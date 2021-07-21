/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package tools

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

type Metric struct {
	name   string
	tags   []string
	values []string
}

type MetricEntry interface {
	Compile() string
}

func NewMetric(name string) *Metric {
	return &Metric{name: name}
}

func (m *Metric) Tag(k string, v interface{}) *Metric {
	m.tags = append(m.tags, fmt.Sprintf("%s=%v", k, v))
	return m
}

func (m *Metric) Value(k string, v interface{}) *Metric {
	m.values = append(m.values, fmt.Sprintf("%s=%v", k, v))
	return m
}

func NowMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (m *Metric) Compile() (data string) {
	var keys, values []string
	for _, v := range m.values {
		tokens := strings.Split(v, "=")
		if len(tokens) == 2 {
			keys = append(keys, fmt.Sprintf("%s.%s", m.name, tokens[0]))
			values = append(values, tokens[1])
		}
	}

	now := NowMS()
	if len(keys) > 0 {
		if len(m.tags) > 0 {
			data = fmt.Sprintf("+%s %s\r\n+%v000000\r\n*%v\r\n+%s\r\n", strings.Join(keys, "|"), strings.Join(m.tags, " "), now, len(keys), strings.Join(values, "\r\n+"))
		} else {
			data = fmt.Sprintf("+%s\r\n+%v000000\r\n*%v\r\n+%s\r\n", strings.Join(keys, "|"), now, len(keys), strings.Join(values, "\r\n+"))
		}
	}
	return
}

type MetricWriter interface {
	Record(MetricEntry)
}

type AkuMetricWriter struct {
	url string
	ch  chan string
}

func (mw *AkuMetricWriter) Record(m MetricEntry) {
	metric := m.Compile()
	if len(metric) == 0 {
		return
	}
	select {
	case mw.ch <- metric:
	default:
		// We drop metrics here if writer stuck some, so it wont stuck the whole program
	}
}

func poll(ch chan string, count int, timeout time.Duration) string {
	entries := []string{}
	timer := time.NewTimer(timeout)
loop:
	for {
		select {
		case metric := <-ch:
			entries = append(entries, metric)
			if len(entries) >= count {
				break loop
			}
		case <-timer.C:
			if len(entries) > 0 {
				break loop
			} else {
				timer.Reset(timeout)
			}
		}
	}
	return strings.Join(entries, "")
}

func NewMetricWriter(url string) (mw MetricWriter, err error) {
	ch := make(chan string, 1000)
	mw = &AkuMetricWriter{url: url, ch: ch}
	go func() {
		conn, err := net.Dial("udp", url)
		for {
			metrics := poll(ch, 100, time.Second*10)
			if len(metrics) == 0 {
				continue
			}
			for {
				if err == nil && conn != nil {
					_, err = conn.Write([]byte(metrics))
					if err == nil {
						continue
					}
				}
				logs.Error("Metric writer connection is unavailable err %v", err)
				time.Sleep(time.Second * 2)
				conn, err = net.Dial("udp", url)
			}
		}
	}()
	return mw, nil
}

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

package metrics

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/polynetwork/bridge-common/base"

	"github.com/beego/beego/v2/server/web"
)

var (
	metrics *Metrics

	metricUpdates = make(chan []byte, 10)
	controller    *MetricController
)

func Init(prefix string) {
	metrics = NewMetrics(prefix)
}

type Metric struct {
	Key   string
	Value interface{}
}

type Metrics struct {
	prefix string
	state  map[string]string
	ch     chan Metric
}

func Record(value interface{}, key string, args ...interface{}) {
	select {
	case metrics.ch <- Metric{Key: fmt.Sprintf(key, args...), Value: value}:
	default:
		// We drop the metrics update here, so the routine wont be blocked
	}
}

func (m *Metrics) start() {
	ticker := time.NewTicker(time.Second)
	for metric := range m.ch {
		m.state[fmt.Sprintf("%s.%s", m.prefix, metric.Key)] = fmt.Sprintf("%v", metric.Value)
		select {
		case <-ticker.C:
			bytes, _ := json.Marshal(m.state)
			metricUpdates <- bytes
		default:
		}
	}
}

func NewMetrics(prefix string) *Metrics {
	if metrics == nil {
		prefix = fmt.Sprintf("%s.%s", prefix, base.ENV)
		metrics = &Metrics{state: map[string]string{}, ch: make(chan Metric, 1000), prefix: prefix}
		go metrics.start()
	}
	return metrics
}

type MetricController struct {
	sync.RWMutex
	state []byte
	web.Controller
}

func (c *MetricController) Metrics() {
	c.RLock()
	state := c.state
	c.RUnlock()
	c.ServeJSON()
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.Output.Body(state)
}

func (c *MetricController) start() {
	for update := range metricUpdates {
		c.Lock()
		c.state = update
		c.Unlock()
	}
}

func NewMetricController() *MetricController {
	if controller != nil {
		controller = &MetricController{state: []byte("{}")}
		go controller.start()
	}
	return controller
}

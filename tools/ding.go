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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/polynetwork/bridge-common/log"
)

var (
	DingUrl string
)

type CardEvent interface {
	Format() (title string, keys []string, values []interface{}, buttons []map[string]string)
}

func PostCardEvent(ev CardEvent) error {
	title, keys, values, buttons := ev.Format()
	return PostDingCardKV(title, keys, values, buttons)
}

func PostDingCardKV(title string, keys []string, values []interface{}, btns []map[string]string) error {
	keys = append(keys, "ReportTime")
	values = append(values, time.Now())
	content := fmt.Sprintf("## %s", title)
	for i, k := range keys {
		if len(values) > i {
			content = fmt.Sprintf("%s\n- %s %v", content, k, values[i])
		}
	}
	err := PostDingCard(title, content, btns)
	if err != nil {
		log.Error("Post dingtalk error", "err", err)
	}
	return err
}

func PostDingCardSimple(title string, body map[string]interface{}, btns []map[string]string) error {
	content := fmt.Sprintf("## %s", title)
	for k, v := range body {
		content = fmt.Sprintf("%s\n- %s %v", content, k, v)
	}
	err := PostDingCard(title, content, btns)
	if err != nil {
		log.Error("Post dingtalk error", "err", err)
	}
	return err
}

func PostDingCard(title, body string, btns interface{}) error {
	payload := map[string]interface{}{}
	payload["msgtype"] = "actionCard"
	card := map[string]interface{}{}
	card["title"] = title
	card["text"] = body
	card["hideAvatar"] = 0
	card["btns"] = btns
	payload["actionCard"] = card
	return PostJson(DingUrl, payload)
}

func PostJson(url string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Info("PostJson response", "Body", string(respBody))
	return nil
}

func PostJsonFor(url string, payload interface{}, result interface{}) error {
	return PostJsonAs(url, nil, payload, result)
}

func PostJsonAs(url string, construct func(*http.Request), payload interface{}, result interface{}) error {
	var body io.Reader
	if payload != nil {
		var data []byte
		var err error
		switch p := payload.(type) {
		case string:
			data = []byte(p)
		case []byte:
			data = p
		default:
			data, err = json.Marshal(payload)
			if err != nil {
				return err
			}
		}
		body = bytes.NewBuffer(data)
	}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	if construct != nil {
		construct(req)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.Error("PostJson response", "Body", string(respBody), "err", err)
	} else {
		log.Debug("PostJson response", "Body", string(respBody))
	}
	return nil
}

func GetJsonFor(url string, result interface{}) error {
	return GetJsonAs(url, nil, result)
}

func GetJsonAs(url string, construct func(*http.Request), result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if construct != nil {
		construct(req)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respBody, result)
	if err != nil {
		log.Error("GetJson response", "url", url, "Body", string(respBody), "err", err)
	} else {
		log.Debug("GetJson response", "url", url, "Body", string(respBody))
	}
	return nil
}

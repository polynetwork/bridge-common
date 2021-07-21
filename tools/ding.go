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
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
)

var (
	DingUrl string
)

func PostDingCardSimple(title string, body map[string]interface{}, btns []map[string]string) error {
	content := fmt.Sprintf("## %s", title)
	for k, v := range body {
		content = fmt.Sprintf("%s\n- %s %v", content, k, v)
	}
	err := PostDingCard(title, content, btns)
	if err != nil {
		logs.Error("Post dingtalk error %s", err)
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
	logs.Info("PostJson response Body:", string(respBody))
	return nil
}

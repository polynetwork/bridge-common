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

package util

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Min(nums ...uint64) (min uint64) {
	if len(nums) > 0 {
		min = nums[0]
	} else {
		return
	}
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return
}

func NonZeroMin(nums ...uint64) (min uint64) {
	for _, num := range nums {
		if min == 0 || (min > num && num != 0) {
			min = num
		}
	}
	return
}

func Fatal(tpl string, args ...interface{}) {
	panic(fmt.Sprintf(tpl, args...))
}

func Json(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func Verbose(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func Reverse(a []byte) (b []byte) {
	size := len(a)
	b = make([]byte, size)
	for i, v := range a {
		b[size-i-1] = v
	}
	return
}

func ReverseHex(a string) (b string) {
	aa, _ := hex.DecodeString(a)
	bb := Reverse(aa)
	return hex.EncodeToString(bb)
}

func LowerHex(a string) string {
	return strings.ToLower(strings.TrimPrefix(a, "0x"))
}

func Concat(data ...[]byte) (b []byte) {
	for _, d := range data {
		b = append(b, d...)
	}
	return
}

func Retry(ctx context.Context, f func() error, interval time.Duration, count int) error {
	c := 0
	var err error
	for {
		if count > 0 {
			if c > count {
				return err
			} else {
				c++
			}
		}
		err = f()
		if err == nil {
			return nil
		}
		select {
		case <-time.After(interval):
		case <-ctx.Done():
			return err
		}
	}
}

func WriteFile(path string, data []byte) error {
	// make sure dir exists
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.Write(data)
	return err
}

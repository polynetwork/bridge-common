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
	"reflect"
	"sync"

	"github.com/polynetwork/bridge-common/log"
)

var _INSTANCES = &SingletonStore{state: map[string]interface{}{}}

type SingletonStore struct {
	sync.Mutex
	state map[string]interface{}
}

func (s *SingletonStore) Single(o Singleton) (interface{}, error) {
	s.Lock()
	defer s.Unlock()
	key := o.Key()
	ins, ok := s.state[key]
	if ok {
		return ins, nil
	}
	log.Info("Creating new singleton instance", "type", reflect.TypeOf(o), "key", key)
	ins, err := o.Create()
	s.state[key] = ins
	return ins, err
}

type Singleton interface {
	Key() string
	Create() (interface{}, error)
}

func Single(s Singleton) (interface{}, error) {
	return _INSTANCES.Single(s)
}

/*
 * Copyright (C) 2022 The poly network Authors
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

package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type LogConfig struct {
	MaxSize uint   // max bytes in MB
	MaxFiles uint  // max log files
	Path string    // main log file path
}

func (c *LogConfig) Writer() io.Writer {
	if c == nil || c.Path == "" {
		return io.Writer(os.Stderr)
	}
	w, err := NewFileWriter(*c)
	if err != nil {
		fmt.Printf("Failed to create file writer, err %v\n", err)
		return io.Writer(os.Stderr)
	}
	return w
}

func NewFileWriter(c LogConfig) (w *FileWriter, err error) {
	w = &FileWriter{LogConfig: c}
	err = w.Init()
	return
}

type FileWriter struct {
	LogConfig
	size, maxSize int
	file *os.File
}


func(w *FileWriter) Init() (err error) {
	if w.MaxSize > 0 {
		w.maxSize = int(w.MaxSize << 20)
		if uint(w.maxSize >> 20) != w.MaxSize {
			return fmt.Errorf("invalid max log file size")
		}
	}
	w.file, err = openFile(w.Path)
	return
}

func (w *FileWriter) Write(p []byte) (n int, err error) {
	if w.maxSize > 0 && w.size + len(p) > w.maxSize {
		err = w.rotate()
		if err != nil {
			fmt.Printf("Failed to rotate log file, path: %s err: %v \n", w.Path, err)
			return
		}
	}
	n, err = w.file.Write(p)
	if err == nil {
		w.size += len(p)
	}
	return
}

func (w *FileWriter) removeLogs() {
	if w.MaxFiles == 0 { return }
	base := filepath.Base(w.Path)
	dir := filepath.Dir(w.Path)
	list, err := os.ReadDir(dir)
	count := uint(0)
	if err == nil {
		for idx := len(list) - 1; idx >= 0; idx-- {
			name := list[idx].Name()
			if list[idx].Type().IsRegular() && strings.HasPrefix(name, base) {
				count++
				if count > w.MaxFiles && name != base {
					os.Remove(filepath.Join(dir, name))
				}
			}
		}
	}
}

func (w *FileWriter) rotate() (err error) {
	if w.file != nil {
		w.file.Sync()
		w.file.Close()
		os.Rename(w.Path, fmt.Sprintf("%s%s", w.Path, time.Now().Format(time.RFC3339)))
		w.removeLogs()
	}
	w.size = 0
	w.file, err = openFile(w.Path)
	return
}

func openFile(path string) (f *os.File, err error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
}


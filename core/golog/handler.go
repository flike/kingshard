// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package golog

import (
	"io"
)

//Handler writes logs to somewhere
type Handler interface {
	Write(p []byte) (n int, err error)
	Close() error
}

//StreamHandler writes logs to a specified io Writer, maybe stdout, stderr, etc...
type StreamHandler struct {
	w io.Writer
}

func NewStreamHandler(w io.Writer) (*StreamHandler, error) {
	h := new(StreamHandler)

	h.w = w

	return h, nil
}

func (h *StreamHandler) Write(b []byte) (n int, err error) {
	return h.w.Write(b)
}

func (h *StreamHandler) Close() error {
	return nil
}

//NullHandler does nothing, it discards anything.
type NullHandler struct {
}

func NewNullHandler() (*NullHandler, error) {
	return new(NullHandler), nil
}

func (h *NullHandler) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (h *NullHandler) Close() error {
	return nil
}

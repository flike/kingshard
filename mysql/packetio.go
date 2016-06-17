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

package mysql

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const (
	defaultReaderSize = 8 * 1024
)

type PacketIO struct {
	rb *bufio.Reader
	wb io.Writer

	Sequence uint8
}

func NewPacketIO(conn net.Conn) *PacketIO {
	p := new(PacketIO)

	p.rb = bufio.NewReaderSize(conn, defaultReaderSize)
	p.wb = conn

	p.Sequence = 0

	return p
}

func (p *PacketIO) ReadPacket() ([]byte, error) {
	header := []byte{0, 0, 0, 0}

	if _, err := io.ReadFull(p.rb, header); err != nil {
		return nil, ErrBadConn
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		return nil, fmt.Errorf("invalid payload length %d", length)
	}

	sequence := uint8(header[3])

	if sequence != p.Sequence {
		return nil, fmt.Errorf("invalid sequence %d != %d", sequence, p.Sequence)
	}

	p.Sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(p.rb, data); err != nil {
		return nil, ErrBadConn
	} else {
		if length < MaxPayloadLen {
			return data, nil
		}

		var buf []byte
		buf, err = p.ReadPacket()
		if err != nil {
			return nil, ErrBadConn
		} else {
			return append(data, buf...), nil
		}
	}
}

//data already have header
func (p *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence

		if n, err := p.wb.Write(data[:4+MaxPayloadLen]); err != nil {
			return ErrBadConn
		} else if n != (4 + MaxPayloadLen) {
			return ErrBadConn
		} else {
			p.Sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	if n, err := p.wb.Write(data); err != nil {
		return ErrBadConn
	} else if n != len(data) {
		return ErrBadConn
	} else {
		p.Sequence++
		return nil
	}
}

func (p *PacketIO) WritePacketBatch(total, data []byte, direct bool) ([]byte, error) {
	if data == nil {
		//only flush the buffer
		if direct == true {
			n, err := p.wb.Write(total)
			if err != nil {
				return nil, ErrBadConn
			}
			if n != len(total) {
				return nil, ErrBadConn
			}
		}
		return total, nil
	}

	length := len(data) - 4
	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence
		total = append(total, data[:4+MaxPayloadLen]...)

		p.Sequence++
		length -= MaxPayloadLen
		data = data[MaxPayloadLen:]
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	total = append(total, data...)
	p.Sequence++

	if direct {
		if n, err := p.wb.Write(total); err != nil {
			return nil, ErrBadConn
		} else if n != len(total) {
			return nil, ErrBadConn
		}
	}
	return total, nil
}

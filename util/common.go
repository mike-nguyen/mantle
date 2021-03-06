// Copyright 2018 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"unsafe"
)

const (
	LITTLE Endian = iota // little endian
	BIG                  // big endian
)

// Endianness of the platform - big or little
type Endian int

var HostEndianness Endian

func init() {
	// Determine endianness - https://stackoverflow.com/questions/51332658/any-better-way-to-check-endianness-in-go
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0x0100)

	switch buf {
	case [2]byte{0x00, 0x01}:
		HostEndianness = LITTLE
	case [2]byte{0x01, 0x00}:
		HostEndianness = BIG
	default:
		HostEndianness = LITTLE
	}
}

func StrToPtr(s string) *string {
	return &s
}

func BoolToPtr(b bool) *bool {
	return &b
}

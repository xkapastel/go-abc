/**
This file is a part of ABC.
Copyright (C) 2018 Matthew Blount

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public
License along with this program.  If not, see
<https://www.gnu.org/licenses/>.
**/

package abc

import (
	"io"
)

const (
	byteBegin byte = iota
	byteApp
	byteBox
	byteCat
	byteCopy
	byteDrop
	byteSwap
	byteEnd
	byteHash
)

// DecodeBlock creates a block from a stream of bytecode.
func DecodeBlock(src io.ByteReader) (Block, error) {
	var build []Block
	var stack [][]Block
	for {
		code, err := src.ReadByte()
		switch {
		case err == io.EOF:
			// len(stack) > 0 not an error, just let it go.
			return newCatN(build...), nil
		case err != nil:
			return nil, err
		}
		switch code {
		case byteBegin:
			stack = append(stack, build)
			build = nil
		case byteEnd:
			if len(stack) == 0 {
				// Not an error, just let it go.
				continue
			}
			body := newCatN(build...)
			wrap := &mkBox{body}
			build = stack[len(stack)-1]
			build = append(build, wrap)
			stack = stack[:len(stack)-1]
		case byteApp:
			build = append(build, opApp{})
		case byteBox:
			build = append(build, opBox{})
		case byteCat:
			build = append(build, opCat{})
		case byteCopy:
			build = append(build, opCopy{})
		case byteDrop:
			build = append(build, opDrop{})
		case byteSwap:
			build = append(build, opSwap{})
		case byteHash:
			var name []byte
			for i := 0; i < 32; i++ {
				code, err := src.ReadByte()
				if err != nil {
					return nil, err
				}
				name = append(name, code)
			}
			link := opLink{name}
			build = append(build, link)
		default:
			// Not an error, just let it go.
		}
	}
}

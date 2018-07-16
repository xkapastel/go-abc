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
	CodeBegin byte = iota
	CodeOpApp
	CodeOpBox
	CodeOpCat
	CodeOpCopy
	CodeOpDrop
	CodeOpSwap
	CodeOpNoCopy
	CodeOpNoDrop
	CodeOpNoSwap
	CodeOpEq
	CodeOpNeq
	CodeOpTag
	CodeOpLink
	CodeEnd
)

// DecodeBlock reads a block from a stream of bytecode.
func DecodeBlock(src io.ByteReader) (Block, error) {
	var dead uint
	var build []Block
	var stack [][]Block
	for {
		code, err := src.ReadByte()
		switch {
		case err == io.EOF:
			return newCatN(build...), nil
		case err != nil:
			return nil, err
		case code == CodeBegin:
			stack = append(stack, build)
			build = nil
		case code == CodeEnd:
			if len(stack) == 0 {
				dead++
				continue
			}
			body := newCatN(build...)
			wrap := body.Box()
			build = stack[len(stack)-1]
			build = append(build, wrap)
			stack = stack[:len(stack)-1]
		case code == CodeOpApp:
			build = append(build, App)
		case code == CodeOpBox:
			build = append(build, Box)
		case code == CodeOpCat:
			build = append(build, Cat)
		case code == CodeOpCopy:
			build = append(build, Copy)
		case code == CodeOpDrop:
			build = append(build, Drop)
		case code == CodeOpSwap:
			build = append(build, Swap)
		case code == CodeOpNoCopy:
			build = append(build, NoCopy)
		case code == CodeOpNoDrop:
			build = append(build, NoDrop)
		case code == CodeOpNoSwap:
			build = append(build, NoSwap)
		case code == CodeOpEq:
			build = append(build, Eq)
		case code == CodeOpNeq:
			build = append(build, Neq)
		case code == CodeOpTag:
			build = append(build, Tag)
		case code == CodeOpLink:
			var buf []byte
			for i := 0; i < 32; i++ {
				value, err := src.ReadByte()
				if err != nil {
					return nil, err
				}
				buf = append(buf, value)
			}
			link := opLink{buf}
			build = append(build, link)
		default:
			dead++
		}
	}
}

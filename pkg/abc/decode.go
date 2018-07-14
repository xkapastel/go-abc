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
	byteId byte = iota
	byteOpId
	byteOpApp
	byteOpBox
	byteOpCat
	byteOpCopy
	byteOpDrop
	byteOpSwap
	byteOpEq
	byteOpLink
	byteMkBox
	byteMkCat
	byteMkLink
	byteRmBox
	byteRmCat
	byteRmLink
	byteMkCatAll
	byteCopy
	byteDrop
	byteSwap
	byteStep
	byteReduce
)

const kDefaultQuota int = 255

func DecodeBlock(src io.ByteReader) (Block, error) {
	ctx := NewBuild()
	for {
		code, err := src.ReadByte()
		switch {
		case err == io.EOF:
			return ctx.Block(), nil
		case err != nil:
			return nil, err
		case code == byteOpId:
			ctx.OpId()
		case code == byteOpApp:
			ctx.OpApp()
		case code == byteOpBox:
			ctx.OpBox()
		case code == byteOpCat:
			ctx.OpCat()
		case code == byteOpCopy:
			ctx.OpCopy()
		case code == byteOpDrop:
			ctx.OpDrop()
		case code == byteOpSwap:
			ctx.OpSwap()
		case code == byteMkBox:
			ctx.MkBox()
		case code == byteMkCat:
			ctx.MkCat()
		case code == byteMkLink:
			ctx.MkLink()
		case code == byteRmBox:
			ctx.RmBox()
		case code == byteRmCat:
			ctx.RmCat()
		case code == byteRmLink:
			ctx.RmLink()
		case code == byteCopy:
			ctx.Copy()
		case code == byteDrop:
			ctx.Drop()
		case code == byteSwap:
			ctx.Swap()
		case code == byteReduce:
			ctx.Reduce()
		}
	}
}

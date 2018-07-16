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
	"fmt"
	"io"
)

type mkBox struct{ body Block }

func (block *mkBox) Box() Block { return &mkBox{block} }
func (block *mkBox) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block *mkBox) Reduce(quota int) Block { return block }
func (block *mkBox) Encode(dst io.ByteWriter) error {
	if err := dst.WriteByte(CodeBegin); err != nil {
		return err
	}
	if err := block.body.Encode(dst); err != nil {
		return err
	}
	return dst.WriteByte(CodeEnd)
}
func (block *mkBox) String() string {
	body := block.body.String()
	return fmt.Sprintf("[%s]", body)
}
func (lhs *mkBox) Eq(rhs Block) bool {
	switch rhs := rhs.(type) {
	case *mkBox:
		return lhs.body.Eq(rhs.body)
	default:
		return false
	}
}
func (block *mkBox) Copy() bool { return block.body.Copy() }
func (block *mkBox) Drop() bool { return block.body.Drop() }
func (block *mkBox) Swap() bool { return block.body.Swap() }
func (block *mkBox) step(ctx *reduce) bool {
	ctx.push(block)
	return false
}

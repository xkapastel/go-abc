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
	"encoding/hex"
	"fmt"
	"io"
)

type opLink struct {
	value []byte
}

func (block opLink) Box() Block { return &mkBox{block} }
func (block opLink) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opLink) Reduce(quota int) Block { return block }
func (block opLink) Encode(dst io.ByteWriter) error {
	if err := dst.WriteByte(CodeOpLink); err != nil {
		return err
	}
	for _, value := range block.value {
		if err := dst.WriteByte(value); err != nil {
			return err
		}
	}
	return nil
}
func (block opLink) String() string {
	name := hex.EncodeToString(block.value)
	return fmt.Sprintf("#%s", name)
}
func (lhs opLink) Eq(rhs Block) bool {
	_, ok := rhs.(opLink)
	return ok
}
func (block opLink) Copy() bool { return true }
func (block opLink) Drop() bool { return true }
func (block opLink) Swap() bool { return true }
func (block opLink) step(ctx *reduce) bool {
	ctx.clear(block)
	return false
}

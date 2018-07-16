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

type opId struct{}

func (block opId) Box() Block { return &mkBox{block} }
func (block opId) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(block, rest)
}
func (block opId) Reduce(quota int) Block         { return block }
func (block opId) Encode(dst io.ByteWriter) error { return nil }
func (block opId) String() string                 { return "id" }
func (lhs opId) Eq(rhs Block) bool {
	_, ok := rhs.(opId)
	return ok
}
func (block opId) Copy() bool            { return true }
func (block opId) Drop() bool            { return true }
func (block opId) Swap() bool            { return true }
func (block opId) step(ctx *reduce) bool { return false }
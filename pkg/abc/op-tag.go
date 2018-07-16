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

type opTag struct{}

func (block opTag) Reduce(quota int) Block { return block }
func (block opTag) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(CodeOpTag)
}
func (block opTag) String() string { return "t" }
func (lhs opTag) Eq(rhs Block) bool {
	_, ok := rhs.(opTag)
	return ok
}
func (block opTag) Copy() bool { return true }
func (block opTag) Drop() bool { return true }
func (block opTag) Swap() bool { return true }
func (block opTag) step(ctx *reduce) bool {
	return false
}

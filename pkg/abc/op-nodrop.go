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

import ()

type opNoDrop struct{}

func (block opNoDrop) String() string { return "%nodrop" }
func (lhs opNoDrop) eq(rhs Block) bool {
	_, ok := rhs.(opNoDrop)
	return ok
}
func (block opNoDrop) Copy() bool { return true }
func (block opNoDrop) Drop() bool { return false }
func (block opNoDrop) Swap() bool { return true }
func (block opNoDrop) step(ctx *reduce) bool {
	ctx.kill.push(block)
	return false
}

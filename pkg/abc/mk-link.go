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
)

type mkLink struct {
	value []byte
}

func (object mkLink) String() string {
	name := hex.EncodeToString(object.value)
	return fmt.Sprintf("#%s", name)
}
func (lhs mkLink) eq(rhs Object) bool {
	_, ok := rhs.(mkLink)
	return ok
}
func (object mkLink) step(ctx *rewrite) bool {
	ctx.clear(object)
	return false
}

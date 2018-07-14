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

type opLink struct{ name []byte }

func (tau opLink) Box() Block { return &mkBox{tau} }
func (tau opLink) Cat(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(tau, rest)
}
func (tau opLink) Reduce(quota int) Block { return tau }
func (tau opLink) Encode(dst io.ByteWriter) error {
	for _, value := range tau.name {
		if err := dst.WriteByte(value); err != nil {
			return err
		}
	}
	return nil
}
func (tau opLink) String() string {
	name := hex.EncodeToString(tau.name)
	return fmt.Sprintf("#%s", name)
}
func (lhs opLink) Eq(rhs Block) bool {
	switch rhs := rhs.(type) {
	case opLink:
		for i := 0; i < 32; i++ {
			if lhs.name[i] != rhs.name[i] {
				return false
			}
		}
		return true
	default:
		return false
	}
}

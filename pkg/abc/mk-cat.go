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
)

type mkCat struct{ fst, snd Block }

func newCat(fst, snd Block) Block {
	var ok bool
	_, ok = fst.(opId)
	if ok {
		return snd
	}
	_, ok = snd.(opId)
	if ok {
		return fst
	}
	switch fst := fst.(type) {
	case *mkCat:
		inner := newCat(fst.snd, snd)
		return newCat(fst.fst, inner)
	default:
		return &mkCat{fst, snd}
	}
}

// NewCat catenates the given blocks.
func NewCat(xs ...Block) Block {
	var block Block = opId{}
	for i := len(xs) - 1; i >= 0; i-- {
		child := xs[i]
		block = newCat(child, block)
	}
	return block
}

// NewCatR catenates the given blocks in reverse.
func NewCatR(xs ...Block) Block {
	var block Block = opId{}
	for _, child := range xs {
		block = newCat(child, block)
	}
	return block
}
func (block *mkCat) String() string {
	var ok bool
	_, ok = block.fst.(opId)
	if ok {
		return block.snd.String()
	}
	_, ok = block.snd.(opId)
	if ok {
		return block.fst.String()
	}
	fst := block.fst.String()
	snd := block.snd.String()
	return fmt.Sprintf("%s %s", fst, snd)
}
func (lhs *mkCat) eq(rhs Block) bool {
	switch rhs := rhs.(type) {
	case *mkCat:
		if lhs.fst.eq(rhs.fst) {
			return lhs.snd.eq(rhs.snd)
		}
		return false
	default:
		return false
	}
}
func (block *mkCat) step(ctx *reduce) bool {
	ctx.work.push(block.snd)
	ctx.work.push(block.fst)
	return false
}

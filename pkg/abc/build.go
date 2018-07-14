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
	"strings"
)

type Build interface {
	OpId()
	OpApp()
	OpBox()
	OpCat()
	OpCopy()
	OpDrop()
	OpSwap()
	MkBox()
	MkCat()
	MkLink()
	RmBox()
	RmCat()
	RmLink()
	Copy()
	Drop()
	Swap()
	Step()
	Reduce()
	Block() Block
}

type build struct {
	data BlockStack
}

func NewBuild() Build {
	return &build{
		data: NewBlockStack(),
	}
}

func (ctx *build) OpId()   { ctx.data.Push(Id) }
func (ctx *build) OpApp()  { ctx.data.Push(App) }
func (ctx *build) OpBox()  { ctx.data.Push(Box) }
func (ctx *build) OpCat()  { ctx.data.Push(Cat) }
func (ctx *build) OpCopy() { ctx.data.Push(Copy) }
func (ctx *build) OpDrop() { ctx.data.Push(Drop) }
func (ctx *build) OpSwap() { ctx.data.Push(Swap) }

func (ctx *build) MkBox() {
	if ctx.data.Len() > 0 {
		lhs := ctx.data.Pop()
		rhs := lhs.Box()
		ctx.data.Push(rhs)
	}
}

func (ctx *build) RmBox() {
	if ctx.data.Len() > 0 {
		block := ctx.data.Peek(0)
		switch block := block.(type) {
		case *mkBox:
			ctx.data.Pop()
			ctx.data.Push(block.body)
		}
	}
}

func (ctx *build) MkCat() {
	if ctx.data.Len() > 1 {
		rhs := ctx.data.Pop()
		lhs := ctx.data.Pop()
		cat := lhs.Cat(rhs)
		ctx.data.Push(cat)
	}
}

func (ctx *build) RmCat() {
	if ctx.data.Len() > 0 {
		block := ctx.data.Peek(0)
		switch block := block.(type) {
		case *mkCat:
			ctx.data.Pop()
			ctx.data.Push(block.fst)
			ctx.data.Push(block.snd)
		}
	}
}

func (ctx *build) MkLink() {
	//
}

func (ctx *build) RmLink() {
	//
}

func (ctx *build) Copy() {
	if ctx.data.Len() > 0 {
		block := ctx.data.Peek(0)
		ctx.data.Push(block)
	}
}

func (ctx *build) Drop() {
	if ctx.data.Len() > 0 {
		ctx.data.Pop()
	}
}

func (ctx *build) Swap() {
	if ctx.data.Len() > 1 {
		fst := ctx.data.Pop()
		snd := ctx.data.Pop()
		ctx.data.Push(fst)
		ctx.data.Push(snd)
	}
}

func (ctx *build) Step() {

}

func (ctx *build) Reduce() {
	if ctx.data.Len() > 0 {
		lhs := ctx.data.Pop()
		rhs := lhs.Reduce(kDefaultQuota)
		ctx.data.Push(rhs)
	}
}

func (ctx *build) Block() Block {
	return ctx.data.Block()
}

func (ctx *build) String() string {
	var buf []string
	ctx.data.Each(func(block Block) {
		buf = append(buf, block.String())
	})
	return strings.Join(buf, " - ")
}

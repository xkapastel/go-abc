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

type reduce struct {
	kill *stack
	data *stack
	work *stack
}

func newReduce(init Block) *reduce {
	work := newStack()
	work.Push(init)
	return &reduce{
		kill: newStack(),
		data: newStack(),
		work: work,
	}
}
func (ctx *reduce) arity() int { return ctx.data.Len() }
func (ctx *reduce) push(block Block) {
	ctx.data.Push(block)
}
func (ctx *reduce) peek(index int) Block {
	return ctx.data.Peek(index)
}
func (ctx *reduce) pop() Block {
	return ctx.data.Pop()
}
func (ctx *reduce) queue(block Block) {
	ctx.work.Push(block)
}
func (ctx *reduce) clear(block Block) {
	ctx.data.Each(ctx.kill.Push)
	ctx.kill.Push(block)
	ctx.data.Clear()
}
func (ctx *reduce) stash(block Block) {
	ctx.kill.Push(block)
}
func (ctx *reduce) step() bool {
	for ctx.work.Len() > 0 {
		block := ctx.work.Pop()
		if block.step(ctx) {
			break
		}
	}
	return ctx.work.Len() > 0
}
func (ctx *reduce) Block() Block {
	var buf []Block
	ctx.work.Each(func(block Block) {
		buf = append(buf, block)
	})
	work := NewCatR(buf...)
	data := ctx.data.Block()
	kill := ctx.kill.Block()
	return NewCat(kill, work, data)
}

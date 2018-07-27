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

// Rewrite rewrites an object until it either reaches a normal
// form or the effort quota is exhausted.
func Rewrite(object Object, quota int) Object {
	ctx := newRewrite(object)
	busy := true
	for busy && quota > 0 {
		busy = ctx.step()
		quota--
	}
	return ctx.Object()
}

type rewrite struct {
	kill *stack
	data *stack
	work *stack
}

func newRewrite(init Object) *rewrite {
	work := newStack()
	work.push(init)
	return &rewrite{
		kill: newStack(),
		data: newStack(),
		work: work,
	}
}
func (ctx *rewrite) clear(object Object) {
	ctx.data.each(ctx.kill.push)
	ctx.kill.push(object)
	ctx.data.clear()
}
func (ctx *rewrite) step() bool {
	for ctx.work.len() > 0 {
		object := ctx.work.pop()
		if object.step(ctx) {
			break
		}
	}
	return ctx.work.len() > 0
}
func (ctx *rewrite) Object() Object {
	var buf []Object
	ctx.work.each(func(object Object) {
		buf = append(buf, object)
	})
	work := newCatsR(buf...)
	data := ctx.data.Object()
	kill := ctx.kill.Object()
	return newCats(kill, work, data)
}

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

type stack struct {
	data []Object
}

func newStack() *stack {
	return &stack{
		data: make([]Object, 0),
	}
}
func (ctx *stack) push(object Object) {
	ctx.data = append(ctx.data, object)
}
func (ctx *stack) peek(index int) Object {
	return ctx.data[len(ctx.data)-1-index]
}
func (ctx *stack) pop() Object {
	object := ctx.data[len(ctx.data)-1]
	ctx.data = ctx.data[:len(ctx.data)-1]
	return object
}
func (ctx *stack) clear() {
	ctx.data = nil
}
func (ctx *stack) len() int { return len(ctx.data) }
func (ctx *stack) Object() Object {
	return newCats(ctx.data...)
}
func (ctx *stack) each(fn func(Object)) {
	for _, value := range ctx.data {
		fn(value)
	}
}

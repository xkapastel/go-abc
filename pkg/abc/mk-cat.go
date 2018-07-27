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

type mkCat struct{ fst, snd Object }

func newCat(fst, snd Object) Object {
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

func newCats(xs ...Object) Object {
	var object Object = opId{}
	for i := len(xs) - 1; i >= 0; i-- {
		child := xs[i]
		object = newCat(child, object)
	}
	return object
}

func newCatsR(xs ...Object) Object {
	var object Object = opId{}
	for _, child := range xs {
		object = newCat(child, object)
	}
	return object
}
func (object *mkCat) String() string {
	var ok bool
	_, ok = object.fst.(opId)
	if ok {
		return object.snd.String()
	}
	_, ok = object.snd.(opId)
	if ok {
		return object.fst.String()
	}
	fst := object.fst.String()
	snd := object.snd.String()
	return fmt.Sprintf("%s %s", fst, snd)
}
func (lhs *mkCat) eq(rhs Object) bool {
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
func (object *mkCat) step(ctx *rewrite) bool {
	ctx.work.push(object.snd)
	ctx.work.push(object.fst)
	return false
}

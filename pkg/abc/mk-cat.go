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
	"io"
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
func newCatN(xs ...Block) Block {
	var block Block = opId{}
	for i := len(xs) - 1; i >= 0; i-- {
		child := xs[i]
		block = newCat(child, block)
	}
	return block
}
func newCatNFlip(xs ...Block) Block {
	var block Block = opId{}
	for _, child := range xs {
		block = newCat(child, block)
	}
	return block
}
func (tau *mkCat) Box() Block { return &mkBox{tau} }
func (tau *mkCat) Catenate(xs ...Block) Block {
	rest := newCatN(xs...)
	return newCat(tau, rest)
}
func (tau *mkCat) Reduce(quota int) Block {
	var trash []Block
	var stack []Block
	var queue []Block = []Block{tau}

	clear := func(block Block) {
		for _, value := range stack {
			trash = append(trash, value)
		}
		trash = append(trash, block)
		stack = nil
	}

	for quota > 0 && len(queue) > 0 {
		block := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if block == nil {
			panic(block)
		}
		switch block := block.(type) {
		case opId:
			//
		case opApp:
			if len(stack) == 0 {
				clear(block)
				continue
			}
			value, ok := stack[len(stack)-1].(*mkBox)
			if !ok {
				clear(block)
				continue
			}
			stack = stack[:len(stack)-1]
			queue = append(queue, value.body)
			quota--
		case opBox:
			if len(stack) == 0 {
				clear(block)
				continue
			}
			value := stack[len(stack)-1]
			boxed := &mkBox{value}
			stack[len(stack)-1] = boxed
			quota--
		case opCat:
			if len(stack) < 2 {
				clear(block)
				continue
			}
			var ok bool
			rhs, ok := stack[len(stack)-1].(*mkBox)
			if !ok {
				clear(block)
				continue
			}
			lhs, ok := stack[len(stack)-2].(*mkBox)
			if !ok {
				clear(block)
				continue
			}
			body := newCat(lhs.body, rhs.body)
			wrap := &mkBox{body}
			stack = stack[:len(stack)-2]
			stack = append(stack, wrap)
			quota--
		case opCopy:
			if len(stack) == 0 {
				clear(block)
				continue
			}
			value := stack[len(stack)-1]
			stack = append(stack, value)
			quota--
		case opDrop:
			if len(stack) == 0 {
				clear(block)
				continue
			}
			stack = stack[:len(stack)-1]
			quota--
		case opSwap:
			if len(stack) < 2 {
				clear(block)
				continue
			}
			tmp := stack[len(stack)-1]
			stack[len(stack)-1] = stack[len(stack)-2]
			stack[len(stack)-2] = tmp
			quota--
		case *mkBox:
			stack = append(stack, block)
			quota--
		case *mkCat:
			queue = append(queue, block.snd)
			queue = append(queue, block.fst)
		default:
			panic("unknown block")
		}
	}
	qq := newCatNFlip(queue...)
	ss := newCatN(stack...)
	tt := newCatN(trash...)
	return newCatN(tt, ss, qq)
}
func (tau *mkCat) Encode(dst io.ByteWriter) error {
	if err := tau.fst.Encode(dst); err != nil {
		return err
	}
	return tau.snd.Encode(dst)
}
func (tau *mkCat) String() string {
	var ok bool
	_, ok = tau.fst.(opId)
	if ok {
		return tau.snd.String()
	}
	_, ok = tau.snd.(opId)
	if ok {
		return tau.fst.String()
	}
	fst := tau.fst.String()
	snd := tau.snd.String()
	return fmt.Sprintf("%s %s", fst, snd)
}
func (lhs *mkCat) Equals(rhs Block) bool {
	switch rhs := rhs.(type) {
	case *mkCat:
		if lhs.fst.Equals(rhs.fst) {
			return lhs.snd.Equals(rhs.snd)
		}
		return false
	default:
		return false
	}
}

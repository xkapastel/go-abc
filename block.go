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

const (
	byteBegin byte = iota
	byteApp
	byteBox
	byteCat
	byteCopy
	byteDrop
	byteSwap
	byteEnd
	byteHash
)

type Block interface {
	Box() Block
	Catenate(...Block) Block
	Reduce(int) Block
	Encode(io.ByteWriter) error
	String() string
}

var Id Block
var Apply Block
var Box Block
var Catenate Block
var Copy Block
var Drop Block
var Swap Block

func init() {
	Id = opId{}
	Apply = opApp{}
	Box = opBox{}
	Catenate = opCat{}
	Copy = opCopy{}
	Drop = opDrop{}
	Swap = opSwap{}
}

func Decode(src io.ByteReader) (Block, error) {
	var build []Block
	var stack [][]Block
	for {
		byte, err := src.ReadByte()
		switch {
		case err == io.EOF:
			if len(stack) != 0 {
				return nil, fmt.Errorf("Unbalanced block")
			}
			return newCat(build...), nil
		case err != nil:
			return nil, err
		}
		switch byte {
		case byteBegin:
			stack = append(stack, build)
			build = nil
		case byteEnd:
			if len(stack) == 0 {
				return nil, fmt.Errorf("Unbalanced block")
			}
			body := newCat(build...)
			wrap := &box{body}
			build = stack[len(stack)-1]
			build = append(build, wrap)
			stack = stack[:len(stack)-1]
		case byteApp:
			build = append(build, opApp{})
		case byteBox:
			build = append(build, opBox{})
		case byteCat:
			build = append(build, opCat{})
		case byteCopy:
			build = append(build, opCopy{})
		case byteDrop:
			build = append(build, opDrop{})
		case byteSwap:
			build = append(build, opSwap{})
		default:
			panic("Unknown bytecode")
		}
	}
}

func newCat(xs ...Block) Block {
	var block Block = opId{}
	for i := len(xs) - 1; i >= 0; i-- {
		child := xs[i]
		block = &cat{child, block}
	}
	return block
}

type opId struct{}
type opApp struct{}
type opBox struct{}
type opCat struct{}
type opCopy struct{}
type opDrop struct{}
type opSwap struct{}
type box struct{ body Block }
type cat struct{ fst, snd Block }

func (tau opId) Box() Block { return &box{tau} }
func (tau opId) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opId) Reduce(quota int) Block         { return tau }
func (tau opId) Encode(dst io.ByteWriter) error { return nil }
func (tau opId) String() string                 { return "" }

func (tau opApp) Box() Block { return &box{tau} }
func (tau opApp) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opApp) Reduce(quota int) Block { return tau }
func (tau opApp) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteApp)
}
func (tau opApp) String() string { return "a" }

func (tau opBox) Box() Block { return &box{tau} }
func (tau opBox) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opBox) Reduce(quota int) Block { return tau }
func (tau opBox) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteBox)
}
func (tau opBox) String() string { return "b" }

func (tau opCat) Box() Block { return &box{tau} }
func (tau opCat) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opCat) Reduce(quota int) Block { return tau }
func (tau opCat) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteCat)
}
func (tau opCat) String() string { return "c" }

func (tau opCopy) Box() Block { return &box{tau} }
func (tau opCopy) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opCopy) Reduce(quota int) Block { return tau }
func (tau opCopy) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteCopy)
}
func (tau opCopy) String() string { return "d" }

func (tau opDrop) Box() Block { return &box{tau} }
func (tau opDrop) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opDrop) Reduce(quota int) Block { return tau }
func (tau opDrop) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteDrop)
}
func (tau opDrop) String() string { return "e" }

func (tau opSwap) Box() Block { return &box{tau} }
func (tau opSwap) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau opSwap) Reduce(quota int) Block { return tau }
func (tau opSwap) Encode(dst io.ByteWriter) error {
	return dst.WriteByte(byteSwap)
}
func (tau opSwap) String() string { return "f" }

func (tau *box) Box() Block { return &box{tau} }
func (tau *box) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau *box) Reduce(quota int) Block { return tau }
func (tau *box) Encode(dst io.ByteWriter) error {
	if err := dst.WriteByte(byteBegin); err != nil {
		return err
	}
	if err := tau.body.Encode(dst); err != nil {
		return err
	}
	return dst.WriteByte(byteEnd)
}
func (tau *box) String() string {
	body := tau.body.String()
	return fmt.Sprintf("[%s]", body)
}

func (tau *cat) Box() Block { return &box{tau} }
func (tau *cat) Catenate(xs ...Block) Block {
	rest := newCat(xs...)
	return &cat{tau, rest}
}
func (tau *cat) Reduce(quota int) Block {
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
			value, ok := stack[len(stack)-1].(*box)
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
			boxed := &box{value}
			stack[len(stack)-1] = boxed
			quota--
		case opCat:
			if len(stack) < 2 {
				clear(block)
				continue
			}
			var ok bool
			rhs, ok := stack[len(stack)-1].(*box)
			if !ok {
				clear(block)
				continue
			}
			lhs, ok := stack[len(stack)-2].(*box)
			if !ok {
				clear(block)
				continue
			}
			body := newCat(lhs.body, rhs.body)
			wrap := &box{body}
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
		case *box:
			stack = append(stack, block)
			quota--
		case *cat:
			queue = append(queue, block.snd)
			queue = append(queue, block.fst)
		default:
			panic("unknown block")
		}
	}
	qq := newCat(queue...)
	ss := newCat(stack...)
	tt := newCat(trash...)
	return newCat(tt, ss, qq)
}
func (tau *cat) Encode(dst io.ByteWriter) error {
	if err := tau.fst.Encode(dst); err != nil {
		return err
	}
	return tau.snd.Encode(dst)
}
func (tau *cat) String() string {
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

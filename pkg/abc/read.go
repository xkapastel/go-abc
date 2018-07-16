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
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var cache map[string]Block
var cycle map[string]bool

func init() {
	cache = make(map[string]Block)
	cycle = make(map[string]bool)
}

func readFile(name string) (Block, error) {
	block, ok := cache[name]
	if ok {
		return block, nil
	}
	if cycle[name] {
		msg := "`%s` contains a cycle"
		err := fmt.Errorf(msg, name)
		return nil, err
	}
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	cycle[name] = true
	block, err = Read(file)
	cycle[name] = false
	if err != nil {
		cache[name] = block
	}
	return block, err
}

func Read(src io.Reader) (Block, error) {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}
	text := string(buf)
	text = strings.Replace(text, "[", "[ ", -1)
	text = strings.Replace(text, "]", " ]", -1)
	words := strings.Split(text, " ")
	ident := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9-]+$")
	var build []Block
	var stack [][]Block
	for _, word := range words {
		switch {
		case word == "[":
			stack = append(stack, build)
			build = nil
		case word == "]":
			if len(stack) == 0 {
				return nil, fmt.Errorf("Unbalanced block")
			}
			body := newCatN(build...)
			wrap := body.Box()
			build = stack[len(stack)-1]
			build = append(build, wrap)
			stack = stack[:len(stack)-1]
		case word == "a":
			build = append(build, App)
		case word == "b":
			build = append(build, Box)
		case word == "c":
			build = append(build, Cat)
		case word == "cp":
			build = append(build, Copy)
		case word == "dp":
			build = append(build, Drop)
		case word == "sp":
			build = append(build, Swap)
		case word == "nc":
			build = append(build, NoCopy)
		case word == "nd":
			build = append(build, NoDrop)
		case word == "ns":
			build = append(build, NoSwap)
		case word == "eq":
			build = append(build, Eq)
		case word == "nq":
			build = append(build, Neq)
		case word == "t":
			// build = append(build, Tag)
		case len(word) == 0:
			continue
		case len(word) <= 2:
			msg := "`%s`: words of length <= 2 are reserved"
			err := fmt.Errorf(msg, word)
			return nil, err
		case ident.MatchString(word):
			block, err := readFile(word)
			if err != nil {
				return nil, err
			}
			build = append(build, block)
		}
	}
	if len(stack) != 0 {
		return nil, fmt.Errorf("Unbalanced block")
	}
	return newCatN(build...), nil
}
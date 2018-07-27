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
	"strconv"
	"strings"
)

var cache map[string]Object
var cycle map[string]bool

func init() {
	cache = make(map[string]Object)
	cycle = make(map[string]bool)
}

func readFile(name string) (Object, error) {
	object, ok := cache[name]
	if ok {
		return object, nil
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
	object, err = Read(file)
	cycle[name] = false
	if err != nil {
		cache[name] = object
	}
	return object, err
}

// Read creates an object from a string. Free variables are resolved
// using files in the current directory, and cyclic definitions are
// not allowed.
func Read(src io.Reader) (Object, error) {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}
	text := string(buf)
	text = strings.Replace(text, "[", "[ ", -1)
	text = strings.Replace(text, "]", " ]", -1)
	text = strings.Replace(text, "\t", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.Replace(text, "\n", " ", -1)
	words := strings.Split(text, " ")
	num := regexp.MustCompile("^(([-+]?[0-9]*\\.?[0-9]+([eE][-+]?[0-9]+)?))$")
	ident := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9-]+$")
	var build []Object
	var stack [][]Object
	for _, word := range words {
		switch {
		case word == "[":
			stack = append(stack, build)
			build = nil
		case word == "]":
			if len(stack) == 0 {
				return nil, fmt.Errorf("Unbalanced block")
			}
			body := newCats(build...)
			wrap := newBox(body)
			build = stack[len(stack)-1]
			build = append(build, wrap)
			stack = stack[:len(stack)-1]
		case word == "a":
			build = append(build, opApp{})
		case word == "b":
			build = append(build, opBox{})
		case word == "c":
			build = append(build, opCat{})
		case word == "d":
			build = append(build, opCopy{})
		case word == "e":
			build = append(build, opDrop{})
		case word == "f":
			build = append(build, opSwap{})
		case len(word) == 0:
			continue
		case num.MatchString(word):
			value, err := strconv.ParseFloat(word, 64)
			if err != nil {
				return nil, err
			}
			object := newNum(value)
			build = append(build, object)
		case len(word) <= 2:
			msg := "`%s`: words of length <= 2 are reserved"
			err := fmt.Errorf(msg, word)
			return nil, err
		case ident.MatchString(word):
			object := newVar(word)
			build = append(build, object)
		}
	}
	if len(stack) != 0 {
		return nil, fmt.Errorf("Unbalanced block")
	}
	return newCats(build...), nil
}

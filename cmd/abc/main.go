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

package main

import (
	"bufio"
	"fmt"
	"github.com/xkapastel/abc/pkg/abc"
	"os"
)

func usage() {
	fmt.Println(`usage: abc [command]

Available commands are:

read    - read a stream of bytecode
parse   - read a string
reduce  - optimize a block
reify   - convert a block to a syntax tree
reflect - convert a syntax tree to a block
`)
	os.Exit(1)
}

const prompt = "abc> "

func main() {
	const defaultQuota = 1000
	if len(os.Args) != 2 {
		usage()
	}
	switch os.Args[1] {
	case "read":
		stdin := bufio.NewReader(os.Stdin)
		block, err := abc.DecodeBlock(stdin)
		if err != nil {
			panic(err)
		}
		fmt.Println(block)
	case "parse":
	case "quote":
	case "reduce":
		stdin := bufio.NewReader(os.Stdin)
		lhs, err := abc.DecodeBlock(stdin)
		if err != nil {
			panic(err)
		}
		stdout := bufio.NewWriter(os.Stdout)
		rhs := lhs.Reduce(defaultQuota)
		rhs.Encode(stdout)
		stdout.Flush()
	case "reify":
	case "reflect":
	default:
		usage()
	}
}

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

decode  : byte -> text = convert bytecode to text
read    : text -> byte = convert text to bytecode
reduce  : byte -> byte = rewrite bytecode
reify   : byte -> byte = convert code to syntax tree
reflect : byte -> byte = convert syntax tree to code
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
	case "decode":
		stdin := bufio.NewReader(os.Stdin)
		block, err := abc.Decode(stdin)
		if err != nil {
			panic(err)
		}
		fmt.Println(block)
	case "read":
		stdin := bufio.NewReader(os.Stdin)
		block, err := abc.Read(stdin)
		if err != nil {
			panic(err)
		}
		stdout := bufio.NewWriter(os.Stdout)
		block.Encode(stdout)
		stdout.Flush()
	case "reduce":
		stdin := bufio.NewReader(os.Stdin)
		lhs, err := abc.Decode(stdin)
		if err != nil {
			panic(err)
		}
		stdout := bufio.NewWriter(os.Stdout)
		rhs := abc.Reduce(lhs, defaultQuota)
		rhs.Encode(stdout)
		stdout.Flush()
	case "reify":
		panic("unimplemented")
	case "reflect":
		panic("unimplemented")
	default:
		usage()
	}
}

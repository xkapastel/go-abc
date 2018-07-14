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
	"strings"
)

func main() {
	ctx := abc.NewBuild()
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("abc> ")
	for stdin.Scan() {
		str := stdin.Text()
		buf := strings.Split(str, " ")
		for _, word := range buf {
			switch word {
			case "opid":
				ctx.OpId()
			case "opapp":
				ctx.OpApp()
			case "opbox":
				ctx.OpBox()
			case "opcat":
				ctx.OpCat()
			case "opcopy":
				ctx.OpCopy()
			case "opdrop":
				ctx.OpDrop()
			case "opswap":
				ctx.OpSwap()
			case "mkbox":
				ctx.MkBox()
			case "mkcat":
				ctx.MkCat()
			case "mklink":
				ctx.MkLink()
			case "rmbox":
				ctx.RmBox()
			case "rmcat":
				ctx.RmCat()
			case "rmlink":
				ctx.RmLink()
			case "copy":
				ctx.Copy()
			case "drop":
				ctx.Drop()
			case "swap":
				ctx.Swap()
			case "reduce":
				ctx.Reduce()
			case "quit":
				block := ctx.Block()
				stdout := bufio.NewWriter(os.Stdout)
				if err := block.Encode(stdout); err != nil {
					panic(err)
				}
				stdout.Flush()
				return
			default:
				fmt.Printf("Unknown command: %s\n", word)
			}
		}
		fmt.Println(ctx)
		fmt.Printf("abc> ")
	}
}

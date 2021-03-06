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
	"github.com/xkapastel/go-abc/pkg/abc"
	"os"
)

func main() {
	const defaultQuota = 1000
	stdin := bufio.NewReader(os.Stdin)
	lhs, err := abc.Read(stdin)
	if err != nil {
		panic(err)
	}
	rhs := abc.Rewrite(lhs, defaultQuota)
	fmt.Println(rhs)
}

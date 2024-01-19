// Copyright (c) Alisdair MacLeod <copying@alisdairmacleod.co.uk>
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
// OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(&analysis.Analyzer{
		Name: "elsinator",
		Doc:  "finds else statements",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			for _, f := range pass.Files {
				ast.Inspect(f, func(n ast.Node) bool {
					ifstmt, ok := n.(*ast.IfStmt)
					if !ok {
						return true
					}

					if ifstmt.Else == nil {
						return true
					}

					pass.Reportf(ifstmt.Else.Pos(), "else statement found")
					return true
				})
			}
			return nil, nil
		},
	})
}

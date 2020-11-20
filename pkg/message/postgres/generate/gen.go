// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"

	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/message/postgres/schema"
)

func main() {
	qg.Generate(
		"sql-generated.go",
		[]string{"time"},
		[]*qg.GenTask{
			{
				QueryName: "selectMessages",
				SQL: qb.Select(
					qb.SelectQuery{
						Cols:    schema.SelectCols,
						From:    schema.Table,
						Joins:   nil,
						Where:   nil,
						GroupBy: []string{},
						Having:  nil,
						OrderBy: []string{fmt.Sprintf("%s DESC", schema.ColCreatedAt)},
						Limit:   "$1",
						Offset:  "$2",
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "limit",
						Type: "int",
					},
					{
						Name: "offset",
						Type: "int",
					},
				},
			},
			{
				QueryName: "selectMessageByID",
				SQL: qb.Select(
					qb.SelectQuery{
						Cols:  schema.SelectCols,
						From:  schema.Table,
						Joins: nil,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColID),
						),
						GroupBy: []string{},
						Having:  nil,
						OrderBy: []string{},
						Limit:   "",
						Offset:  "",
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
				},
			},
			{
				QueryName: "deleteMessage",
				SQL: qb.Delete(
					qb.DeleteQuery{
						From: schema.Table,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColID),
						)},
				),
				QueryType: qg.Exec,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
				},
			},
			{
				QueryName: "insertMessage",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: schema.Table,
						Cols: []string{
							schema.ColUID,
							schema.ColText,
							schema.ColCreatedByUserID,
							schema.ColCreatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4",
						},
						Returning: []string{schema.ColID},
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "UID",
						Type: "string",
					},
					{
						Name: "text",
						Type: "string",
					},
					{
						Name: "CreatedByUserID",
						Type: "int",
					},
					{
						Name: "createdAt",
						Type: "time.Time",
					},
				},
			},
		})
}

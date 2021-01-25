// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"

	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/node/postgres/schema"
	userSchema "github.com/dwethmar/atami/pkg/user/postgres/schema"
)

var IDCol = schema.WithTbl(schema.ColID)
var UIDCol = schema.WithTbl(schema.ColUID)

var defaultCols = append(
	schema.SelectCols,
	userSchema.WithTbl(userSchema.ColID),
	userSchema.WithTbl(userSchema.ColUID),
	userSchema.WithTbl(userSchema.ColUsername),
)

var defaultJoin = qb.NewJoin().
	Left(fmt.Sprintf(
		"%s ON %v = %v",
		userSchema.Table,
		schema.WithTbl(schema.ColCreatedByUserID),
		userSchema.WithTbl(userSchema.ColID),
	))

var defaultOrderBy = []string{fmt.Sprintf("%s DESC", schema.WithTbl(schema.ColCreatedAt))}

func main() {
	qg.Generate(
		"generated--queries.go",
		"postgres",
		[]string{
			"time",
			"github.com/dwethmar/atami/pkg/node",
		},
		[]*qg.GenerateQuery{
			{
				Name: "selectNodes",
				SQL: qb.Select(
					qb.SelectQuery{
						Select:  defaultCols,
						From:    schema.Table,
						Joins:   defaultJoin,
						Where:   nil,
						GroupBy: []string{},
						Having:  nil,
						OrderBy: defaultOrderBy,
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
				MapFunc:    "mapNodeWithUser",
				ReturnType: "*node.Node",
			},
			{
				Name: "selectNodeByID",
				SQL: qb.Select(
					qb.SelectQuery{
						Select: defaultCols,
						From:   schema.Table,
						Joins:  defaultJoin,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", IDCol),
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
				MapFunc:    "mapNodeWithUser",
				ReturnType: "*node.Node",
			},
			{
				Name: "selectNodeByUID",
				SQL: qb.Select(
					qb.SelectQuery{
						Select: defaultCols,
						From:   schema.Table,
						Joins:  defaultJoin,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", UIDCol),
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
						Name: "UID",
						Type: "string",
					},
				},
				MapFunc:    "mapNodeWithUser",
				ReturnType: "*node.Node",
			},
			{
				Name: "deleteNode",
				SQL: qb.Delete(
					qb.DeleteQuery{
						From: schema.Table,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", IDCol),
						)},
				),
				QueryType: qg.Exec,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*node.Node",
			},
			{
				Name: "insertNode",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: schema.Table,
						Select: []string{
							schema.ColUID,
							schema.ColText,
							schema.ColCreatedByUserID,
							schema.ColCreatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4",
						},
						Returning: schema.SelectCols,
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
				MapFunc:    "defaultMap",
				ReturnType: "*node.Node",
			},
		})
}

// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"

	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/domain/entity/message/schema"
	userSchema "github.com/dwethmar/atami/pkg/domain/entity/user/schema"
)

var IDCol = schema.WithTbl(schema.ColID)
var UIDCol = schema.WithTbl(schema.ColUID)

var defaultCols = append(
	schema.SelectCols,
	userSchema.WithTbl(userSchema.ColID),
	userSchema.WithTbl(userSchema.ColUID),
	userSchema.WithTbl(userSchema.ColUsername),
)

var defaultFrom = qb.From(schema.Table)
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
		"postgres_queries.go",
		"message",
		[]string{
			"time",
			"github.com/dwethmar/atami/pkg/domain/entity",
		},
		[]*qg.GenerateQuery{
			{
				Name: "selectMessages",
				SQL: qb.Select(
					qb.SelectQuery{
						SelectCols: defaultCols,
						From:       defaultFrom,
						Joins:      defaultJoin,
						Where:      nil,
						GroupBy:    []string{},
						Having:     nil,
						OrderBy:    defaultOrderBy,
						Limit:      "$1",
						Offset:     "$2",
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "limit",
						Type: "uint",
					},
					{
						Name: "offset",
						Type: "uint",
					},
				},
				MapFunc:    "mapMessageWithUser",
				ReturnType: "*Message",
			},
			{
				Name: "selectMessageByID",
				SQL: qb.Select(
					qb.SelectQuery{
						SelectCols: defaultCols,
						From:       defaultFrom,
						Joins:      defaultJoin,
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
				MapFunc:    "mapMessageWithUser",
				ReturnType: "*Message",
			},
			{
				Name: "selectMessageByUID",
				SQL: qb.Select(
					qb.SelectQuery{
						SelectCols: defaultCols,
						From:       defaultFrom,
						Joins:      defaultJoin,
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
				MapFunc:    "mapMessageWithUser",
				ReturnType: "*Message",
			},
			{
				Name: "deleteMessage",
				SQL: qb.Delete(
					qb.DeleteQuery{
						From: defaultFrom,
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
				ReturnType: "*Message",
			},
			{
				Name: "insertMessage",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: schema.Table,
						InsertCols: []string{
							schema.ColUID,
							schema.ColText,
							schema.ColCreatedByUserID,
							schema.ColCreatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4",
						},
						Returning: []string{
							schema.ColID,
						},
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
				MapFunc:    "insertMap",
				ReturnType: "entity.ID",
			},
		})
}

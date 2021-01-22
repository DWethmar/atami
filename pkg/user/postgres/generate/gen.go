package main

import (
	"fmt"
	"strconv"

	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/user/postgres/schema"
)

var defaultCols = schema.SelectCols

func main() {
	qg.Generate(
		"sql-generated.go",
		"postgres",
		[]string{
			"time",
			"github.com/dwethmar/atami/pkg/user",
		},
		[]*qg.GenerateQuery{
			{
				Name: "selectUsernameUniqueCheck",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: []string{"1"},
						Where: qb.NewWhere().
							And(fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColUsername))),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "username",
						Type: "string",
					},
				},
				MapFunc:    "mapIsUniqueCheck",
				ReturnType: "bool",
			},

			{
				Name: "selectEmailUniqueCheck",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: []string{"1"},
						Where: qb.NewWhere().
							And(fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColEmail))),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "email",
						Type: "string",
					},
				},
				MapFunc:    "mapIsUniqueCheck",
				ReturnType: "bool",
			},

			{
				Name: "insertUser",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: schema.Table,
						Select: []string{
							schema.ColUID,
							schema.ColUsername,
							schema.ColEmail,
							schema.ColPassword,
							schema.ColCreatedAt,
							schema.ColUpdatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4", "$5", "$6",
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
						Name: "username",
						Type: "string",
					},
					{
						Name: "email",
						Type: "string",
					},
					{
						Name: "password",
						Type: "string",
					},
					{
						Name: "createdAt",
						Type: "time.Time",
					},
					{
						Name: "updateddAt",
						Type: "time.Time",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},

			{
				Name: "deleteUser",
				SQL: qb.Delete(
					qb.DeleteQuery{
						From:  schema.Table,
						Where: qb.NewWhere().And(fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColID))),
					},
				),
				QueryType: qg.Exec,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},

			{
				Name: "selectUsers",
				SQL: qb.Select(
					qb.SelectQuery{
						Select:  defaultCols,
						From:    schema.Table,
						OrderBy: []string{fmt.Sprintf("%s ASC", schema.WithTbl(schema.ColCreatedAt))},
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
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},

			{
				Name: "selectUserByID",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColID)),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},

			{
				Name: "selectUserByUID",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColUID)),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "UID",
						Type: "string",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},

			{
				Name: "selectUserByEmail",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColEmail)),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "email",
						Type: "string",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},

			{
				Name: "selectUserByEmailWithPassword",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: append(schema.SelectCols, schema.WithTbl(schema.ColPassword)),
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColEmail)),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "email",
						Type: "string",
					},
				},
				MapFunc:    "mapWithPassword",
				ReturnType: "*user.User",
			},

			{
				Name: "selectUserByUsername",
				SQL: qb.Select(
					qb.SelectQuery{
						From:   schema.Table,
						Select: schema.SelectCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.WithTbl(schema.ColUsername)),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "username",
						Type: "string",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},
		})
}

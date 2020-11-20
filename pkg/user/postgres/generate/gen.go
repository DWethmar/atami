// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"
	"strconv"

	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/user/postgres/schema"
)

func main() {
	qg.Generate(
		"sql-generated.go",
		[]string{"time"},
		[]*qg.GenTask{

			{
				QueryName: "selectUsernameUniqueCheck",
				SQL: qb.Select(
					qb.SelectQuery{
						From:  schema.Table,
						Cols:  []string{"1"},
						Where: qb.NewWhere().And(fmt.Sprintf("%s = $1", schema.ColUsername)),
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
			},

			{
				QueryName: "selectEmailUniqueCheck",
				SQL: qb.Select(
					qb.SelectQuery{
						From:  schema.Table,
						Cols:  []string{"1"},
						Where: qb.NewWhere().And(fmt.Sprintf("%s = $1", schema.ColEmail)),
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
			},

			{
				QueryName: "insertUser",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: schema.Table,
						Cols: []string{
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
			},

			{
				QueryName: "deleteUser",
				SQL: qb.Delete(
					qb.DeleteQuery{
						From:  schema.Table,
						Where: qb.NewWhere().And(fmt.Sprintf("%s = $1", schema.ColID)),
					},
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
				QueryName: "selectUsers",
				SQL: qb.Select(
					qb.SelectQuery{
						Cols:    schema.SelectCols,
						From:    schema.Table,
						OrderBy: []string{"created_at ASC"},
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
				QueryName: "selectUserByID",
				SQL: qb.Select(
					qb.SelectQuery{
						From: schema.Table,
						Cols: schema.SelectCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColID),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
				},
			},

			{
				QueryName: "selectUserByUID",
				SQL: qb.Select(
					qb.SelectQuery{
						From: schema.Table,
						Cols: schema.SelectCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColUID),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "UID",
						Type: "string",
					},
				},
			},

			{
				QueryName: "selectUserByEmail",
				SQL: qb.Select(
					qb.SelectQuery{
						From: schema.Table,
						Cols: schema.SelectCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColEmail),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "email",
						Type: "string",
					},
				},
			},

			{
				QueryName: "selectUserByEmailWithPassword",
				SQL: qb.Select(
					qb.SelectQuery{
						From: schema.Table,
						Cols: append(schema.SelectCols, schema.ColPassword),
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColEmail),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "email",
						Type: "string",
					},
				},
			},

			{
				QueryName: "selectUserByUsername",
				SQL: qb.Select(
					qb.SelectQuery{
						From: schema.Table,
						Cols: schema.SelectCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", schema.ColUsername),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.Query,
				FuncArgs: []qg.FuncArg{
					{
						Name: "username",
						Type: "string",
					},
				},
			},
		})
}

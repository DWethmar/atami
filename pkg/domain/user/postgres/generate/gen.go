package main

import (
	"fmt"
	"strconv"

	"github.com/dwethmar/atami/pkg/database/querybuilder"
	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/domain/user/postgres/schema"
)

var defaultCols = schema.SelectCols
var defaultFrom = querybuilder.From(schema.Table);

var usernameCol = schema.WithTbl(schema.ColUsername)
var biographyCol = schema.WithTbl(schema.ColBiography)
var emailCol = schema.WithTbl(schema.ColEmail)
var idCol = schema.WithTbl(schema.ColID)
var createdAtCol = schema.WithTbl(schema.ColCreatedAt)
var updatedAtCol = schema.WithTbl(schema.ColUpdatedAt)
var uidCol = schema.WithTbl(schema.ColUID)
var passwordCol = schema.WithTbl(schema.ColPassword)

func main() {
	qg.Generate(
		"generated--queries.go",
		"postgres",
		[]string{
			"time",
			"github.com/dwethmar/atami/pkg/domain/user",
		},
		[]*qg.GenerateQuery{
			{
				Name: "selectUsernameUniqueCheck",
				SQL: qb.Select(
					qb.SelectQuery{
						From: defaultFrom,
						SelectCols: []string{"1"},
						Where: qb.NewWhere().
							And(fmt.Sprintf("%s = $1", usernameCol)),
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
						From: defaultFrom,
						SelectCols: []string{"1"},
						Where: qb.NewWhere().
							And(fmt.Sprintf("%s = $1", emailCol)),
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
						InsertCols: []string{
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
						From:  defaultFrom,
						Where: qb.NewWhere().And(fmt.Sprintf("%s = $1", idCol)),
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
						SelectCols:  defaultCols,
						From: defaultFrom,
						OrderBy: []string{fmt.Sprintf("%s ASC", createdAtCol)},
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
						From: defaultFrom,
						SelectCols: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", idCol),
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
						From: defaultFrom,
						SelectCols: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", uidCol),
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
						From: defaultFrom,
						SelectCols: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", emailCol),
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
						From: defaultFrom,
						SelectCols: append(schema.SelectCols, passwordCol),
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", emailCol),
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
						From: defaultFrom,
						SelectCols: schema.SelectCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", usernameCol),
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

			{
				Name: "updateUser",
				SQL: qb.Update(
					qb.UpdateQuery{
						Table: schema.Table,
						Set: []qb.UpdateCol{
							{Name: schema.ColBiography, Value: "$2"},
							{Name: schema.ColUpdatedAt, Value: "$3"},
						},
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", idCol),
						),
						Returning: schema.SelectCols,
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: "ID",
						Type: "int",
					},
					{
						Name: "biography",
						Type: "string",
					},
					{
						Name: "updatedAt",
						Type: "time.Time",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*user.User",
			},
		})
}

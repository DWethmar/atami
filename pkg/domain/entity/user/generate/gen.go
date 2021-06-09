package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dwethmar/atami/pkg/database/querybuilder"
	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	qg "github.com/dwethmar/atami/pkg/database/querygenerator"
	"github.com/dwethmar/atami/pkg/domain/user/postgres/schema"
	"github.com/iancoleman/strcase"
)

var defaultCols = schema.SelectCols
var defaultFrom = querybuilder.From(schema.Table)

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
		"sql.go",
		"user",
		[]string{
			"time",
			"github.com/dwethmar/atami/pkg/domain/entity",
		},
		[]*qg.GenerateQuery{
			{
				Name: "insertUser",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: schema.Table,
						InsertCols: []string{
							schema.ColUID,
							schema.ColUsername,
							schema.ColBiography,
							schema.ColEmail,
							schema.ColPassword,
							schema.ColCreatedAt,
							schema.ColUpdatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4", "$5", "$6", "$7",
						},
						Returning: schema.SelectCols,
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: strings.ToUpper(schema.ColUID),
						Type: "string",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColUsername),
						Type: "string",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColBiography),
						Type: "string",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColEmail),
						Type: "string",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColPassword),
						Type: "string",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColCreatedAt),
						Type: "time.Time",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColUpdatedAt),
						Type: "time.Time",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
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
						Name: strings.ToUpper(schema.ColID),
						Type: "entity.ID",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
			},

			{
				Name: "selectUsers",
				SQL: qb.Select(
					qb.SelectQuery{
						SelectCols: defaultCols,
						From:       defaultFrom,
						OrderBy:    []string{fmt.Sprintf("%s DESC", createdAtCol)},
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
				MapFunc:    "defaultMap",
				ReturnType: "*User",
			},

			{
				Name: "selectUserByID",
				SQL: qb.Select(
					qb.SelectQuery{
						From:       defaultFrom,
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
						Name: strings.ToUpper(schema.ColID),
						Type: "entity.ID",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
			},

			{
				Name: "selectUserByUID",
				SQL: qb.Select(
					qb.SelectQuery{
						From:       defaultFrom,
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
						Name: strings.ToUpper(schema.ColUID),
						Type: "entity.UID",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
			},

			{
				Name: "selectUserByEmail",
				SQL: qb.Select(
					qb.SelectQuery{
						From:       defaultFrom,
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
						Name: strcase.ToLowerCamel(schema.ColEmail),
						Type: "string",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
			},

			{
				Name: "selectUserCredentials",
				SQL: qb.Select(
					qb.SelectQuery{
						From:       defaultFrom,
						SelectCols: []string{usernameCol, emailCol, passwordCol},
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", emailCol),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: strcase.ToLowerCamel(schema.ColEmail),
						Type: "string",
					},
				},
				MapFunc:    "mapCredentials",
				ReturnType: "*UserCredentials",
			},

			{
				Name: "selectUserByUsername",
				SQL: qb.Select(
					qb.SelectQuery{
						From:       defaultFrom,
						SelectCols: defaultCols,
						Where: qb.NewWhere().And(
							fmt.Sprintf("%s = $1", usernameCol),
						),
						Limit: strconv.Itoa(1),
					},
				),
				QueryType: qg.QueryRow,
				FuncArgs: []qg.FuncArg{
					{
						Name: strcase.ToLowerCamel(schema.ColUsername),
						Type: "string",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
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
						Name: strings.ToUpper(schema.ColID),
						Type: "entity.ID",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColBiography),
						Type: "string",
					},
					{
						Name: strcase.ToLowerCamel(schema.ColUpdatedAt),
						Type: "time.Time",
					},
				},
				MapFunc:    "defaultMap",
				ReturnType: "*User",
			},
		})
}

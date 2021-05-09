// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"

	qb "github.com/dwethmar/atami/pkg/database/querybuilder"
	sg "github.com/dwethmar/atami/pkg/database/seedergenerator"

	"github.com/dwethmar/atami/pkg/domain/entity/message/schema"
	messageSchema "github.com/dwethmar/atami/pkg/domain/entity/message/schema"
	userSchema "github.com/dwethmar/atami/pkg/domain/entity/user/schema"
)

var defaultOrderBy = []string{fmt.Sprintf("%s DESC", schema.WithTbl(schema.ColCreatedAt))}

func main() {
	sg.Generate(
		"seed_queries.go",
		"seed",
		[]string{
			"time",
			"github.com/dwethmar/atami/pkg/domain/entity",
		},
		[]*sg.Seeder{
			{
				Name: "User",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: userSchema.Table,
						InsertCols: []string{
							userSchema.ColUID,
							userSchema.ColUsername,
							userSchema.ColEmail,
							userSchema.ColBiography,
							userSchema.ColCreatedAt,
							userSchema.ColUpdatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4", "$5", "$6",
						},
					},
				),
				FuncArgs: []sg.FuncArg{
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
				Name: "Message",
				SQL: qb.Insert(
					qb.InsertQuery{
						Into: messageSchema.Table,
						InsertCols: []string{
							messageSchema.ColUID,
							messageSchema.ColText,
							messageSchema.ColCreatedByUserID,
							messageSchema.ColCreatedAt,
							messageSchema.ColUpdatedAt,
						},
						Values: []interface{}{
							"$1", "$2", "$3", "$4", "$5",
						},
					},
				),
				FuncArgs: []sg.FuncArg{
					{
						Name: "UID",
						Type: "string",
					},
					{
						Name: "text",
						Type: "string",
					},
					{
						Name: "createdByUserID",
						Type: "entity.ID",
					},
					{
						Name: "createdAt",
						Type: "time.Time",
					},
					{
						Name: "updatedAt",
						Type: "time.Time",
					},
				},
			},
		})
}

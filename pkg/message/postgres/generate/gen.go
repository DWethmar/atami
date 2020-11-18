// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/dwethmar/atami/pkg/database/querybuilder"
	"github.com/dwethmar/atami/pkg/message/postgres/schema"
)

var queries = []struct {
	Name  string
	Query string
}{
	{
		"selectMessages",
		querybuilder.Select(
			querybuilder.SelectQuery{
				Cols:      schema.SelectCols,
				From:      schema.Table,
				Joins:     nil,
				Where:     nil,
				GroupBy:   []string{},
				Having:    nil,
				OrderBy:   []string{fmt.Sprintf("%s DESC", schema.ColCreatedAt)},
				LimitStr:  "$1",
				OffsetStr: "$2",
			},
		),
	},
	{
		"selectMessageByID",
		querybuilder.Select(
			querybuilder.SelectQuery{
				Cols:    schema.SelectCols,
				From:    schema.Table,
				Joins:   nil,
				Where:   querybuilder.NewWhere().And("id = $1"),
				GroupBy: []string{},
				Having:  nil,
				OrderBy: []string{},
				Limit:   0,
				Offset:  0,
			},
		),
	},
	{
		"insertMessage",
		querybuilder.Insert(
			querybuilder.InsertQuery{
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
			},
		),
	},
	{
		"deleteMessage",
		querybuilder.Delete(
			querybuilder.DeleteQuery{
				From:  schema.Table,
				Where: querybuilder.NewWhere().And("id = $1"),
			},
		),
	},
}

func main() {
	f, err := os.Create("sql.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp time.Time
		Queries   []struct {
			Name  string
			Query string
		}
	}{
		Timestamp: time.Now(),
		Queries:   queries,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// This file was generated by robots; DO NOT EDIT.
// run: 'make generate' to regenerate this file.

package postgres
{{ range .Queries }}
// {{ .Name }} sql query
var {{ .Name }} = ` + "`" + `{{ .Query }}` + "`" + `
{{ end }}`))

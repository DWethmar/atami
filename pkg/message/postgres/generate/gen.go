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

	"github.com/dwethmar/atami/pkg/message/postgres/schema"
	"github.com/dwethmar/atami/pkg/postgres"
)

var queries = map[string]string{
	"getMessages": postgres.Select(
		postgres.SelectQuery{
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
	"getMessageByID": postgres.Select(
		postgres.SelectQuery{
			Cols:      schema.SelectCols,
			From:      schema.Table,
			Joins:     nil,
			Where:     postgres.NewWhere().And("id = $1"),
			GroupBy:   []string{},
			Having:    nil,
			OrderBy:   []string{fmt.Sprintf("%s DESC", schema.ColCreatedAt)},
			LimitStr:  "$1",
			OffsetStr: "$2",
		},
	),
}

func main() {
	f, err := os.Create("sql.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp time.Time
		Queries   map[string]string
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

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package postgres
{{ range $key, $value := .Queries }}
// {{ $key }} sql query
var {{ $key }} = ` + "`" + `{{ $value }}` + "`" + `
{{ end }}


`))

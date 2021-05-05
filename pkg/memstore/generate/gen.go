package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"strings"
)

var storeTemplate = template.Must(
	template.New("").
		Funcs(template.FuncMap{
			"Title": strings.Title,
		}).
		Parse(storeTemplateText),
)

func main() {
	wordPtr := flag.String("name", "name", "a string")
	fileOut := flag.String("fileOut", fmt.Sprintf("%s.go", *wordPtr), "out file")

	flag.Parse()

	f := strings.Trim(*fileOut, "\"")

	fmt.Println(*wordPtr)
	fmt.Println(f)

	GenerateStore(f, "memstore", *wordPtr)
}

// GenerateStore queries to file
func GenerateStore(fileOut string, packageName string, name string) {
	f, err := os.Create(fileOut)
	die(err)
	defer f.Close()

	var buf bytes.Buffer
	storeTemplate.Execute(&buf, struct {
		PackageName string
		Name        string
	}{
		PackageName: packageName,
		Name:        name,
	})

	p, err := format.Source(buf.Bytes())
	die(err)
	f.Write(p)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

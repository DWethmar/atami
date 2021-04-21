package querybuilder

import (
	"fmt"
	"strings"
)

type SelectCols []string

func (s SelectCols) String() string {
	var selectPart = ""
	if len(s) > 0 {
		selectPart = fmt.Sprintf(`SELECT%s`, "\n\t"+strings.Join(s, ",\n\t"))
	} else {
		selectPart =  "SELECT *"
	}
	return selectPart
}

type InsertCols []string

func (i InsertCols) String() string {
	return fmt.Sprintf(`(%s)`, "\n\t"+strings.Join(i, ",\n\t")+"\n")
}

type From string 

func (f From) String() string {
	return fmt.Sprintf(`FROM %s`, string(f))
}
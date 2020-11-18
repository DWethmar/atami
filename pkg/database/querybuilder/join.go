package querybuilder

import (
	"fmt"
	"strings"
)

type joinStatement struct {
	joinType  string
	predicate string
}

// Join clause of sql statement
// (INNER) JOIN: Returns records that have matching values in both tables
// LEFT (OUTER) JOIN: Returns all records from the left table, and the matched records from the right table
// RIGHT (OUTER) JOIN: Returns all records from the right table, and the matched records from the left table
// FULL (OUTER) JOIN: Returns all records when there is a match in either left or right table
type Join struct {
	statements []joinStatement
}

// Inner adds inner join
// example: .Inner("user ON user.id = other.user_id")
func (j *Join) Inner(predicate string) *Join {
	j.statements = append(j.statements, joinStatement{
		"INNER JOIN",
		predicate,
	})
	return j
}

// Left adds left join
// example: .LEft("user ON user.id = other.user_id")
func (j *Join) Left(predicate string) *Join {
	j.statements = append(j.statements, joinStatement{
		"LEFT JOIN",
		predicate,
	})
	return j
}

// Right adds right join
// example: .Right("user ON user.id = other.user_id")
func (j *Join) Right(predicate string) *Join {
	j.statements = append(j.statements, joinStatement{
		"RIGHT JOIN",
		predicate,
	})
	return j
}

// Full adds full join
// example: .Full("user ON user.id = other.user_id")
func (j *Join) Full(predicate string) *Join {
	j.statements = append(j.statements, joinStatement{
		"FULL JOIN",
		predicate,
	})
	return j
}

// String to string
func (j *Join) String() string {
	join := make([]string, len(j.statements))

	for i, c := range j.statements {
		join[i] = fmt.Sprintf("%s %s", c.joinType, c.predicate)

	}

	return strings.Join(join, "\n")
}

// NewJoin create new where
func NewJoin() *Join {
	return &Join{
		statements: make([]joinStatement, 0),
	}
}

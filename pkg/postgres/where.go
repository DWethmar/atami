package postgres

import "strings"

type condition struct {
	operator  string
	predicate string
}

// Where clause of sql statement
type Where struct {
	conditions []condition
}

// And adds and condition
func (w *Where) And(predicate string) {
	var operator string
	if len(w.conditions) > 0 {
		operator = "AND"
	}
	w.conditions = append(w.conditions, condition{
		operator,
		predicate,
	})
}

// Or adds or condition
func (w *Where) Or(predicate string) {
	var operator string
	if len(w.conditions) > 0 {
		operator = "OR"
	}
	w.conditions = append(w.conditions, condition{
		operator,
		predicate,
	})
}

// String to string
func (w *Where) String() string {
	where := make([]string, len(w.conditions))

	for i, c := range w.conditions {
		if c.operator == "" {
			where[i] = c.predicate
		} else {
			where[i] = c.operator + " " + c.predicate
		}
	}

	return strings.Join(where, "\n")
}

// NewWhere create new where
func NewWhere() *Where {
	return &Where{
		conditions: make([]condition, 0),
	}
}

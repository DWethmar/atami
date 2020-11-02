package query

type Condition struct {
	Field    string
	Operator string
	value    interface{}
}

type Query struct {
	conditions []*Condition
	And        bool
}

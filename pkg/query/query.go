package query

var (
	EQ = "EQ"
)

type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}

type Query struct {
	Conditions []*Condition
	And        bool
}

type QueryError struct {
	Query string
	Err   error
}

func (e QueryError) Error() string {
	return e.Query + ": " + e.Err.Error()
}

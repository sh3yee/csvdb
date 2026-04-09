package query

type Query struct {
	path string
}

type Condition struct {
	Column string
	Op     string // "=", "!=", ">", "<", ">=", "<=", "like"
	Value  string
}

func New(path string) *Query {
	return &Query{
		path: path,
	}
}

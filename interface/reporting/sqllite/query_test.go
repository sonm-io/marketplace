package sqllite

type unknownQuery struct{}

func (c unknownQuery) QueryID() string {
	return "UnknownQuery"
}
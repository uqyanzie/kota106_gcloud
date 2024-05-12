package model

type Request struct {
	projection []string
	connection Connection
	conditions []Condition
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) SetProjection(projection []string) {
	r.projection = projection
}

func (r *Request) SetConnection(connection Connection) {
	r.connection = connection
}

func (r *Request) SetConditions(conditions []Condition) {
	r.conditions = conditions
}

func (r *Request) GetProjection() []string {
	return r.projection
}

func (r *Request) GetConnection() Connection {
	return r.connection
}

func (r *Request) GetConditions() []Condition {
	return r.conditions
}

package model

type Condition struct {
	column   string
	operator Operator
	value    interface{}
}

func NewCondition() *Condition {
	return &Condition{}
}

func (c *Condition) SetColumn(column string) {
	c.column = column
}

func (c *Condition) SetOperator(operator Operator) {
	c.operator = operator
}

func (c *Condition) SetValue(value interface{}) {
	c.value = value
}

func (c *Condition) GetColumn() string {
	return c.column
}

func (c *Condition) GetOperator() Operator {
	return c.operator
}

func (c *Condition) GetValue() interface{} {
	return c.value
}

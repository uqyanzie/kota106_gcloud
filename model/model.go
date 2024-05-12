package model

type Operator int

type Connection int

const (
	EQUAL Operator = iota
	NOT_EQUAL
	GREATER_THAN
	LESS_THAN
	GREATER_THAN_OR_EQUAL
	LESS_THAN_OR_EQUAL
)

const (
	AND Connection = iota
	OR
)

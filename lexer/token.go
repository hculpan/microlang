package lexer

import "fmt"

type Token interface {
	String() string
}

type StringToken struct {
	Value    string
	Line     int
	Position int
}

func (s StringToken) String() string {
	return fmt.Sprintf("StringToken(%s) [Line %d, Position %d]", s.Value, s.Line, s.Position)
}

type IntegerToken struct {
	Value    int64
	Line     int
	Position int
}

func (n IntegerToken) String() string {
	return fmt.Sprintf("IntegerToken(%d) [Line %d, Position %d]", n.Value, n.Line, n.Position)
}

type FloatToken struct {
	Value    float64
	Line     int
	Position int
}

func (n FloatToken) String() string {
	return fmt.Sprintf("FloatToken(%f) [Line %d, Position %d]", n.Value, n.Line, n.Position)
}

type KeywordToken struct {
	Value    string
	Line     int
	Position int
}

func (k KeywordToken) String() string {
	return fmt.Sprintf("KeywordToken(%s) [Line %d, Position %d]", k.Value, k.Line, k.Position)
}

type OperatorToken struct {
	Value    string
	Line     int
	Position int
}

func (o OperatorToken) String() string {
	return fmt.Sprintf("OperatorToken(%s) [Line %d, Position %d]", o.Value, o.Line, o.Position)
}

type IdentifierToken struct {
	Value    string
	Line     int
	Position int
}

func (i IdentifierToken) String() string {
	return fmt.Sprintf("IdentifierToken(%s) [Line %d, Position %d]", i.Value, i.Line, i.Position)
}

type CommentToken struct {
	Value    string
	Line     int
	Position int
}

func (c CommentToken) String() string {
	return fmt.Sprintf("CommentToken(%s) [Line %d, Position %d]", c.Value, c.Line, c.Position)
}

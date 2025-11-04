package frontend

type NodeType string

const (
	ProgramNode NodeType = "Program"
	NumericLiteraNode NodeType = "NumericLiteral"
	IdentifierNode NodeType = "Identifier"
	BinaryExprNode NodeType = "BinaryExpr"
	
)
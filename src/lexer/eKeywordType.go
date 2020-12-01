package lexer

// ELuaKeywordType Lua关键字类型
type ELuaKeywordType int

const (
	KwAnd ELuaKeywordType = iota
	KwBreak
	KwDo
	KwElse
	KwElseIf
	KwEnd
	KwFalse
	KwFor
	KwFunction
	KwGoto
	KwIf
	KwIn
	KwLocal
	KwNil
	KwNot
	KwOr
	KwRepeat
	KwReturn
	KwThen
	KwTrue
	KwUntil
	KwWhile
)

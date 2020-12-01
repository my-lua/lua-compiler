package lexer

// ELuaKeyword Lua关键字枚举
type ELuaKeyword int

const (
	// KwMistake 错误的关键字，即不属于Lua关键字
	KwMistake ELuaKeyword = iota
	KwAnd
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

// getLuaKeywordTable 获取Lua关键字表
func getLuaKeywordTable() []string {
	return []string{
		"[mistake!]",
		"and",
		"break",
		"do",
		"else",
		"elseif",
		"end",
		"false",
		"for",
		"function",
		"goto",
		"if",
		"in",
		"local",
		"nil",
		"not",
		"or",
		"repeat",
		"return",
		"then",
		"true",
		"until",
		"while",
	}
}

// getLuaKeywordMap 获取Lua关键字映射
func getLuaKeywordMap() map[string]ELuaKeyword {
	kwdMap := map[string]ELuaKeyword{}
	for index, item := range getLuaKeywordTable() {
		kwdMap[item] = ELuaKeyword(index)
	}
	return kwdMap
}

// IsLuaKeyword 判断一个字符串是不是Lua关键字
func IsLuaKeyword(text string) bool {
	if getLuaKeywordMap()[text] != KwMistake {
		return true
	}
	return false
}

// Name 关键字名称
func (me ELuaKeyword) Name() string {
	return getLuaKeywordTable()[int(me)]
}

// NewLuaKeyword 构造函数
func NewLuaKeyword(text string) ELuaKeyword {
	if IsLuaKeyword(text) {
		return getLuaKeywordMap()[text]
	}
	panic("ELuaKeyword NewLuaKeyword: 文本不是Lua关键字")
}

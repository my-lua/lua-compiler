package lexer

import (
	"regexp"
	"strings"
)

func (me *LuaLexer) scan(re *regexp.Regexp) string {
	return ""
}

func (me *LuaLexer) scanShortString() string {
	return ""
}

// collectLongString 采集长字符串
func (me *LuaLexer) collectLongString(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		if me.chunkTopIsPairOfNewLine() {
			me.chunkNextN(2)
			me.chunkNewLine()
			result += "\n"
		} else if me.chunkTopIsNewLine() {
			me.chunkNext()
			me.chunkNewLine()
			result += "\n"
		} else {
			result += me.chunkNext()
		}
	}
	// 如果行首为换行符，则删除
	if strings.HasPrefix(result, "\n") {
		result = result[1:]
	}
	return result
}

// scanLongString 扫描长字符串
func (me *LuaLexer) scanLongString() string {
	// 检查长字符串头部括号
	matchs := me.ReLongStringOpeningBracket().FindStringSubmatch(me.chunk)
	if len(matchs) < 2 {
		panic("lexer scanLongString: 长字符串头部括号不匹配")
	}
	// 获取头部括号长度, 其中等号个数
	openingBracketLen, equalSignNum := len(matchs[0]), len(matchs[1])
	// 跳过头部括号
	me.chunkNextN(openingBracketLen)
	// 向后寻找相匹配的尾部括号
	indexs := me.ReLongStringClosingBracket(equalSignNum).FindStringIndex(me.chunk)
	if len(indexs) < 2 {
		panic("lexer scanLongString: 长字符串没有正确的尾部括号")
	}
	// 获取尾部括号开始位置
	closingBracketStartIndex := indexs[0]
	// 收集并处理长字符串
	result := me.collectLongString(closingBracketStartIndex)
	// 跳过尾部括号
	me.chunkNextN(openingBracketLen)
	return result
}

func (me *LuaLexer) scanNumber() string {
	return ""
}

func (me *LuaLexer) scanIdentifier() string {
	return ""
}

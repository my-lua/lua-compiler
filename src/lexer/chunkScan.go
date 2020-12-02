package lexer

import (
	"regexp"
	"strings"
)

// scan 扫描的基础方法
func (me *Chunk) scan(re *regexp.Regexp) string {
	if rst := re.FindString(me.Text()); rst != "" {
		me.NextN(len(rst))
		return rst
	}
	panic("Chunk scan: 没有扫描到匹配的字符串")
}

// ScanNumber 扫描数字
func (me *Chunk) ScanNumber() string {
	return me.scan(ReNumber())
}

// ScanIdentifier 扫描标识符
func (me *Chunk) ScanIdentifier() string {
	return me.scan(ReIdentifier())
}

// ScanShortString 扫描短字符串
func (me *Chunk) ScanShortString() string {
	shortStr := me.scan(ReShortStr())
	// 去除两边引号
	return shortStr[1 : len(shortStr)-1]
}

// collectLongString 采集长字符串
func (me *Chunk) collectLongString(n int) string {
	rst := ""
	for i := 0; i < n; i++ {
		if me.Top().IsPairOfNewLine() {
			me.NextN(2)
			me.NewLine()
			rst += "\n"
		} else if me.Top().IsNewLine() {
			me.Next()
			me.NewLine()
			rst += "\n"
		} else {
			rst += me.Next()
		}
	}
	// 如果行首为换行符，则删除
	if strings.HasPrefix(rst, "\n") {
		rst = rst[1:]
	}
	return rst
}

// ScanLongString 扫描长字符串
func (me *Chunk) ScanLongString() string {
	// 检查长字符串头部括号
	matchs := ReLongStringBracketStart().FindStringSubmatch(me.Text())
	if len(matchs) < 2 {
		panic("Chunk ScanLongString: 长字符串头部括号不匹配")
	}
	// 获取头部括号长度, 其中等号个数
	openingBracketLen, equalSignNum := len(matchs[0]), len(matchs[1])
	// 跳过头部括号
	me.NextN(openingBracketLen)
	// 向后寻找相匹配的尾部括号
	indexs := ReLongStringBracketEnd(equalSignNum).FindStringIndex(me.Text())
	if len(indexs) < 2 {
		panic("Chunk ScanLongString: 长字符串没有正确的尾部括号")
	}
	// 获取尾部括号开始位置
	closingBracketStartIndex := indexs[0]
	// 收集并处理长字符串
	result := me.collectLongString(closingBracketStartIndex)
	// 跳过尾部括号
	me.NextN(openingBracketLen)
	return result
}

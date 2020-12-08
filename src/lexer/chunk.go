package lexer

// Chunk 在这里被抽象封装成为一个待处理的文本序列，并且提供一系列的操作方法
type Chunk struct {
	text    string
	curLine int
	curChar int
}

// Text 当前待解析的文本
func (me *Chunk) Text() string {
	return me.text
}

// CurLine 当前行
func (me *Chunk) CurLine() int {
	return me.curLine
}

// CurChar 当前列
func (me *Chunk) CurChar() int {
	return me.curChar
}

// Top 待处理文本的头部，实际就是待处理的文本，绑定了一系列操作头部的方法
func (me *Chunk) Top() ChunkTop {
	return ChunkTop(me.text)
}

// IsEmpty Chunk是否为空
func (me *Chunk) IsEmpty() bool {
	return len(me.text) < 1
}

// Next 下移处理一个字符，返回当前字符
func (me *Chunk) Next() byte {
	top := me.Top().Char()
	me.text = me.text[1:]
	me.curChar++
	return top
}

// NextN 下移处理n个字符，返回字符数组
func (me *Chunk) NextN(n int) []byte {
	rst := []byte{}
	for i := 0; i < n; i++ {
		rst = append(rst, me.Next())
	}
	return rst
}

// NewLine 新的一行
func (me *Chunk) NewLine() {
	me.curLine++
	me.curChar = 1
}

// NewLineN 新的n行
func (me *Chunk) NewLineN(n int) {
	for i := 0; i < n; i++ {
		me.NewLine()
	}
}

// NewChunk 构造函数
func NewChunk(text string) *Chunk {
	return &Chunk{
		text:    text,
		curLine: 1,
		curChar: 1,
	}
}

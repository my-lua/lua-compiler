package lexer

// Chunk Chunk类型
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

// Top Chunk头部
func (me *Chunk) Top() ChunkTop {
	return ChunkTop(me.text)
}

// IsEmpty Chunk是否为空
func (me *Chunk) IsEmpty() bool {
	return len(me.text) < 1
}

// Next 下移处理一个字符
func (me *Chunk) Next() byte {
	top := me.Top().Char()
	me.text = me.text[1:]
	return top
}

// NextN 下移处理n个字符
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

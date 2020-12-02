package lexer

// Chunk 词法分析器待处理的Chunk
type Chunk string

// IsEmpty Chunk是否为空
func (me Chunk) IsEmpty() bool {
	return len(me) < 1
}

// Top 获取Chunk头部
func (me Chunk) Top() ChunkTop {
	return ChunkTop(me)
}

// Next s
func (me Chunk) Next() string {
	return ""
}

// NextN s
func (me Chunk) NextN(n int) string {
	return ""
}

func (me Chunk) NewLine() {

}

func (me Chunk) NewLineN(n int) {

}

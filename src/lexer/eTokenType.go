package lexer

// ETokenType 单词类型
type ETokenType int

const (
	TokenEof        ETokenType     = iota // end-of-file
	TokenVararg                           // ...
	TokenSepSemi                          // ;
	TokenSepComma                         // ,
	TokenSepDot                           // .
	TokenSepColon                         // :
	TokenSepLabel                         // ::
	TokenSepLParen                        // (
	TokenSepRParen                        // )
	TokenSepLBrack                        // [
	TokenSepRBrack                        // ]
	TokenSepLCurly                        // {
	TokenSepRCurly                        // }
	TokenOpAssign                         // =
	TokenOpMinus                          // - (sub or unm)
	TokenOpWave                           // ~ (bnot or bxor)
	TokenOpAdd                            // +
	TokenOpMul                            // *
	TokenOpDiv                            // /
	TokenOpIDiv                           // //
	TokenOpPow                            // ^
	TokenOpMod                            // %
	TokenOpBAnd                           // &
	TokenOpBOr                            // |
	TokenOpShr                            // >>
	TokenOpShl                            // <<
	TokenOpConcat                         // ..
	TokenOpLt                             // <
	TokenOpLe                             // <=
	TokenOpGt                             // >
	TokenOpGe                             // >=
	TokenOpEq                             // ==
	TokenOpNe                             // ~=
	TokenOpLen                            // #
	TokenOpAnd                            // and
	TokenOpOr                             // or
	TokenOpNot                            // not
	TokenKwBreak                          // break
	TokenKwDo                             // do
	TokenKwElse                           // else
	TokenKwElseIf                         // elseif
	TokenKwEnd                            // end
	TokenKwFalse                          // false
	TokenKwFor                            // for
	TokenKwFunction                       // function
	TokenKwGoto                           // goto
	TokenKwIf                             // if
	TokenKwIn                             // in
	TokenKwLocal                          // local
	TokenKwNil                            // nil
	TokenKwRepeat                         // repeat
	TokenKwReturn                         // return
	TokenKwThen                           // then
	TokenKwTrue                           // true
	TokenKwUntil                          // until
	TokenKwWhile                          // while
	TokenIdentifier                       // identifier
	TokenNumber                           // number literal
	TokenString                           // string literal
	TokenOpUnm      = TokenOpMinus        // unary minus
	TokenOpSub      = TokenOpMinus
	TokenOpBNot     = TokenOpWave
	TokenOpBXor     = TokenOpWave
)

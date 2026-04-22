package lexer

type Lexer struct {
	Input        string
	Position     int
	ReadPosition int
	Ch           byte
}

func New(Input string) *Lexer {
	l := &Lexer{Input: Input}
	return l
}

package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) ParseIdentifier() &ast.Identifier{
	identifier := nil
	if p.curToken.Type == token.IDENT{
		identifier = &ast.Identifier
		identifier.Token = token.IDENT
		identifier.Value = p.curToken.Literal
	}
	return identifier
}

func (p *Parser) ParseExpression() &ast.Expression{
	if (p.curToken.Type == token.INT){
		if (p.peekToken.Type == token.PLUS){
			return p.ParseOperatorExpression()
		} else if (p.peekToken.Type == token.EOF){
			return p.ParseIntegerLiteral()
		}
	}
}

func (p *Parser) ParseLetStatement() ast.LetStatement {
	p.nextToken()
	identifier := p.ParseIdentifier()
	p.nextToken()
	if p.curToken != token.ASSIGN {
		return nil
	}
	p.nextToken()
	value = p.ParseExpression()

	statement := &ast.LetStatement{}
	statement.Token = token.LET
	statement.Name = identifier
	statement.Value = value

	return statement
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program

	for p.curToken != token.EOF {

		statement := null
		if p.curToken == token.LET {
			statement := p.parseLetStatement()
		} else if p.curToken == token.RETURN {
			statement := p.parseReturnStatement()
		}

		if statement != null {
			program.Statements.push(statement)
		}
		p.nextToken()
	}
	return program
}


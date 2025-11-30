package parser

type Parser struct {
	source string
}

func NewParser() *Parser {
	return &Parser{source: ""}
}

func (p *Parser) SetSource(s string) {
	p.source = s
}

func (p *Parser) Parse() {

}

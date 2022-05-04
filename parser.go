package parser

import (
	"bufio"
	"io"
	"unicode"
)

const (
	t_section = iota
	t_entry_key
	t_entry_value
)

type FluentBitConfParser struct {
	reader *bufio.Reader
	Conf   *FluentBitConf
	token  int
}

func NewFluentBitConfParser(reader io.Reader) *FluentBitConfParser {
	return &FluentBitConfParser{
		reader: bufio.NewReader(reader),
		Conf: &FluentBitConf{
			Sections: []Section{},
		},
		token: t_section,
	}
}

func (parser *FluentBitConfParser) Parse() *FluentBitConf {
	var currSection *Section = nil
	var currKey string

	for {
		r, _, err := parser.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if currSection != nil {
					parser.Conf.Sections = append(parser.Conf.Sections, *currSection)
				}
				return parser.Conf
			}
			return parser.Conf
		}
		switch r {
		case '\n':
			continue
		case '[':
			// save last config item
			if currSection != nil {
				parser.Conf.Sections = append(parser.Conf.Sections, *currSection)
			}
			// create new config item
			currSection = &Section{
				Name:    "",
				Entries: []Entry{},
			}
			parser.token = t_section
		default:
			if unicode.IsSpace(r) {
				continue
			}

			strValue, _ := parser.parseString()
			switch parser.token {
			case t_section:
				currSection.Name = strValue
				parser.token = t_entry_key
			case t_entry_key:
				currKey = strValue
				parser.token = t_entry_value
			case t_entry_value:
				currSection.bindEntry(currKey, strValue)

				currKey = ""
				parser.token = t_entry_key
			}
		}

	}
}

func (parser *FluentBitConfParser) parseString() (string, error) {
	var val string = ""

	if err := parser.reader.UnreadRune(); err != nil {
		return "", err
	}
	for {
		r, _, err := parser.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return val, nil
			}
			return "", err
		}

		if parser.token == t_entry_key && unicode.IsSpace(r) {
			return val, nil
		}

		if parser.token == t_entry_value && r == '\n' {
			return val, nil
		}

		if parser.token == t_section && r == ']' {
			return val, nil
		}

		val = val + string(r)
	}
}

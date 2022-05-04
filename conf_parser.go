package fluentbit_conf_parser

import (
	"bufio"
	"io"
	"unicode"
)

const (
	SECTION = iota
	ENTRY_KEY
	ENTRY_VALUE
)

type FluentBitConfParser struct {
	reader        *bufio.Reader
	Configuration []Section
	token         int
}

func NewFbReader(reader io.Reader) *FluentBitConfParser {
	return &FluentBitConfParser{
		reader:        bufio.NewReader(reader),
		Configuration: []Section{},
		token:         SECTION,
	}
}

func (parser *FluentBitConfParser) Parse() *FluentBitConfParser {
	var currSection *Section = nil
	var currKey string

	for {
		r, _, err := parser.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if currSection != nil {
					parser.Configuration = append(parser.Configuration, *currSection)
				}
				return parser
			}
			return parser
		}
		switch r {
		case '\n':
			continue
		case '[':
			// save last config item
			if currSection != nil {
				parser.Configuration = append(parser.Configuration, *currSection)
			}
			// create new config item
			currSection = &Section{
				Name:    "",
				Entries: []Entry{},
			}
			parser.token = SECTION
		default:
			if unicode.IsSpace(r) {
				continue
			}

			strValue, _ := parser.parseString()
			switch parser.token {
			case SECTION:
				currSection.Name = strValue
				parser.token = ENTRY_KEY
			case ENTRY_KEY:
				currKey = strValue
				parser.token = ENTRY_VALUE
			case ENTRY_VALUE:
				currSection.BindEntry(currKey, strValue)

				currKey = ""
				parser.token = ENTRY_KEY
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

		if parser.token == ENTRY_KEY && unicode.IsSpace(r) {
			return val, nil
		}

		if parser.token == ENTRY_VALUE && r == '\n' {
			return val, nil
		}

		if parser.token == SECTION && r == ']' {
			return val, nil
		}

		val = val + string(r)
	}
}

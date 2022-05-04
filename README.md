# go-fluentbit-conf-parser

![github](https://img.shields.io/badge/go-1.17-blue.svg) ![github](https://img.shields.io/badge/fluentbit-v1.9-lightblue.svg) ![github](https://img.shields.io/badge/License-MIT-green.svg)

Go package for parsering [Fluentbit](https://fluentbit.io/) `.conf` configuration file.

> more info: [Fluentbit Configuration Document](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/format-schema)

## Install

```shell
go get -u github.com/stevedsun/go-fluentbit-conf-parser
```

## Usage

```go
package main

import (
	"fmt"
	"os"

	parser "github.com/stevedsun/go-fluentbit-conf-parser"
)

func main() {
	confFile, _ := os.Open("td-agent-bit.conf")
	conf := parser.NewFluentBitConfParser(confFile).Parse()
	for _, section := range conf.Sections {
		fmt.Printf("Section: %v \n", section.Name)
		for _, entry := range section.Entries {
			fmt.Printf("Entry: %v %v \n", entry.Key, entry.Value)
		}
	}
}

```

# Todo

- Support [Commands](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/commands)

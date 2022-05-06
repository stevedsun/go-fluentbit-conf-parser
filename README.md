# Fluent-Bit configuration file parser for Golang

![github](https://img.shields.io/badge/go-1.17-blue.svg) ![github](https://img.shields.io/badge/fluentbit-v1.9-lightblue.svg) ![github](https://img.shields.io/badge/License-MIT-green.svg)

Go package for parsering [Fluentbit](https://fluentbit.io/) `.conf` configuration file.

> Read more: [Fluentbit Configuration Document](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/format-schema)

## Features

- Support Section and Entry objects
- Support [Commands](https://docs.fluentbit.io/manual/administration/configuring-fluent-bit/classic-mode/commands)
- Export all entries of a section into a map object (`Section.EntryMap()`).

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
	confFile, _ := os.Open("fluentbit.conf")
	defer confFile.Close()

	conf := parser.NewFluentBitConfParser(confFile).Parse()

	for _, include := range conf.Includes {
		fmt.Printf("@INCLUDE %v\n", include)
	}

	for key, value := range conf.Sets {
		fmt.Printf("@SET %v=%v\n", key, value)
	}

	for _, section := range conf.Sections {
		fmt.Printf("[%v]\n", section.Name)
		for _, entry := range section.Entries {
			fmt.Printf("    %v %v\n", entry.Key, entry.Value)
		}
	}
}

```

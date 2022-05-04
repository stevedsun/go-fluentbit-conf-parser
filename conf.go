package parser

import (
	"reflect"
	"strconv"
)

type FluentBitConf struct {
	Sections []Section
}

type Section struct {
	Name    string
	Entries []Entry
}

type Entry struct {
	Key   string
	Value interface{}
}

func NewSection(name string) *Section {
	return &Section{
		Name:    name,
		Entries: []Entry{},
	}
}

func (section *Section) BindEntry(key string, strValue string) {

	var value interface{}
	if intValue, err := strconv.Atoi(strValue); err == nil {
		value = intValue
	} else if boolValue, err := strconv.ParseBool(strValue); err == nil {
		value = boolValue
	} else {
		value = strValue
	}

	entry := Entry{
		Key:   key,
		Value: value,
	}

	section.Entries = append(section.Entries, entry)
}

func (section *Section) EntryMap() map[string]interface{} {
	var entryMap = make(map[string]interface{})
	for _, entry := range section.Entries {
		if existValue, ok := entryMap[entry.Key]; !ok {
			entryMap[entry.Key] = entry.Value
		} else {
			if reflect.TypeOf(existValue).Kind().String() == "slice" {
				existValue = append(existValue.([]interface{}), entry.Value)
				entryMap[entry.Key] = existValue
			} else {
				arr := []interface{}{}
				arr = append(arr, existValue)
				arr = append(arr, entry.Value)
				entryMap[entry.Key] = arr
			}
		}
	}
	return entryMap
}

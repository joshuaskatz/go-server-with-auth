package utils

import (
	"io/ioutil"
)

func ParseFile(filePath string) string {
	c, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	sql := string(c)

	return sql
}

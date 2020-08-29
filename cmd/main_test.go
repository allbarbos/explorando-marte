package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainProgram(t *testing.T) {
	filePath = "./testdata/valid.csv"
	main()
}

func TestMainProgram_Error(t *testing.T) {
	filePath = "./testdata/invalid-axis-x.csv"

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, `strconv.Atoi: parsing "a": invalid syntax`, r)
		}
	}()

	main()
}

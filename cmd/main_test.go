package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainProgram(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		filePath = "./testdata/valid.csv"
		main()
	})

	t.Run("when invalid axis x", func(t *testing.T) {
		filePath = "./testdata/invalid-axis-x.csv"

		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, `strconv.Atoi: parsing "a": invalid syntax`, r)
			}
		}()

		main()
	})

	t.Run("when invalid csv", func(t *testing.T) {
		filePath = "invalid"

		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, "couldn't open fileopen invalid: no such file or directory", r)
			}
		}()

		main()
	})
}

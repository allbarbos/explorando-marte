package utils_test

import (
	"encoding/csv"
	"explorando-marte/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenFile(t *testing.T) {
	t.Run("when valid path", func(t *testing.T) {
		filePath := "./testdata/valid.csv"
		f := utils.OpenFile(filePath)

		assert.NotNil(t, f)
	})

	t.Run("when invalid path ", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, `couldn't open fileopen path_invalid: no such file or directory`, r)
			}
		}()

		filePath := "path_invalid"
		utils.OpenFile(filePath)
	})

}

func TestReadLine(t *testing.T) {
	t.Run("when valid csv", func(t *testing.T) {
		filePath := "./testdata/valid.csv"
		file := utils.OpenFile(filePath)
		defer file.Close()
		r := csv.NewReader(file)
		rec := utils.ReadLine(r)

		assert.NotNil(t, rec)
		assert.Len(t, rec, 3)

		rec = utils.ReadLine(r)
		assert.Nil(t, rec)
	})

	t.Run("when invalid csv ", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, "csv: invalid field or comment delimiter", r)
			}
		}()

		filePath := "./testdata/invalid.csv"
		file := utils.OpenFile(filePath)
		defer file.Close()
		r := csv.NewReader(file)
		r.Comma = '�'
		r.Comment = 1
		utils.ReadLine(r)
	})
}

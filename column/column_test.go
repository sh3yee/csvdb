package column

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumn_Add(t *testing.T) {
	t.Run("should_return_nil_when_add_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.Add("city")
		assert.Nil(t, err)

		header, table := readCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", ""}}, table)
	})
}

func TestColumn_Alter(t *testing.T) {
	t.Run("should_return_nil_when_alter_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.Alter("name", "username")
		assert.Nil(t, err)

		header, table := readCSV(t, path)
		assert.Equal(t, []string{"username", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}}, table)
	})

	t.Run("should_not_affect_other_when_alter_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.Alter("age", "years")
		assert.Nil(t, err)

		header, table := readCSV(t, path)
		assert.Equal(t, []string{"name", "years", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", "beijing"}}, table)
	})
}

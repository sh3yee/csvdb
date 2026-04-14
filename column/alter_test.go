package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sh3yee/csvdb/testutil"
)

func TestColumn_Alter(t *testing.T) {
	t.Run("should_return_nil_when_alter_column_by_name", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.Alter("name", "username")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"username", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}}, table)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.Alter("not_exist", "username")
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestColumn_AlterByIndex(t *testing.T) {
	t.Run("should_return_nil_when_alter_first_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.AlterByIndex(0, "username")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"username", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_alter_middle_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.AlterByIndex(1, "years")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "years", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_alter_last_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.AlterByIndex(2, "address")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "address"}, header)
		assert.Equal(t, [][]string{{"tom", "20", "beijing"}}, table)
	})

	t.Run("should_return_error_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.AlterByIndex(100, "username")
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.AlterByIndex(-1, "username")
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
}

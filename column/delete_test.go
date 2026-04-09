package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gycsv/testutil"
)

func TestColumn_DeleteByName(t *testing.T) {
	t.Run("should_return_nil_when_delete_column_in_multi_columns_csv", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "shanghai"},
		})

		column := New(path)
		err := column.DeleteByName("age")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "beijing"}, {"jerry", "shanghai"}}, table)
	})

	t.Run("should_return_nil_when_delete_only_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name"}, [][]string{
			{"tom"},
			{"jerry"},
		})

		column := New(path)
		err := column.DeleteByName("name")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{}, header)
		assert.Equal(t, [][]string{}, table)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.DeleteByName("not_exist")
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestColumn_DeleteByIndex(t *testing.T) {
	t.Run("should_return_nil_when_delete_first_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.DeleteByIndex(0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"age", "city"}, header)
		assert.Equal(t, [][]string{{"20", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_delete_last_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.DeleteByIndex(2)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}}, table)
	})

	t.Run("should_return_nil_when_delete_middle_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.DeleteByIndex(1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_delete_only_column", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name"}, [][]string{
			{"tom"},
		})

		column := New(path)
		err := column.DeleteByIndex(0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{}, header)
		assert.Equal(t, [][]string{}, table)
	})

	t.Run("should_return_error_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.DeleteByIndex(100)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.DeleteByIndex(-1)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
}

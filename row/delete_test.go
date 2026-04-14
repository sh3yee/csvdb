package row

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sh3yee/csvdb/testutil"
)

func TestRow_Delete(t *testing.T) {
	t.Run("should_return_nil_when_delete_first_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		err := r.Delete(0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_delete_last_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		err := r.Delete(1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}}, table)
	})

	t.Run("should_return_nil_when_delete_middle_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		r := New(path)
		err := r.Delete(1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"spike", "30"}}, table)
	})

	t.Run("should_return_nil_when_delete_only_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.Delete(0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{}, table)
	})

	t.Run("should_return_error_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.Delete(100)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.Delete(-1)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
}

func TestRow_DeleteBy(t *testing.T) {
	t.Run("should_return_nil_when_delete_rows_by_column_value", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
			{"jerry", "shanghai"},
			{"spike", "beijing"},
		})

		r := New(path)
		err := r.DeleteBy("city", "beijing")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"jerry", "shanghai"}}, table)
	})

	t.Run("should_return_nil_when_no_matching_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
		})

		r := New(path)
		err := r.DeleteBy("city", "shanghai")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "beijing"}}, table)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.DeleteBy("not_exist", "value")
		assert.Equal(t, ErrRowNotFound, err)
	})
}

package row

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sheye/csvdb/testutil"
)

func TestRow_Get(t *testing.T) {
	t.Run("should_return_row_when_get_by_valid_index", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		row, err := r.Get(0)
		assert.Nil(t, err)
		assert.Equal(t, []string{"tom", "20"}, row)
	})

	t.Run("should_return_last_row_when_get_last_index", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		row, err := r.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, []string{"jerry", "25"}, row)
	})

	t.Run("should_return_error_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		_, err := r.Get(100)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		_, err := r.Get(-1)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_empty_table", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{})

		r := New(path)
		_, err := r.Get(0)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
}

func TestRow_GetBy(t *testing.T) {
	t.Run("should_return_rows_when_find_by_column_value", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
			{"jerry", "shanghai"},
			{"spike", "beijing"},
		})

		r := New(path)
		rows, err := r.GetBy("city", "beijing")
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "beijing"}, {"spike", "beijing"}}, rows)
	})

	t.Run("should_return_empty_when_no_matching_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
		})

		r := New(path)
		rows, err := r.GetBy("city", "shanghai")
		assert.Nil(t, err)
		assert.Equal(t, [][]string{}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		_, err := r.GetBy("not_exist", "value")
		assert.Equal(t, ErrRowNotFound, err)
	})
}

func TestRow_GetAll(t *testing.T) {
	t.Run("should_return_all_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		rows, err := r.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, rows)
	})

	t.Run("should_return_empty_when_no_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{})

		r := New(path)
		rows, err := r.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{}, rows)
	})
}

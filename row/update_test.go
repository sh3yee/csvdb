package row

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sheye/csvdb/testutil"
)

func TestRow_Update(t *testing.T) {
	t.Run("should_return_nil_when_update_first_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		err := r.Update(0, []string{"tommy", "21"})
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tommy", "21"}, {"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_update_last_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		err := r.Update(1, []string{"jerry2", "26"})
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry2", "26"}}, table)
	})

	t.Run("should_return_nil_when_update_middle_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		r := New(path)
		err := r.Update(1, []string{"jerry2", "26"})
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry2", "26"}, {"spike", "30"}}, table)
	})

	t.Run("should_return_error_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.Update(100, []string{"test", "0"})
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.Update(-1, []string{"test", "0"})
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
}

func TestRow_UpdateBy(t *testing.T) {
	t.Run("should_return_nil_when_update_rows_by_column_value", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
			{"jerry", "shanghai"},
			{"spike", "beijing"},
		})

		r := New(path)
		err := r.UpdateBy("city", "beijing", []string{"updated", "beijing"})
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"updated", "beijing"}, {"jerry", "shanghai"}, {"updated", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_no_matching_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
		})

		r := New(path)
		err := r.UpdateBy("city", "shanghai", []string{"updated", "shanghai"})
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
		err := r.UpdateBy("not_exist", "value", []string{"test", "0"})
		assert.Equal(t, ErrRowNotFound, err)
	})
}

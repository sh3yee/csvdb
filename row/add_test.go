package row

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sheye/csvdb/testutil"
)

func TestRow_Add(t *testing.T) {
	t.Run("should_return_nil_when_add_row_to_end", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.Add([]string{"jerry", "25"})
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_add_row_to_empty_file", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{})

		r := New(path)
		err := r.Add([]string{"tom", "20"})
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}}, table)
	})
}

func TestRow_AddAt(t *testing.T) {
	t.Run("should_return_nil_when_add_row_at_beginning", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"jerry", "25"},
		})

		r := New(path)
		err := r.AddAt([]string{"tom", "20"}, 0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_add_row_at_middle", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		r := New(path)
		err := r.AddAt([]string{"spike", "30"}, 1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"spike", "30"}, {"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_add_row_at_end", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.AddAt([]string{"jerry", "25"}, 1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"jerry", "25"},
		})

		r := New(path)
		err := r.AddAt([]string{"tom", "20"}, -1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, table)
	})

	t.Run("should_return_nil_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		r := New(path)
		err := r.AddAt([]string{"jerry", "25"}, 100)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, table)
	})
}

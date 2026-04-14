package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sheye/csvdb/testutil"
)

func TestColumn_Add(t *testing.T) {
	t.Run("should_return_nil_when_add_column_to_end", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.Add("city")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", ""}}, table)
	})
}

func TestColumn_AddAt(t *testing.T) {
	t.Run("should_return_nil_when_add_column_at_beginning", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"age", "city"}, [][]string{
			{"20", "beijing"},
		})

		column := New(path)
		err := column.AddAt("name", 0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"", "20", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_add_column_at_middle", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
		})

		column := New(path)
		err := column.AddAt("age", 1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_add_column_at_end", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.AddAt("city", 2)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", ""}}, table)
	})

	t.Run("should_return_nil_when_index_negative", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"age", "city"}, [][]string{
			{"20", "beijing"},
		})

		column := New(path)
		err := column.AddAt("name", -1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"", "20", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_index_out_of_range", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name"}, [][]string{
			{"tom"},
		})

		column := New(path)
		err := column.AddAt("age", 100)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", ""}}, table)
	})
}

func TestColumn_AddWithDefault(t *testing.T) {
	t.Run("should_return_nil_when_add_column_with_default_value", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		column := New(path)
		err := column.AddWithDefault("city", "unknown")
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "20", "unknown"}, {"jerry", "25", "unknown"}}, table)
	})
}

func TestColumn_AddAtWithDefault(t *testing.T) {
	t.Run("should_return_nil_when_add_column_at_position_with_default_value", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "city"}, [][]string{
			{"tom", "beijing"},
		})

		column := New(path)
		err := column.AddAtWithDefault("age", "0", 1)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "0", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_add_column_at_beginning_with_default_value", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"age", "city"}, [][]string{
			{"20", "beijing"},
		})

		column := New(path)
		err := column.AddAtWithDefault("name", "unknown", 0)
		assert.Nil(t, err)

		header, table := testutil.ReadCSV(t, path)
		assert.Equal(t, []string{"name", "age", "city"}, header)
		assert.Equal(t, [][]string{{"unknown", "20", "beijing"}}, table)
	})
}

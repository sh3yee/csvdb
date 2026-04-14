package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sheye/csvdb/testutil"
)

func TestQuery_Find(t *testing.T) {
	t.Run("should_return_rows_when_equal_condition", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "shanghai"},
			{"spike", "30", "beijing"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "city", Op: "=", Value: "beijing"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "20", "beijing"}, {"spike", "30", "beijing"}}, rows)
	})

	t.Run("should_return_rows_when_not_equal_condition", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "name", Op: "!=", Value: "tom"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"jerry", "25"}}, rows)
	})

	t.Run("should_return_rows_when_greater_than_condition", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "age", Op: ">", Value: "24"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"jerry", "25"}, {"spike", "30"}}, rows)
	})

	t.Run("should_return_rows_when_less_than_condition", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "age", Op: "<", Value: "26"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, rows)
	})

	t.Run("should_return_rows_when_like_condition_with_prefix", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "email"}, [][]string{
			{"tom", "tom@example.com"},
			{"jerry", "jerry@test.com"},
			{"tim", "tim@example.com"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "email", Op: "like", Value: "%example.com"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "tom@example.com"}, {"tim", "tim@example.com"}}, rows)
	})

	t.Run("should_return_rows_when_like_condition_with_suffix", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "email"}, [][]string{
			{"tom", "tom@example.com"},
			{"tim", "tim@test.com"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "name", Op: "like", Value: "t%"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "tom@example.com"}, {"tim", "tim@test.com"}}, rows)
	})

	t.Run("should_return_rows_when_like_condition_with_contains", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "email"}, [][]string{
			{"tom", "tom@example.com"},
			{"jerry", "jerry@test.com"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "email", Op: "like", Value: "%example%"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "tom@example.com"}}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		_, err := q.Find(Condition{Column: "not_exist", Op: "=", Value: "value"})
		assert.Equal(t, ErrColumnNotFound, err)
	})

	t.Run("should_return_empty_when_no_match", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "name", Op: "=", Value: "not_exist"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{}, rows)
	})
}

func TestQuery_FindAll(t *testing.T) {
	t.Run("should_return_rows_when_multiple_conditions", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "beijing"},
			{"spike", "30", "shanghai"},
		})

		q := New(path)
		rows, err := q.FindAll(
			Condition{Column: "age", Op: ">", Value: "22"},
			Condition{Column: "city", Op: "=", Value: "beijing"},
		).Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"jerry", "25", "beijing"}}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		_, err := q.FindAll(
			Condition{Column: "not_exist", Op: "=", Value: "value"},
		).Get()
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestQuery_FindIn(t *testing.T) {
	t.Run("should_return_rows_when_values_in_list", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"id", "name"}, [][]string{
			{"1", "tom"},
			{"2", "jerry"},
			{"3", "spike"},
			{"4", "tyke"},
		})

		q := New(path)
		rows, err := q.FindIn("id", []string{"1", "3"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"1", "tom"}, {"3", "spike"}}, rows)
	})

	t.Run("should_return_empty_when_no_match", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"id", "name"}, [][]string{
			{"1", "tom"},
		})

		q := New(path)
		rows, err := q.FindIn("id", []string{"99", "100"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name"}, [][]string{
			{"tom"},
		})

		q := New(path)
		_, err := q.FindIn("not_exist", []string{"1"})
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestQuery_FindNotIn(t *testing.T) {
	t.Run("should_return_rows_when_values_not_in_list", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"id", "name"}, [][]string{
			{"1", "tom"},
			{"2", "jerry"},
			{"3", "spike"},
		})

		q := New(path)
		rows, err := q.FindNotIn("id", []string{"1"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"2", "jerry"}, {"3", "spike"}}, rows)
	})
}

func TestResult_Select(t *testing.T) {
	t.Run("should_return_selected_columns", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "shanghai"},
		})

		q := New(path)
		rows, err := q.FindAll(Condition{Column: "age", Op: ">", Value: "18"}).Select("name", "city").Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "beijing"}, {"jerry", "shanghai"}}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		_, err := q.FindAll().Select("not_exist").Get()
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestResult_Limit(t *testing.T) {
	t.Run("should_limit_result_count", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		rows, err := q.FindAll().Limit(2).Get()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(rows))
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, rows)
	})

	t.Run("should_return_all_when_limit_greater_than_count", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		rows, err := q.FindAll().Limit(10).Get()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(rows))
	})
}

func TestResult_Offset(t *testing.T) {
	t.Run("should_skip_first_n_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		rows, err := q.FindAll().Offset(1).Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"jerry", "25"}, {"spike", "30"}}, rows)
	})

	t.Run("should_return_empty_when_offset_greater_than_count", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		rows, err := q.FindAll().Offset(10).Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{}, rows)
	})
}

func TestResult_First(t *testing.T) {
	t.Run("should_return_first_row", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
		})

		q := New(path)
		row, err := q.FindAll().First()
		assert.Nil(t, err)
		assert.Equal(t, []string{"tom", "20"}, row)
	})

	t.Run("should_return_error_when_no_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{})

		q := New(path)
		_, err := q.FindAll().First()
		assert.Equal(t, ErrRowNotFound, err)
	})
}

func TestResult_Count(t *testing.T) {
	t.Run("should_return_count", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		count, err := q.FindAll().Count()
		assert.Nil(t, err)
		assert.Equal(t, 3, count)
	})
}

func TestResult_Exists(t *testing.T) {
	t.Run("should_return_true_when_rows_exist", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		exists, err := q.FindAll(Condition{Column: "name", Op: "=", Value: "tom"}).Exists()
		assert.Nil(t, err)
		assert.True(t, exists)
	})

	t.Run("should_return_false_when_no_rows", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		exists, err := q.FindAll(Condition{Column: "name", Op: "=", Value: "not_exist"}).Exists()
		assert.Nil(t, err)
		assert.False(t, exists)
	})
}

func TestResult_OrderBy(t *testing.T) {
	t.Run("should_sort_ascending", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "30"},
			{"jerry", "20"},
			{"spike", "25"},
		})

		q := New(path)
		rows, err := q.FindAll().OrderBy("age", "asc").Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{
			{"jerry", "20"},
			{"spike", "25"},
			{"tom", "30"},
		}, rows)
	})

	t.Run("should_sort_descending", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "30"},
			{"spike", "25"},
		})

		q := New(path)
		rows, err := q.FindAll().OrderBy("age", "desc").Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{
			{"jerry", "30"},
			{"spike", "25"},
			{"tom", "20"},
		}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		_, err := q.FindAll().OrderBy("not_exist", "asc").Get()
		assert.Equal(t, ErrColumnNotFound, err)
	})

	t.Run("should_sort_with_limit", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "30"},
			{"jerry", "20"},
			{"spike", "25"},
		})

		q := New(path)
		rows, err := q.FindAll().OrderBy("age", "asc").Limit(2).Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{
			{"jerry", "20"},
			{"spike", "25"},
		}, rows)
	})
}

func TestResult_ThenBy(t *testing.T) {
	t.Run("should_sort_by_multiple_columns", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "shanghai"},
			{"spike", "20", "shanghai"},
			{"tyke", "25", "beijing"},
		})

		q := New(path)
		rows, err := q.FindAll().OrderBy("age", "asc").ThenBy("city", "asc").Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{
			{"tom", "20", "beijing"},
			{"spike", "20", "shanghai"},
			{"tyke", "25", "beijing"},
			{"jerry", "25", "shanghai"},
		}, rows)
	})

	t.Run("should_sort_by_multiple_columns_with_different_orders", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "shanghai"},
			{"spike", "20", "shanghai"},
		})

		q := New(path)
		rows, err := q.FindAll().OrderBy("age", "desc").ThenBy("name", "asc").Get()
		assert.Nil(t, err)
		assert.Equal(t, [][]string{
			{"jerry", "25", "shanghai"},
			{"spike", "20", "shanghai"},
			{"tom", "20", "beijing"},
		}, rows)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		q := New(path)
		_, err := q.FindAll().OrderBy("age", "asc").ThenBy("not_exist", "asc").Get()
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestQuery_Find_Additional(t *testing.T) {
	t.Run("should_return_rows_when_greater_equal_condition", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "age", Op: ">=", Value: "25"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"jerry", "25"}, {"spike", "30"}}, rows)
	})

	t.Run("should_return_rows_when_less_equal_condition", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
			{"jerry", "25"},
			{"spike", "30"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "age", Op: "<=", Value: "25"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "20"}, {"jerry", "25"}}, rows)
	})

	t.Run("should_return_rows_when_like_condition_without_wildcard", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name", "email"}, [][]string{
			{"tom", "tom@example.com"},
			{"jerry", "jerry@test.com"},
		})

		q := New(path)
		rows, err := q.Find(Condition{Column: "name", Op: "like", Value: "tom"})
		assert.Nil(t, err)
		assert.Equal(t, [][]string{{"tom", "tom@example.com"}}, rows)
	})
}

func TestQuery_FindNotIn_Additional(t *testing.T) {
	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := testutil.CreateCSV(t, []string{"name"}, [][]string{
			{"tom"},
		})

		q := New(path)
		_, err := q.FindNotIn("not_exist", []string{"1"})
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

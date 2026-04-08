package column

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestColumn_DeleteByName(t *testing.T) {
	t.Run("should_return_nil_when_delete_column_in_multi_columns_csv", func(t *testing.T) {
		// 创建临时测试文件
		path := createTestCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
			{"jerry", "25", "shanghai"},
		})

		column := New(path)
		err := column.DeleteByName("age")
		assert.Nil(t, err)

		// 验证删除结果
		header, table := readCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "beijing"}, {"jerry", "shanghai"}}, table)
	})

	t.Run("should_return_nil_when_delete_only_column", func(t *testing.T) {
		// 只有一列的情况
		path := createTestCSV(t, []string{"name"}, [][]string{
			{"tom"},
			{"jerry"},
		})

		column := New(path)
		err := column.DeleteByName("name")
		assert.Nil(t, err)

		// 验证删除结果：删除唯一列后，header 和 table 都为空
		header, table := readCSV(t, path)
		assert.Equal(t, []string{}, header)
		assert.Equal(t, [][]string{}, table)
	})

	t.Run("should_return_error_when_column_not_found", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.DeleteByName("not_exist")
		assert.Equal(t, ErrColumnNotFound, err)
	})
}

func TestColumn_DeleteByIndex(t *testing.T) {
	t.Run("should_return_nil_when_delete_first_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.DeleteByIndex(0)
		assert.Nil(t, err)

		header, table := readCSV(t, path)
		assert.Equal(t, []string{"age", "city"}, header)
		assert.Equal(t, [][]string{{"20", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_delete_last_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.DeleteByIndex(2)
		assert.Nil(t, err)

		header, table := readCSV(t, path)
		assert.Equal(t, []string{"name", "age"}, header)
		assert.Equal(t, [][]string{{"tom", "20"}}, table)
	})

	t.Run("should_return_nil_when_delete_middle_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age", "city"}, [][]string{
			{"tom", "20", "beijing"},
		})

		column := New(path)
		err := column.DeleteByIndex(1)
		assert.Nil(t, err)

		header, table := readCSV(t, path)
		assert.Equal(t, []string{"name", "city"}, header)
		assert.Equal(t, [][]string{{"tom", "beijing"}}, table)
	})

	t.Run("should_return_nil_when_delete_only_column", func(t *testing.T) {
		path := createTestCSV(t, []string{"name"}, [][]string{
			{"tom"},
		})

		column := New(path)
		err := column.DeleteByIndex(0)
		assert.Nil(t, err)

		// 删除唯一列后，header 和 table 都为空
		header, table := readCSV(t, path)
		assert.Equal(t, []string{}, header)
		assert.Equal(t, [][]string{}, table)
	})

	t.Run("should_return_error_when_index_out_of_range", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.DeleteByIndex(100)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})

	t.Run("should_return_error_when_index_negative", func(t *testing.T) {
		path := createTestCSV(t, []string{"name", "age"}, [][]string{
			{"tom", "20"},
		})

		column := New(path)
		err := column.DeleteByIndex(-1)
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
}

// createTestCSV 创建临时测试 CSV 文件
func createTestCSV(t *testing.T, header []string, rows [][]string) string {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "test_*.csv")
	assert.Nil(t, err)
	defer tmpFile.Close()

	writer := csv.NewWriter(tmpFile)
	err = writer.Write(header)
	assert.Nil(t, err)

	for _, row := range rows {
		err = writer.Write(row)
		assert.Nil(t, err)
	}
	writer.Flush()

	// 注册清理函数
	t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	return tmpFile.Name()
}

// readCSV 读取 CSV 文件内容
func readCSV(t *testing.T, path string) ([]string, [][]string) {
	t.Helper()

	file, err := os.Open(path)
	assert.Nil(t, err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	assert.Nil(t, err)

	if len(records) == 0 {
		return []string{}, [][]string{}
	}

	return records[0], records[1:]
}

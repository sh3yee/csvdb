package testutil

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CreateCSV 创建临时测试 CSV 文件
func CreateCSV(t *testing.T, header []string, rows [][]string) string {
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

	t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	return tmpFile.Name()
}

// ReadCSV 读取 CSV 文件内容
func ReadCSV(t *testing.T, path string) ([]string, [][]string) {
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

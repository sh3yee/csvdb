package row

import "github.com/sheye/csvdb/internal/file"

// Update 更新指定索引的行
func (r *Row) Update(index int, values []string) error {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	if index < 0 || index >= len(rwf.Table) {
		return ErrIndexOutOfRange
	}

	rwf.Table[index] = values

	return rwf.Write()
}

// UpdateBy 更新匹配条件的行
func (r *Row) UpdateBy(column, value string, newValues []string) error {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	// 查找列索引
	colIndex := -1
	for i, h := range rwf.Header {
		if h == column {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		return ErrRowNotFound
	}

	// 更新匹配的行
	for i, row := range rwf.Table {
		if colIndex < len(row) && row[colIndex] == value {
			rwf.Table[i] = newValues
		}
	}

	return rwf.Write()
}

package row

import "gycsv/internal/file"

// Delete 删除指定索引的行
func (r *Row) Delete(index int) error {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	if index < 0 || index >= len(rwf.Table) {
		return ErrIndexOutOfRange
	}

	rwf.Table = append(rwf.Table[:index], rwf.Table[index+1:]...)

	return rwf.Write()
}

// DeleteBy 删除匹配条件的行
func (r *Row) DeleteBy(column, value string) error {
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

	// 过滤掉匹配的行
	newTable := make([][]string, 0, len(rwf.Table))
	for _, row := range rwf.Table {
		if colIndex >= len(row) || row[colIndex] != value {
			newTable = append(newTable, row)
		}
	}

	rwf.Table = newTable

	return rwf.Write()
}

package column

import "gycsv/internal/file"

// DeleteByName 根据列名删除列
func (c *Column) DeleteByName(name string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	// 查找列索引
	index := -1
	for i, h := range rwf.Header {
		if h == name {
			index = i
			break
		}
	}

	if index == -1 {
		return ErrColumnNotFound
	}

	return c.deleteColumn(rwf, index)
}

// DeleteByIndex 根据索引删除列
func (c *Column) DeleteByIndex(index int) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	if index < 0 || index >= len(rwf.Header) {
		return ErrIndexOutOfRange
	}

	return c.deleteColumn(rwf, index)
}

// deleteColumn 删除指定索引的列
func (c *Column) deleteColumn(rwf *file.Data, index int) error {
	// 从 header 中删除
	rwf.Header = append(rwf.Header[:index], rwf.Header[index+1:]...)

	// 从 table 每行中删除
	for i, row := range rwf.Table {
		if index < len(row) {
			rwf.Table[i] = append(row[:index], row[index+1:]...)
		}
	}

	return rwf.Write()
}

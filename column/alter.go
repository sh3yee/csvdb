package column

import "gycsv/internal/file"

// Alter 根据列名修改列名
func (c *Column) Alter(oldName string, newName string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	// 查找列索引
	index := -1
	for i, h := range rwf.Header {
		if h == oldName {
			index = i
			break
		}
	}

	if index == -1 {
		return ErrColumnNotFound
	}

	return c.alterColumn(rwf, index, newName)
}

// AlterByIndex 根据索引修改列名
func (c *Column) AlterByIndex(index int, newName string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	if index < 0 || index >= len(rwf.Header) {
		return ErrIndexOutOfRange
	}

	return c.alterColumn(rwf, index, newName)
}

// alterColumn 修改指定索引的列名
func (c *Column) alterColumn(rwf *file.Data, index int, newName string) error {
	rwf.Header[index] = newName

	return rwf.Write()
}

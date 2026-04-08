package column

import "gycsv/internal/file"

// Add 添加新列
func (c *Column) Add(field string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	// 在 header 中添加新列
	rwf.Header = append(rwf.Header, field)

	// 在 table 每行中添加空值
	for i := range rwf.Table {
		rwf.Table[i] = append(rwf.Table[i], "")
	}

	return rwf.Write()
}

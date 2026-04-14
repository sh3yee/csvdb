package column

import "github.com/sheye/csvdb/internal/file"

// Add 添加新列到末尾
func (c *Column) Add(field string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	// 在 header 末尾添加新列
	rwf.Header = append(rwf.Header, field)

	// 在 table 每行末尾添加空值
	for i := range rwf.Table {
		rwf.Table[i] = append(rwf.Table[i], "")
	}

	return rwf.Write()
}

// AddAt 添加新列到指定位置
func (c *Column) AddAt(field string, index int) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	return c.addColumn(rwf, field, "", index)
}

// AddWithDefault 添加带默认值的新列到末尾
func (c *Column) AddWithDefault(field string, defaultValue string) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	return c.addColumn(rwf, field, defaultValue, len(rwf.Header))
}

// AddAtWithDefault 添加带默认值的新列到指定位置
func (c *Column) AddAtWithDefault(field string, defaultValue string, index int) error {
	rwf := file.New(c.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	return c.addColumn(rwf, field, defaultValue, index)
}

// addColumn 添加新列的核心方法
func (c *Column) addColumn(rwf *file.Data, field string, defaultValue string, index int) error {
	// 处理索引边界
	if index < 0 {
		index = 0
	}
	if index > len(rwf.Header) {
		index = len(rwf.Header)
	}

	// 在 header 中插入新列
	rwf.Header = append(rwf.Header[:index], append([]string{field}, rwf.Header[index:]...)...)

	// 在 table 每行中插入默认值
	for i, row := range rwf.Table {
		if index > len(row) {
			rwf.Table[i] = append(row, defaultValue)
		} else {
			rwf.Table[i] = append(row[:index], append([]string{defaultValue}, row[index:]...)...)
		}
	}

	return rwf.Write()
}

package row

import "github.com/sheye/csvdb/internal/file"

// Add 添加新行到末尾
func (r *Row) Add(values []string) error {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	rwf.Table = append(rwf.Table, values)

	return rwf.Write()
}

// AddAt 添加新行到指定位置
func (r *Row) AddAt(values []string, index int) error {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return err
	}

	// 处理索引边界
	if index < 0 {
		index = 0
	}
	if index > len(rwf.Table) {
		index = len(rwf.Table)
	}

	rwf.Table = append(rwf.Table[:index], append([][]string{values}, rwf.Table[index:]...)...)

	return rwf.Write()
}

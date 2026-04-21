package row

import "github.com/sh3yee/csvdb/internal/file"

// Get 获取指定索引的行
func (r *Row) Get(index int) ([]string, error) {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return nil, err
	}

	if index < 0 || index >= len(rwf.Table) {
		return nil, ErrIndexOutOfRange
	}

	return rwf.Table[index], nil
}

// GetBy 获取匹配条件的行
func (r *Row) GetBy(column, value string) ([][]string, error) {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return nil, err
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
		return nil, ErrRowNotFound
	}

	// 筛选匹配的行
	result := make([][]string, 0)
	for _, row := range rwf.Table {
		if colIndex < len(row) && row[colIndex] == value {
			result = append(result, row)
		}
	}

	return result, nil
}

// GetAll 获取所有行
func (r *Row) GetAll() ([][]string, error) {
	rwf := file.New(r.Path)

	if err := rwf.Read(); err != nil {
		return nil, err
	}

	if rwf.Table == nil {
		return [][]string{}, nil
	}

	return rwf.Table, nil
}

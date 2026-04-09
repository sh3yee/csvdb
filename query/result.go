package query

// Result 查询结果，支持链式处理
type Result struct {
	rows   [][]string
	header []string
	err    error
}

// Select 选择特定列
func (r *Result) Select(columns ...string) *Result {
	if r.err != nil {
		return r
	}

	// 查找列索引
	colIndices := make([]int, 0, len(columns))
	for _, col := range columns {
		found := false
		for i, h := range r.header {
			if h == col {
				colIndices = append(colIndices, i)
				found = true
				break
			}
		}
		if !found {
			r.err = ErrColumnNotFound
			return r
		}
	}

	// 提取指定列
	result := make([][]string, 0, len(r.rows))
	for _, row := range r.rows {
		newRow := make([]string, 0, len(colIndices))
		for _, idx := range colIndices {
			if idx < len(row) {
				newRow = append(newRow, row[idx])
			} else {
				newRow = append(newRow, "")
			}
		}
		result = append(result, newRow)
	}

	r.rows = result
	r.header = columns
	return r
}

// Limit 限制结果数量
func (r *Result) Limit(n int) *Result {
	if r.err != nil {
		return r
	}

	if n >= len(r.rows) {
		return r
	}

	r.rows = r.rows[:n]
	return r
}

// Offset 跳过前 N 条
func (r *Result) Offset(n int) *Result {
	if r.err != nil {
		return r
	}

	if n >= len(r.rows) {
		r.rows = nil
		return r
	}

	r.rows = r.rows[n:]
	return r
}

// Get 获取所有结果
func (r *Result) Get() ([][]string, error) {
	return r.rows, r.err
}

// First 获取第一条结果
func (r *Result) First() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}

	if len(r.rows) == 0 {
		return nil, ErrRowNotFound
	}

	return r.rows[0], nil
}

// Count 统计结果数量
func (r *Result) Count() (int, error) {
	return len(r.rows), r.err
}

// Exists 判断是否存在匹配结果
func (r *Result) Exists() (bool, error) {
	if r.err != nil {
		return false, r.err
	}

	return len(r.rows) > 0, nil
}

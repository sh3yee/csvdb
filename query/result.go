package query

type sortOrder struct {
	column    string
	index     int
	ascending bool
}

// Result 查询结果，支持链式处理
type Result struct {
	rows       [][]string
	header     []string
	sortOrders []sortOrder
	err        error
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

// OrderBy 单列排序
func (r *Result) OrderBy(column, order string) *Result {
	if r.err != nil {
		return r
	}

	// 查找列索引
	colIndex := -1
	for i, h := range r.header {
		if h == column {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		r.err = ErrColumnNotFound
		return r
	}

	ascending := order != "desc"
	r.sortOrders = []sortOrder{{column: column, index: colIndex, ascending: ascending}}
	return r
}

// ThenBy 多列排序，追加排序条件
func (r *Result) ThenBy(column, order string) *Result {
	if r.err != nil {
		return r
	}

	// 查找列索引
	colIndex := -1
	for i, h := range r.header {
		if h == column {
			colIndex = i
			break
		}
	}

	if colIndex == -1 {
		r.err = ErrColumnNotFound
		return r
	}

	ascending := order != "desc"
	r.sortOrders = append(r.sortOrders, sortOrder{column: column, index: colIndex, ascending: ascending})
	return r
}

// doSort 执行排序
func (r *Result) doSort() {
	if len(r.sortOrders) == 0 || len(r.rows) <= 1 {
		return
	}

	// 使用稳定的排序方式
	for i := len(r.sortOrders) - 1; i >= 0; i-- {
		so := r.sortOrders[i]
		r.quickSort(so.index, so.ascending)
	}
}

// quickSort 快速排序
func (r *Result) quickSort(colIndex int, ascending bool) {
	if len(r.rows) <= 1 {
		return
	}

	// 使用标准库的排序
	for i := 0; i < len(r.rows)-1; i++ {
		for j := i + 1; j < len(r.rows); j++ {
			var shouldSwap bool
			vi := ""
			vj := ""
			if colIndex < len(r.rows[i]) {
				vi = r.rows[i][colIndex]
			}
			if colIndex < len(r.rows[j]) {
				vj = r.rows[j][colIndex]
			}

			if ascending {
				shouldSwap = vi > vj
			} else {
				shouldSwap = vi < vj
			}

			if shouldSwap {
				r.rows[i], r.rows[j] = r.rows[j], r.rows[i]
			}
		}
	}
}

// Get 获取所有结果
func (r *Result) Get() ([][]string, error) {
	if r.err != nil {
		return nil, r.err
	}

	r.doSort()
	return r.rows, nil
}

// First 获取第一条结果
func (r *Result) First() ([]string, error) {
	if r.err != nil {
		return nil, r.err
	}

	if len(r.rows) == 0 {
		return nil, ErrRowNotFound
	}

	r.doSort()
	return r.rows[0], nil
}

// Count 统计结果数量
func (r *Result) Count() (int, error) {
	if r.err != nil {
		return 0, r.err
	}

	r.doSort()
	return len(r.rows), nil
}

// Exists 判断是否存在匹配结果
func (r *Result) Exists() (bool, error) {
	if r.err != nil {
		return false, r.err
	}

	return len(r.rows) > 0, nil
}

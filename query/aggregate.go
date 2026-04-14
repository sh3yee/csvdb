package query

import (
	"strconv"
)

// AggregateResult 聚合结果
type AggregateResult struct {
	rows   [][]string
	header []string
	err    error
}

// Sum 计算指定列的总和
func (r *Result) Sum(column string) (float64, error) {
	if r.err != nil {
		return 0, r.err
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
		return 0, ErrColumnNotFound
	}

	// 执行排序和分页
	r.doSort()
	r.applyOffset()
	r.applyLimit()

	var sum float64
	for _, row := range r.rows {
		if colIndex < len(row) {
			val, err := strconv.ParseFloat(row[colIndex], 64)
			if err != nil {
				continue // 跳过非数字值
			}
			sum += val
		}
	}

	return sum, nil
}

// Avg 计算指定列的平均值
func (r *Result) Avg(column string) (float64, error) {
	if r.err != nil {
		return 0, r.err
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
		return 0, ErrColumnNotFound
	}

	// 执行排序和分页
	r.doSort()
	r.applyOffset()
	r.applyLimit()

	var sum float64
	var count int
	for _, row := range r.rows {
		if colIndex < len(row) {
			val, err := strconv.ParseFloat(row[colIndex], 64)
			if err != nil {
				continue // 跳过非数字值
			}
			sum += val
			count++
		}
	}

	if count == 0 {
		return 0, nil
	}

	return sum / float64(count), nil
}

// Min 获取指定列的最小值（字符串比较）
func (r *Result) Min(column string) (string, error) {
	if r.err != nil {
		return "", r.err
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
		return "", ErrColumnNotFound
	}

	// 执行排序和分页
	r.doSort()
	r.applyOffset()
	r.applyLimit()

	if len(r.rows) == 0 {
		return "", ErrRowNotFound
	}

	minVal := r.rows[0][colIndex]
	for _, row := range r.rows {
		if colIndex < len(row) {
			if row[colIndex] < minVal {
				minVal = row[colIndex]
			}
		}
	}

	return minVal, nil
}

// Max 获取指定列的最大值（字符串比较）
func (r *Result) Max(column string) (string, error) {
	if r.err != nil {
		return "", r.err
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
		return "", ErrColumnNotFound
	}

	// 执行排序和分页
	r.doSort()
	r.applyOffset()
	r.applyLimit()

	if len(r.rows) == 0 {
		return "", ErrRowNotFound
	}

	maxVal := r.rows[0][colIndex]
	for _, row := range r.rows {
		if colIndex < len(row) {
			if row[colIndex] > maxVal {
				maxVal = row[colIndex]
			}
		}
	}

	return maxVal, nil
}

// MinFloat 获取指定列的最小值（数值比较）
func (r *Result) MinFloat(column string) (float64, error) {
	if r.err != nil {
		return 0, r.err
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
		return 0, ErrColumnNotFound
	}

	// 执行排序和分页
	r.doSort()
	r.applyOffset()
	r.applyLimit()

	if len(r.rows) == 0 {
		return 0, ErrRowNotFound
	}

	var minVal float64
	var found bool
	for _, row := range r.rows {
		if colIndex < len(row) {
			val, err := strconv.ParseFloat(row[colIndex], 64)
			if err != nil {
				continue
			}
			if !found || val < minVal {
				minVal = val
				found = true
			}
		}
	}

	if !found {
		return 0, ErrRowNotFound
	}

	return minVal, nil
}

// MaxFloat 获取指定列的最大值（数值比较）
func (r *Result) MaxFloat(column string) (float64, error) {
	if r.err != nil {
		return 0, r.err
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
		return 0, ErrColumnNotFound
	}

	// 执行排序和分页
	r.doSort()
	r.applyOffset()
	r.applyLimit()

	if len(r.rows) == 0 {
		return 0, ErrRowNotFound
	}

	var maxVal float64
	var found bool
	for _, row := range r.rows {
		if colIndex < len(row) {
			val, err := strconv.ParseFloat(row[colIndex], 64)
			if err != nil {
				continue
			}
			if !found || val > maxVal {
				maxVal = val
				found = true
			}
		}
	}

	if !found {
		return 0, ErrRowNotFound
	}

	return maxVal, nil
}

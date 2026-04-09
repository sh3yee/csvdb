package query

import (
	"gycsv/internal/file"
	"strings"
)

// Find 单条件查询，直接返回结果
func (q *Query) Find(cond Condition) ([][]string, error) {
	rwf := file.New(q.path)

	if err := rwf.Read(); err != nil {
		return nil, err
	}

	colIndex := q.findColumnIndex(rwf.Header, cond.Column)
	if colIndex == -1 {
		return nil, ErrColumnNotFound
	}

	result := make([][]string, 0)
	for _, row := range rwf.Table {
		if colIndex >= len(row) {
			continue
		}
		if q.match(row[colIndex], cond.Op, cond.Value) {
			result = append(result, row)
		}
	}

	return result, nil
}

// FindAll 多条件查询，返回 Result 支持链式处理
func (q *Query) FindAll(conds ...Condition) *Result {
	rwf := file.New(q.path)

	if err := rwf.Read(); err != nil {
		return &Result{rows: [][]string{}, header: nil, err: err}
	}

	// 验证所有列是否存在
	for _, cond := range conds {
		if q.findColumnIndex(rwf.Header, cond.Column) == -1 {
			return &Result{rows: [][]string{}, header: rwf.Header, err: ErrColumnNotFound}
		}
	}

	result := make([][]string, 0)
	for _, row := range rwf.Table {
		if q.matchAll(row, rwf.Header, conds) {
			result = append(result, row)
		}
	}

	return &Result{rows: result, header: rwf.Header, err: nil}
}

// FindIn IN 查询
func (q *Query) FindIn(column string, values []string) ([][]string, error) {
	rwf := file.New(q.path)

	if err := rwf.Read(); err != nil {
		return nil, err
	}

	colIndex := q.findColumnIndex(rwf.Header, column)
	if colIndex == -1 {
		return nil, ErrColumnNotFound
	}

	valueSet := make(map[string]bool)
	for _, v := range values {
		valueSet[v] = true
	}

	result := make([][]string, 0)
	for _, row := range rwf.Table {
		if colIndex < len(row) && valueSet[row[colIndex]] {
			result = append(result, row)
		}
	}

	return result, nil
}

// FindNotIn NOT IN 查询
func (q *Query) FindNotIn(column string, values []string) ([][]string, error) {
	rwf := file.New(q.path)

	if err := rwf.Read(); err != nil {
		return nil, err
	}

	colIndex := q.findColumnIndex(rwf.Header, column)
	if colIndex == -1 {
		return nil, ErrColumnNotFound
	}

	valueSet := make(map[string]bool)
	for _, v := range values {
		valueSet[v] = true
	}

	result := make([][]string, 0)
	for _, row := range rwf.Table {
		if colIndex < len(row) && !valueSet[row[colIndex]] {
			result = append(result, row)
		}
	}

	return result, nil
}

// findColumnIndex 查找列索引
func (q *Query) findColumnIndex(header []string, column string) int {
	for i, h := range header {
		if h == column {
			return i
		}
	}
	return -1
}

// match 判断值是否匹配条件
func (q *Query) match(cellValue, op, value string) bool {
	switch op {
	case "=":
		return cellValue == value
	case "!=":
		return cellValue != value
	case ">":
		return cellValue > value
	case "<":
		return cellValue < value
	case ">=":
		return cellValue >= value
	case "<=":
		return cellValue <= value
	case "like":
		return q.matchLike(cellValue, value)
	default:
		return false
	}
}

// matchLike 模糊匹配，支持 % 通配符
func (q *Query) matchLike(cellValue, pattern string) bool {
	// %xxx% -> contains
	// %xxx  -> has suffix
	// xxx%  -> has prefix
	// xxx   -> equals

	if strings.HasPrefix(pattern, "%") && strings.HasSuffix(pattern, "%") {
		return strings.Contains(cellValue, pattern[1:len(pattern)-1])
	}
	if strings.HasPrefix(pattern, "%") {
		return strings.HasSuffix(cellValue, pattern[1:])
	}
	if strings.HasSuffix(pattern, "%") {
		return strings.HasPrefix(cellValue, pattern[:len(pattern)-1])
	}
	return cellValue == pattern
}

// matchAll 判断行是否匹配所有条件
func (q *Query) matchAll(row []string, header []string, conds []Condition) bool {
	for _, cond := range conds {
		colIndex := q.findColumnIndex(header, cond.Column)
		if colIndex >= len(row) {
			return false
		}
		if !q.match(row[colIndex], cond.Op, cond.Value) {
			return false
		}
	}
	return true
}

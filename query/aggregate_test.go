package query

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sheye/csvdb/testutil"
)

func TestAggregate_Sum(t *testing.T) {
	dir, err := os.MkdirTemp("", "csvdb_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.csv")
	testutil.CreateCSV(path, [][]string{
		{"name", "age", "salary"},
		{"Alice", "25", "5000"},
		{"Bob", "30", "6000"},
		{"Charlie", "35", "7000"},
	})

	q := New(path)
	result := q.FindAll()

	sum, err := result.Sum("salary")
	if err != nil {
		t.Fatalf("Sum failed: %v", err)
	}

	expected := 18000.0
	if sum != expected {
		t.Errorf("Sum = %f, want %f", sum, expected)
	}

	// 测试条件过滤后的求和
	result2 := New(path).Find("age", ">", "25")
	sum2, err := result2.Sum("salary")
	if err != nil {
		t.Fatalf("Sum with condition failed: %v", err)
	}

	expected2 := 13000.0 // Bob + Charlie
	if sum2 != expected2 {
		t.Errorf("Sum with condition = %f, want %f", sum2, expected2)
	}
}

func TestAggregate_Avg(t *testing.T) {
	dir, err := os.MkdirTemp("", "csvdb_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.csv")
	testutil.CreateCSV(path, [][]string{
		{"name", "age", "salary"},
		{"Alice", "25", "5000"},
		{"Bob", "30", "6000"},
		{"Charlie", "35", "7000"},
	})

	q := New(path)
	result := q.FindAll()

	avg, err := result.Avg("age")
	if err != nil {
		t.Fatalf("Avg failed: %v", err)
	}

	expected := 30.0
	if avg != expected {
		t.Errorf("Avg = %f, want %f", avg, expected)
	}

	// 测试空结果
	result2 := New(path).Find("age", ">", "100")
	avg2, err := result2.Avg("age")
	if err != nil {
		t.Fatalf("Avg on empty result failed: %v", err)
	}

	if avg2 != 0 {
		t.Errorf("Avg on empty = %f, want 0", avg2)
	}
}

func TestAggregate_MinMax(t *testing.T) {
	dir, err := os.MkdirTemp("", "csvdb_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.csv")
	testutil.CreateCSV(path, [][]string{
		{"name", "age", "salary"},
		{"Alice", "25", "5000"},
		{"Bob", "30", "6000"},
		{"Charlie", "35", "7000"},
	})

	q := New(path)
	result := q.FindAll()

	// 测试字符串 Min/Max
	minName, err := result.Min("name")
	if err != nil {
		t.Fatalf("Min failed: %v", err)
	}
	if minName != "Alice" {
		t.Errorf("Min = %s, want Alice", minName)
	}

	maxName, err := result.Max("name")
	if err != nil {
		t.Fatalf("Max failed: %v", err)
	}
	if maxName != "Charlie" {
		t.Errorf("Max = %s, want Charlie", maxName)
	}

	// 测试数值 Min/Max
	minAge, err := result.MinFloat("age")
	if err != nil {
		t.Fatalf("MinFloat failed: %v", err)
	}
	if minAge != 25 {
		t.Errorf("MinFloat = %f, want 25", minAge)
	}

	maxAge, err := result.MaxFloat("age")
	if err != nil {
		t.Fatalf("MaxFloat failed: %v", err)
	}
	if maxAge != 35 {
		t.Errorf("MaxFloat = %f, want 35", maxAge)
	}
}

func TestAggregate_ColumnNotFound(t *testing.T) {
	dir, err := os.MkdirTemp("", "csvdb_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.csv")
	testutil.CreateCSV(path, [][]string{
		{"name", "age"},
		{"Alice", "25"},
	})

	q := New(path)
	result := q.FindAll()

	_, err = result.Sum("nonexistent")
	if err != ErrColumnNotFound {
		t.Errorf("Sum with nonexistent column: err = %v, want %v", err, ErrColumnNotFound)
	}

	_, err = result.Avg("nonexistent")
	if err != ErrColumnNotFound {
		t.Errorf("Avg with nonexistent column: err = %v, want %v", err, ErrColumnNotFound)
	}

	_, err = result.Min("nonexistent")
	if err != ErrColumnNotFound {
		t.Errorf("Min with nonexistent column: err = %v, want %v", err, ErrColumnNotFound)
	}

	_, err = result.Max("nonexistent")
	if err != ErrColumnNotFound {
		t.Errorf("Max with nonexistent column: err = %v, want %v", err, ErrColumnNotFound)
	}
}

func TestAggregate_WithNonNumericValues(t *testing.T) {
	dir, err := os.MkdirTemp("", "csvdb_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.csv")
	testutil.CreateCSV(path, [][]string{
		{"name", "value"},
		{"Alice", "100"},
		{"Bob", "not_a_number"},
		{"Charlie", "200"},
	})

	q := New(path)
	result := q.FindAll()

	// Sum 应该跳过非数字值
	sum, err := result.Sum("value")
	if err != nil {
		t.Fatalf("Sum failed: %v", err)
	}

	expected := 300.0 // 100 + 200, 跳过 "not_a_number"
	if sum != expected {
		t.Errorf("Sum = %f, want %f", sum, expected)
	}

	// Avg 应该只计算有效数字
	avg, err := result.Avg("value")
	if err != nil {
		t.Fatalf("Avg failed: %v", err)
	}

	expectedAvg := 150.0 // (100 + 200) / 2
	if avg != expectedAvg {
		t.Errorf("Avg = %f, want %f", avg, expectedAvg)
	}
}

func TestAggregate_WithLimit(t *testing.T) {
	dir, err := os.MkdirTemp("", "csvdb_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "test.csv")
	testutil.CreateCSV(path, [][]string{
		{"name", "value"},
		{"A", "10"},
		{"B", "20"},
		{"C", "30"},
		{"D", "40"},
	})

	q := New(path)
	result := q.FindAll().Limit(2)

	sum, err := result.Sum("value")
	if err != nil {
		t.Fatalf("Sum with limit failed: %v", err)
	}

	expected := 30.0 // 10 + 20
	if sum != expected {
		t.Errorf("Sum with limit = %f, want %f", sum, expected)
	}
}

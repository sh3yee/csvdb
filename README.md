# csvdb

像操作数据库一样操作 CSV 文件的 Go 库。

## 安装

```bash
go get github.com/sh3yee/csvdb
```

## 导入方式

本项目主要通过子包使用：

```go
import (
    "github.com/sh3yee/csvdb/column"
    "github.com/sh3yee/csvdb/row"
    "github.com/sh3yee/csvdb/query"
)
```

说明：根路径 `github.com/sh3yee/csvdb` 主要用于模块索引与文档入口，实际功能请从子包导入。

如果你使用 gopkg.in，可使用：

```bash
go get gopkg.in/sh3yee/csvdb.v1
```

## 快速开始

```go
package main

import (
    "github.com/sh3yee/csvdb/column"
    "github.com/sh3yee/csvdb/row"
)

func main() {
    // 列操作
    col := column.New("users.csv")
    col.Add("email")
    col.Alter("name", "username")
    col.DeleteByName("age")

    // 行操作
    r := row.New("users.csv")
    r.Add([]string{"1", "tom", "tom@example.com"})
    r.Update(0, []string{"1", "tommy", "tommy@example.com"})
    r.Delete(0)
}
```

## API

### Column - 列操作

#### 添加列

| 方法 | 说明 |
|------|------|
| `Add(field string) error` | 添加空列到末尾 |
| `AddAt(field string, index int) error` | 添加空列到指定位置 |
| `AddWithDefault(field, defaultValue string) error` | 添加带默认值的列到末尾 |
| `AddAtWithDefault(field, defaultValue string, index int) error` | 添加带默认值的列到指定位置 |

#### 修改列

| 方法 | 说明 |
|------|------|
| `Alter(oldName, newName string) error` | 根据列名修改列名 |
| `AlterByIndex(index int, newName string) error` | 根据索引修改列名 |

#### 删除列

| 方法 | 说明 |
|------|------|
| `DeleteByName(name string) error` | 根据列名删除 |
| `DeleteByIndex(index int) error` | 根据索引删除 |

### Row - 行操作

#### 添加行

| 方法 | 说明 |
|------|------|
| `Add(values []string) error` | 添加行到末尾 |
| `AddAt(values []string, index int) error` | 添加行到指定位置 |

#### 修改行

| 方法 | 说明 |
|------|------|
| `Update(index int, values []string) error` | 更新指定索引的行 |
| `UpdateBy(column, value string, newValues []string) error` | 更新匹配条件的行 |

#### 删除行

| 方法 | 说明 |
|------|------|
| `Delete(index int) error` | 删除指定索引的行 |
| `DeleteBy(column, value string) error` | 删除匹配条件的行 |

#### 查询行

| 方法 | 说明 |
|------|------|
| `Get(index int) ([]string, error)` | 获取指定索引的行 |
| `GetBy(column, value string) ([][]string, error)` | 获取匹配条件的行 |
| `GetAll() ([][]string, error)` | 获取所有行 |

### Query - 查询操作

#### 条件查询

| 方法 | 说明 |
|------|------|
| `Find(cond Condition) ([][]string, error)` | 单条件查询 |
| `FindAll(conds ...Condition) *Result` | 多条件查询，返回 Result 支持链式处理 |
| `FindIn(column string, values []string) ([][]string, error)` | IN 查询 |
| `FindNotIn(column string, values []string) ([][]string, error)` | NOT IN 查询 |

#### Condition 结构体

```go
type Condition struct {
    Column string
    Op     string // "=", "!=", ">", "<", ">=", "<=", "like"
    Value  string
}
```

#### Result 链式方法

| 方法 | 说明 |
|------|------|
| `Select(columns ...string) *Result` | 选择特定列 |
| `OrderBy(column, order string) *Result` | 单列排序，order 为 "asc" 或 "desc" |
| `ThenBy(column, order string) *Result` | 多列排序，追加排序条件 |
| `Limit(n int) *Result` | 限制结果数量 |
| `Offset(n int) *Result` | 跳过前 N 条 |
| `Get() ([][]string, error)` | 获取所有结果 |
| `First() ([]string, error)` | 获取第一条 |
| `Count() (int, error)` | 统计数量 |
| `Exists() (bool, error)` | 判断是否存在 |

#### 查询示例

```go
q := query.New("users.csv")

// 单条件查询
rows, _ := q.Find(query.Condition{Column: "age", Op: ">", Value: "18"})

// 多条件查询 + 选列 + 分页
rows, _ := q.FindAll(
    query.Condition{Column: "age", Op: ">=", Value: "18"},
    query.Condition{Column: "city", Op: "=", Value: "Beijing"},
).Select("name", "email").Limit(10).Offset(5).Get()

// IN 查询
rows, _ := q.FindIn("id", []string{"1", "2", "3"})

// 模糊匹配
rows, _ := q.Find(query.Condition{Column: "name", Op: "like", Value: "tom%"})

// 排序
rows, _ := q.FindAll().OrderBy("age", "desc").Get()

// 多列排序
rows, _ := q.FindAll().OrderBy("age", "asc").ThenBy("name", "desc").Get()

// 查询 + 排序 + 分页
rows, _ := q.FindAll(
    query.Condition{Column: "status", Op: "=", Value: "active"},
).OrderBy("created_at", "desc").Limit(10).Get()

// 判断是否存在
exists, _ := q.Find(query.Condition{Column: "id", Op: "=", Value: "1"}).Exists()
```

#### 聚合统计

| 方法 | 说明 |
|------|------|
| `Count() (int, error)` | 统计数量 |
| `Sum(column string) (float64, error)` | 计算指定列的总和 |
| `Avg(column string) (float64, error)` | 计算指定列的平均值 |
| `Min(column string) (string, error)` | 获取指定列的最小值（字符串比较） |
| `Max(column string) (string, error)` | 获取指定列的最大值（字符串比较） |
| `MinFloat(column string) (float64, error)` | 获取指定列的最小值（数值比较） |
| `MaxFloat(column string) (float64, error)` | 获取指定列的最大值（数值比较） |

#### 聚合示例

```go
q := query.New("users.csv")

// 统计数量
count, _ := q.FindAll().Count()

// 总和
total, _ := q.FindAll().Sum("amount")

// 平均值
avg, _ := q.FindAll().Avg("age")

// 最小值/最大值（字符串比较）
minName, _ := q.FindAll().Min("name")
maxName, _ := q.FindAll().Max("name")

// 最小值/最大值（数值比较）
minAge, _ := q.FindAll().MinFloat("age")
maxAge, _ := q.FindAll().MaxFloat("age")

// 条件过滤后聚合
total, _ := q.FindAll(
    query.Condition{Column: "status", Op: "=", Value: "active"},
).Sum("amount")
```

### 错误类型

| 错误 | 说明 |
|------|------|
| `ErrColumnNotFound` | 列不存在 |
| `ErrRowNotFound` | 行不存在 |
| `ErrIndexOutOfRange` | 索引越界 |

## 待开发功能

| 模块 | 功能 | 状态 |
|------|------|------|
| Row | 行操作（增删改查） | ✅ 已完成 |
| Query | 条件查询、筛选 | ✅ 已完成 |
| Sort | 排序功能 | ✅ 已完成 |
| Aggregate | 聚合统计（COUNT/SUM/AVG/MIN/MAX） | ✅ 已完成 |

## License

[MIT](LICENSE)

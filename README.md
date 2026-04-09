# csvdb

像操作数据库一样操作 CSV 文件的 Go 库。

## 安装

```bash
go get gycsv
```

## 快速开始

```go
package main

import (
    "gycsv/column"
    "gycsv/row"
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
| Query | 条件查询、筛选 | 📌 待开发 |
| Sort | 排序功能 | ⏳ 计划中 |
| Aggregate | 聚合统计（COUNT/SUM/AVG/MIN/MAX） | ⏳ 计划中 |
| Join | 多 CSV 文件关联 | ⏳ 计划中 |

## License

[MIT](LICENSE)

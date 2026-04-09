# csvdb

像操作数据库一样操作 CSV 文件的 Go 库。

## 安装

```bash
go get gycsv
```

## 快速开始

```go
package main

import "gycsv/column"

func main() {
    col := column.New("users.csv")

    // 添加列
    col.Add("email")
    col.AddAt("id", 0)
    col.AddWithDefault("status", "active")

    // 修改列名
    col.Alter("name", "username")

    // 删除列
    col.DeleteByName("age")
    col.DeleteByIndex(0)
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

### 错误类型

| 错误 | 说明 |
|------|------|
| `ErrColumnNotFound` | 列不存在 |
| `ErrIndexOutOfRange` | 索引越界 |

## 待开发功能

| 模块 | 功能 | 状态 |
|------|------|------|
| Row | 行操作（增删改查） | 📌 待开发 |
| Query | 条件查询、筛选 | ⏳ 计划中 |
| Sort | 排序功能 | ⏳ 计划中 |
| Aggregate | 聚合统计（COUNT/SUM/AVG/MIN/MAX） | ⏳ 计划中 |
| Join | 多 CSV 文件关联 | ⏳ 计划中 |

## License

[MIT](LICENSE)

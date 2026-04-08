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
| `Alter(before, after string) error` | 修改列名 |

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

## License

[MIT](LICENSE)

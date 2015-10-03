# 基础类型

- 布尔类型：bool
- 整型：int8、byte、int16、int、uint、uintptr等
- 浮点类型：float32、float64
- 复数类型：complex64、complex128
- 字符串：string
- 字符类型：rune
- 错误类型：error

# 复合类型

- 指针：pointer
- 数组：array
- 切片：slice
- 字典：map
- 通道：chan
- 结构体：struct
- 接口：interface

# 常用命令

- ` godoc --http=:6060 ` 开启本地文档，然后访问 http://localhost:6060
- ` go run main.go ` 运行一个go程序
- ` go run --work main.go ` 运行的时候将生成的临时文件路径打印出来
- ` go build main.go ` 构建go程序并生成可执行文件

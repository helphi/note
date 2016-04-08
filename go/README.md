# 链接

- 官网 <http://golang.org/>
- 在线运行go代码 <http://play.golang.org/>
- 下载 <https://golang.org/dl/>
- go on github <https://github.com/golang/go>
- go doc <https://godoc.org/>
- go语言中文网 <http://studygolang.com/> // 提供下载及文档
- golang中国 <http://www.golangtc.com/> // 提供下载及文档
- gobyexample <https://gobyexample.com/>
- go数据库开发 <http://go-database-sql.org/>

# 同时安装多套环境

```sh
mkdir d:%HOMEPATH%\gosrc
mkdir d:%HOMEPATH%\gopath-1.3.1
mkdir d:%HOMEPATH%\gopath-1.4.2
mklink /j d:%HOMEPATH%\gopath-1.3.1\src d:%HOMEPATH%\gosrc
mklink /j d:%HOMEPATH%\gopath-1.4.2\src d:%HOMEPATH%\gosrc
```

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

# 整型

类型 | 长度（字节） | 值范围
-----|--------------|-------
int8 | 1 |  -128 ~ 127
uint8（即byte） | 1 |  0 ~ 255
int16 | 2 |  -32 768 ~ 32 767
uint16 | 2 |  0 ~ 65 535
int32 | 4 |  -2 147 483 648 ~ 2 147 483 647
uint32 | 4 |  0 ~ 4 294 967 295
int64 | 8 |  -9 223 372 036 854 775 808 ~ 9 223 372 036 854 775 807
uint64 | 8 |  0 ~ 18 446 744 073 709 551 615
int | 平台相关 | 平台相关
uint | 平台相关 | 平台相关
uintptr | 同指针 | 在32位平台下为4字节，64位平台下为8字节

# 常用命令

- ` godoc --http=:6060 ` 开启本地文档，然后访问 http://localhost:6060
- ` go run main.go ` 运行一个go程序
- ` go run --work main.go ` 运行的时候将生成的临时文件路径打印出来
- ` go build main.go ` 构建go程序并生成可执行文件

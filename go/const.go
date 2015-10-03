package main

func main() {
  const Pi float64 = 3.14159265358979323846
  const zero = 0.0 // 无类型浮点常量

  const (
    size int64 = 1024
    eof = -1 // 无类型整型常量
  )

  const m, n float32 = 0, 3 // m = 0.0, n = 3.0，常量的多重赋值
  const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量
  const mask = 1 << 3
  // const Home = os.Getenv("HOME") // ERROR: const initializer os.Getenv("HOME") is not a constant

  const ( // iota被重设为0
    c0 = iota // c0 == 0
    c1 = iota // c1 == 1
    c2 = iota // c2 == 2
  )

  const ( // iota被重设为0
    c00 = iota // c00 == 0
    c11 // c11 == 1
    c22 // c22 == 2
  )

  const (
    c3 = 1 << iota // c3 == 1 (iota在每个const开头被重设为0)
    c4 = 1 << iota // c4 == 2
    c5 = 1 << iota // c5 == 4
  )

  const (
    c33 = 1 <<iota // c33 == 1 (iota在每个const开头被重设为0)
    c44 // c44 == 2
    c55 // c55 == 4
  )

  const (
    u = iota * 42 // u == 0
    v float64 = iota * 42 // v == 42.0
    w = iota * 42 // w == 84
  )

  const x = iota // x == 0 (因为iota又被重设为0了)
  const y = iota // y == 0 (同上)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出，包外不可见
	)
}

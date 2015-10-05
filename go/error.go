package main

import "fmt"
import "os"
import "io"

func main() {
	_, err := os.Stat("a.txt")
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			fmt.Println("error")
		}
	}

	CopyFile("", "")

	ErrorFlow()
}

type PathError struct {
	Op string
	Path string
	Err error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)

	defer func() {
		fmt.Println("srcFile close")
		srcFile.Close()
	}()

	dstFile, err := os.Create(dst)

	defer func() {
		fmt.Println("dstFile close")
		dstFile.Close()
	}()

	return io.Copy(dstFile, srcFile)
}

func ErrorFlow() {
	fmt.Println("ErrorFlow start")
	defer func() {
		fmt.Println("ErrorFlow defer")
		if r := recover(); r != nil {
			fmt.Printf("ErrorFlow recover:%v\n", r)
		}
	}()

	inErrorFlow()
	fmt.Println("ErrorFlow end")
}

func inErrorFlow() {
	fmt.Println("inErrorFlow start")
	_, err := os.Open("")
	if err != nil {
		panic(err)
	}
	fmt.Println("inErrorFlow end")
}

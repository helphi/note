# 在linux中构建与执行


```bash
export GOPATH=$GOPATH:$PWD
mkdir bin
cd bin
go test algorithm/qsort 
go test algorithm/bubblesort 
go build sorter
./sorter -i unsorted.dat -o sorted.dat -a bubblesort
```

# 在windows中构建与执行

```cmd
set GOPATH=%GOPATH%;%CD%
mkdir bin
cd bin
go test algorithm/qsort 
go test algorithm/bubblesort 
go build sorter
sorter.exe -i unsorted.dat -o sorted.dat -a bubblesort
```

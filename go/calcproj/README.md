# 在linux中构建与执行


```bash
export GOPATH=$GOPATH:$PWD
mkdir bin
cd bin
go test simplemath
go build calc
./calc add 5 6
./calc sqrt 9

```

# 在windows中构建与执行

```cmd
set GOPATH=%GOPATH%;%CD%
mkdir bin
cd bin
go test simplemath
go build calc
calc.exe add 5 6
calc.exe sqrt 9
```

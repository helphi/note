set GOROOT=D:\dev\go1.4.2.windows-amd64\go
set GOPATH=%HOMEDRIVE%%HOMEPATH%\gopath
set Path=%Path%;%GOROOT%\bin
echo "go run main.go | go run --work main.go | go build main.go"
cmd

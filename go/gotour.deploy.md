# 中文版 gotour 部署
---------------------------

### 下载相关文件

```sh
curl -O ftp://ftp.cqrd.com/software/google/go/go1.5.1.linux-amd64.tar.gz
curl -O ftp://ftp.cqrd.com/software/google/go/gotour/go-tour-zh.zip
curl -O ftp://ftp.cqrd.com/software/google/go/gotour/net.zip
curl -O ftp://ftp.cqrd.com/software/google/go/gotour/tools.zip
```

### 安装 golang 运行环境

```sh
tar xzf go1.5.1.linux-amd64.tar.gz

vi ~/.bashrc

export GOROOT=/opt/app/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/opt/app/gopath

. ~/.bashrc
```

### 安装 gotour

```sh
mkdir -p /opt/app/gopath/src/bitbucket.org/mikespook
mkdir -p /opt/app/gopath/src/golang.org/x/

unzip go-tour-zh.zip -d /opt/app/gopath/src/bitbucket.org/mikespook
unzip net.zip -d /opt/app/gopath/src/golang.org/x/
unzip tools.zip -d /opt/app/gopath/src/golang.org/x/

mv /opt/app/gopath/src/bitbucket.org/mkiespook/mikespook-go-tour-zh-d850789fb984/ /opt/app/gopath/src/bitbucket.org/mikespook/go-tour-zh
mv /opt/app/gopath/src/golang.org/x/net-master/ /opt/app/gopath/src/golang.org/x/net
mv /opt/app/gopath/src/golang.org/x/tools-master/ /opt/app/gopath/src/golang.org/x/tools

go install bitbucket.org/mikespook/go-tour-zh/gotour
```

### 启动 gotour

```sh
/opt/app/gopath/bin/gotour
```

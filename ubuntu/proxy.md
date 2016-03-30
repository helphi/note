脚本如下

```sh
export http_proxy=http://IP:PORT
export https_proxy=http://IP:PORT
export ftp_proxy=http://IP:PORT
export no_proxy=localhost,127.0.0.1,192.168.,172.,10.
```

- 如果有用户名密码使用 `http://user:pwd@IP:PORT` 格式
- 直接在终端执行只在该终端有效
- 在 `~/.bashrc` 文件末尾加入以上指令，然后执行 `. ~/.bashrc` 可一直生效
- 在 `/etc/apt/apt.conf` 文件末尾加入以上指令只对 `apt-*` 指令生效

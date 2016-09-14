# AIX系统相关命令

- `hostid` 显示主机ID号。
- `isainfo` 显示系统内核支持的应用程序的比特位数。
- `isalist` 显示在x86平台上系统支持应用程序的比特位数。
- `localeadm` 进行系统的时区设置。
- `prtconf` 列出系统硬件信息。
- `psrinfo` 显示CPU的类型。
- `showrev` 显示主机名、主机ID、内核版本、应用程序架构、硬件提供者信息等。
- `uname` 显示操作系统的名称、版本、节点名、硬件名和CPU类型。

# 根据占用的端口号查看进程号

* 首先使用 `nestat -Ant | grep 端口号` 获得该端口号的PCB（Protocol control block），组查看协议tcp。（输出信息中第一列既是PCB）
* 如果是tcp连接，那么用 `rmsock <PCB> tcpcb`
* 如果是udp连接，那么用 `rmsock <PCB> inpcb`

**例：**

```bash
# netstat -Ant | grep 8000
f1000e000288c3b8 tcp4 0 0 127.0.0.1.8000 *.* LISTEN
# rmsock f1000e000288c3b8 tcpcb
The socket 0xf1000e000288c008 is being held by proccess 10551550 (abc).
```
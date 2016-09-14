# 根据 ip 找 mac 地址

```bash
arp -d *
ping $ip
arp -a
```

> arp 指令只能查找本网段的 mac 地址

# 根据 ip 查主机名

- `nbtstat -A $ip`
- `ping -a $ip`

> linux 使用 `nmblookup -A $ip`

# ping 一个网段的 ip

```bash
for /l %%i in (2,1,254) do ping -n 1 -w 1 172.26.3.%%i
```

# 删除某个目录下所有的 `.svn` 文件夹

```bash
for /r %%d in (.) do if exist "%%d\.svn" rd /s /q "%%d\.svn"
```
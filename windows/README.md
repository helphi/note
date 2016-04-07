# 根据ip找mac地址

```sh
arp -d *
ping $ip
arp -a
```

> arp 指令只能查找本网段的 mac 地址

# 根据ip查主机名

- `nbtstat -A $ip`
- `ping -a $ip`

> linux 使用 `nmblookup -A $ip`

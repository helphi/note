# 网络文件夹共享

- 服务器端配置

服务器ip=192.168.1.2

```bash
vi /etc/exports
/share 192.168.1.2(insecure,rw,sync,no_root_squash)
service rpcbind start
service nfs restart
```

- 客户端配置

客户端ip=192.168.1.2

```bash
mount -t nfs 192.168.1.1:/share /share
```

- 查看

```bash
fuser -m -v /share
```
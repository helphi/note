# install

```sh
wget https://get.daocloud.io/docker/builds/Linux/x86_64/docker-1.10.3
chmod +x docker-1.10.3
mv docker-1.10.3 /usr/local/bin/docker
docker daemon
```

- 其他版本

```sh
wget https://get.docker.com/builds/Linux/x86_64/docker-1.9.1
```

- 常用启动指令 `http_proxy=http://172.17.18.84:8080 https_proxy=http://172.17.18.84:8080 docker --registry-mirror=https://6pzhi4th.mirror.aliyuncs.com daemon --bip=172.172.172.1/24`

- 错误处理1：
```
FATA[0000] Error mounting devices cgroup: mountpoint for devices not found
```

这个错误为cgroup在宿主机上没有挂载，在 `/etc/fstab` 结尾添加如下代码然后重启

```
none        /sys/fs/cgroup        cgroup        defaults    0    0
```

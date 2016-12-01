# 在线安装

```bash
curl -sSL https://get.docker.com/ | sh
```

# 离线安装

```bash
wget https://get.daocloud.io/docker/builds/Linux/x86_64/docker-1.12.3.tgz
tar zxf docker-1.12.3.tgz
sudo mv docker/* /usr/local/bin/
```

* 其他版本

```bash
wget https://get.daocloud.io/docker/builds/Linux/x86_64/docker-1.11.2.tgz
wget https://get.daocloud.io/docker/builds/Linux/x86_64/docker-1.10.3
wget https://get.docker.com/builds/Linux/x86_64/docker-1.9.1
```

* 常见错误：

```
FATA[0000] Error mounting devices cgroup: mountpoint for devices not found
```

这个错误为cgroup在宿主机上没有挂载，在 `/etc/fstab` 结尾添加如下代码然后重启

```bash
echo "none /sys/fs/cgroup cgroup defaults 0 0" >> /etc/fstab
```

# 启动

* 启动使用代理并指定容器网段 `all_proxy=http://proxyserver:port dockerd`
* 启动时使用镜像 `dockerd --registry-mirror=https://6pzhi4th.mirror.aliyuncs.com` 其中 `--registry-mirror` 可以指定多个
* 示例

```bash
sudo su
groupadd docker
usermod -aG docker hefei
dockerd --registry-mirror=https://6pzhi4th.mirror.aliyuncs.com --bip=172.172.172.1/24 > /var/log/docker 2>&1 &
```

# 常用指令

- 从非官方库拉取镜像
```bash
docker pull 6pzhi4th.mirror.aliyuncs.com/library/ubuntu:14.04
```
- 删除所有已停止容器
```bash
docker rm $(docker ps -aq)
```
- 删除所有 `<none>` 镜像 
```bash
docker rmi $(docker images -qf dangling=true)
```
- 删除所有未在使用的数据卷
```bash
docker volume rm $(docker volume ls -qf dangling=true)
```
- 进入容器 
```bash
docker exec -it ID bash
```


# ubuntu 上安装 git

```bash
sudo add-apt-repository ppa:git-core/ppa # 使用最新的 git 源安装最新的 git，不执行这句则使用默认源安装固定的版本
sudo apt-get update
sudo apt-get install git
```

# 常用命令

- 代理设置 `git config http.proxy http://ip:port`
- 全局代理设置 `git config --global http.proxy http://ip:port`
- 单库代理设置 `git config remote.origin.proxy http://ip:port` // `origin` 可替换为其他远程库名字

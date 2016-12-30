# 常用

- Ubuntu 16.04 Unity 启动器移动到桌面底部 `gsettings set com.canonical.Unity.Launcher launcher-position Bottom`
- Ubuntu 16.04 Unity 启动器移动到桌面 `gsettings set com.canonical.Unity.Launcher launcher-position Left`
- 安装媒体解码器 `sudo apt install ubuntu-restricted-extras`
- 安装常用软件 `sudo apt-get install git vim openssh-server wps-office`
- 16.04 安装搜狗输入法 
```bash
sudo echo "deb http://archive.ubuntukylin.com:10006/ubuntukylin xenial main" > /etc/apt/sources.list.d/ubuntukylin.list
sudo apt update && apt install sogoupinyin  
```
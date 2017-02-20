# 常用

- Ubuntu 16.04 Unity 启动器移动到桌面底部 `gsettings set com.canonical.Unity.Launcher launcher-position Bottom`
- Ubuntu 16.04 Unity 启动器移动到桌面 `gsettings set com.canonical.Unity.Launcher launcher-position Left`
- 软件卸载 `sudo apt remove libreoffice-common unity-webapps-common fonts-noto-cjk`
- 安装媒体解码器 `sudo apt install ubuntu-restricted-extras`
- 安装常用软件 `sudo apt-get install git vim openssh-server wps-office unrar ttf-wqy-microhei`
- 16.04 安装搜狗输入法 
```bash
sudo echo "deb http://archive.ubuntukylin.com:10006/ubuntukylin xenial main" > /etc/apt/sources.list.d/ubuntukylin.list
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys D259B7555E1D3C58
sudo apt update && apt install sogoupinyin  
```
- 允许远程桌面连接
```sh
gsettings set org.gnome.Vino enabled true
gsettings set org.gnome.Vino require-encryption false
```

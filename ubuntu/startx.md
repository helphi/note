# ubuntu 14.04

文件 `/etc/default/grub` 中以下两行分别表示开机启动图形界面和文字界面

```
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash" 
GRUB_CMDLINE_LINUX_DEFAULT="text" 
```
 
> Tips:
> - 修改文件后需要使用 `sudo update-grub` 使配置生效
> - 在文字界面使用 `sudo su -` 和 `service lightdm start` 启动图形界面，注意不要在普通用户 home 目录下执行 starx 否则会让该普通用户 home 目录下的 .Xauthority 所有者变成 root 从而造成该普通用户无法登陆图形界面

# ubuntu 16.04

- 禁用图形界面开机启动： `sudo rm -rf /etc/systemd/system/display-manager.service`
- 恢复： `sudo ln -s /lib/systemd/system/lightdm.service /etc/systemd/system/display-manager.service`
文件 `/etc/default/grub` 中以下两行分别表示开机启动图形界面和文字界面

```
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash" 
GRUB_CMDLINE_LINUX_DEFAULT="text" 
```
 
> Tips:
> 修改文件后需要使用 `sudo update-grub` 使配置生效
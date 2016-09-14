# 修复 grub

很多时候，特别是在linux调整分区后，开机重启时会出现如下字样，系统就是进不去：

```bash
error : unknow filesystem
grub rescue>
```

这一般是由于分区调整或分区UUID改变造成grub2不能正常启动，从而进入修复模式了（grub rescue)，也称救援模式，在救援模式下只有很少的命令可以用：set, ls, insmod, root, prefix
 
1. `set`  查看环境变量，这里可以查看启动路径和分区。
1. `ls`   查看设备
1. `insmod`  加载模块
1. `root`  指定用于启动系统的分区,在救援模式下设置grub启动分区
1. `prefix` 设定grub启动路径
 
**修复步骤：**

1、查看分区

```bash
grub rescue> ls
(hd0) (hd0,msdos9) (hd0,msdos8) (hd0,msdos7) (hd0,msdos6) (hd0,msdos5) (hd0,msdos2) (hd0,msdos1)
```

> 以上结果每个人可能不一样

2、寻找ubuntu所在分区

```bash
grub rescue> ls (hd0,msdos1)/
```

> 若出现 unknown filesystem 字样，则尝试下一个，若出现的是你的 ubuntu 主文件夹下的文件夹和文件的名字，那就是的要找的分区了
 
3、修改启动分区

假如你找到的启动分区是 `hd0,msdos8`

```bash
grub rescue>root=(hd0,msdos8)
grub rescue>prefix=/boot/grub                         //grub路径设置
grub rescue>set root=(hd0,msdos8)
grub rescue>set prefix=(hd0,msdos8)/boot/grub
grub rescue>insmod normal                            //启动normal启动
grub rescue>normal
```
 
4、进入命令行启动 ubuntu

进入系统启动选项界面后还是进不去，因为还没有真正的修改grub，这个要到ubuntu里修改
进入系统启动项界面后，按 C 进入命令行模式

```bash
grub >set root=hd0,msdos8
grub >set prefix=(hd0,msdos8)/boot/grub
grub >linux /vmlinuz-xxx-xxx root=/dev/sda8 //里边的 xxxx 可以按 Tab 键，如果有 acpi 问题,在最后加一句 acpi=off
grub >initrd /initrd.img-xxx-xxx
grub >boot
``` 

5、进入ubuntu修复grub

```bash
sudo update-grub
sudo grub-install /dev/sda
```

6、重启
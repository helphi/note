# vim 添加到 windows 右键菜单

## 使用自帶的安裝文件
1. 执行安装目录下的 install.exe
1. 使用(d 14)执行第14个脚本


## 修改注册表

1. 删掉注册表中的 `HKEY_CLASSES_ROOT/*/shellex/ContextMenuHandlers/gvim`
1. 直接进入注册表，在 `HKEY_CLASSES_ROOT/*/shell` 下添加项 `用 vim 编辑`， 再在项 `用 vim 编辑` 下添加子项 `command` [右键-新建-项] ，在其右边窗口把其键值设定为 `"D:/Program Files/vim/vim72/gvim.exe" -p --remote-tab-silent "%1"`

> 也可把下面的内容保存为文件 gvim.reg，并将之导入注册表，注意 vim 程序的路径应为实际路径

```reg
Windows Registry Editor Version 5.00
[HKEY_CLASSES_ROOT/*/shell/Edit with &Vim]
[HKEY_CLASSES_ROOT/*/shell/Edit with &Vim/command]
@="/"D://Program Files//Vim//vim72//gvim.exe/" -p --remote-tab-silent /"%1/""
```
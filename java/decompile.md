# 介绍

Java 反编译工具有很多种，以下介绍其中一些工具

# jad

jad 是一款使用非常广泛地 Java 反编译工具，官网 <http://www.varaneckas.com/jad> 提供下载

### 批量反编译 .class 示例

`jad.exe -r -ff -d src -s java classes/**/*.class`

> 其中 `-r` 表示恢复源文件的目录结构，`-ff` 表示将类属性定义放在类方法定义之前， `-d` 表示输出目录，`-s` 表示文件的扩展名，更多参数可直接执行 `java.exe` 或查看 根目录下的 `Readme.txt`

### 链接

- 360盘 `所有文件\software\jad`

# jadclipse

JadClipse 是基于上面介绍的 jad 的 eclipse 插件，是一款非常实用而且方便地 Java 反编译插件

### 使用

1. 从官网 <https://sourceforge.net/projects/jadclipse/> 下载 jadClipse 的 jar 包到 eclipse 根目录下的 `dropins` 目录中（没有这个目录则新建）
1. 重启 eclipse 导航到 `Windows` -> `Preferences` -> `Java` -> `jadClipse` （如果没有出现这一项，删除 eclipse 根目录下的 `configuration/org.eclipse.update` 文件夹后重启 eclipse）
1. 在 `Path to decompiler` 中输入下载的 jad 的路径，如 `D:\dev\jad158g.win\jad.exe`
1. 导航到 `Windows` -> `Perference` -> `General` -> `Editors` -> `File Associations` 将 `*.class` 和  `*.class without source` 默认编辑器都设置为 `JadClipse Class File Viewer`
1. 然后就可以直接打开 .class 文件进行查看了

### 链接

- 360盘 `所有文件\software\eclipse-plugins\jadclipse`

# JD-GUI

JD-GUI 是一个用 C++ 开发的 Java 反编译工具，由 Pavel Kouznetsov 开发，支持Windows、Linux 和 Mac OS 三个平台。而且提供了 Eclipse 插件 JD-Eclipse 和 IntelliJ 插件 JD-IntelliJ。JD-GUI不需要安装，直接点击运行，可以反编译 jar 和 class 文件。

### 链接

- 官网 <http://jd.benow.ca/>
- GitHub <https://github.com/java-decompiler/jd-gui>
- 360盘 `所有文件\software\jad`

# JD-Eclipse

### 安装（摘自官网）

1. Download and unzip the JD-Eclipse Update Site,
1. Launch Eclipse,
1. Click on "Help > Install New Software...",
1. Click on button "Add..." to add an new repository,
1. Enter "JD-Eclipse Update Site" and select the local site directory,
1. Check "Java Decompiler Eclipse Plug-in",
1. Next, next, next... and restart Eclipse.

### 链接

- 官网 <http://jd.benow.ca/>
- GitHub <https://github.com/java-decompiler/jd-eclipse>
- 360盘 `所有文件\software\eclipse-plugins\jd-eclipse`

- eclipse 提供在线安装插件的机制，但是在没有网络的情况下我们可以用以下几种方式手动安装插件，其中前三种方式并不是每种插件都适用的，比如有些插件涉及到版本兼容，依赖，额外处理等内容就可能安装不成功
- 从官方下载的离线更新包里面一般都有 `content.jar`、`artifacts.jar`、`site.xml` 等文件，这种包最好用 eclipse 插件安装向导进行安装，用其他方法安装可能会成功但不推荐
- 有些离线包解压缩后只有 features 和 plugins 两个子目录而没有额外的东西，这种一般使用 dropins 或 links 的方式就可以，如果加载不到则需要直接放到 eclipse 的 features 和 plugins 目录下

# 安装插件到 plugins 和 features 目录下

- features 目录: 插件的配置信息，启动信息,签名，使用权，还有一些图片等一些资源;
- plugins 目录: 真正实现插件功能的代码jar包，还有一些必要的配置信息等。

# 使用 links 方式

比如 jd 反编译插件位置在 `D:\eclipse-plugins\jd` 目录中，插件目录下的结构如下（这种目录结构最好保持，否则可能加载失败）：

```
eclipse     // 目录
-- plugins  // 目录
-- features // 目录
-- site.xml // 这个文件可有可无，不影响效果
```

然后在 eclipse 根目录下建立名为 `links` 的目录，目录下建立一个 `jd.link` 文件（不一定以 link 结尾，也可以是其他后缀），文件内容为 `path=D:\eclipse-plugins\jd`

# 使用 dropins 方式（eclipse3.4以上）

在 eclipse 根目录下建立名为 `dropins` 的目录，以 jd 反编译插件为例，文件存放结构一般如下：

```
jd            // 文件夹的建立非必须，可以直接在 dropins 文件夹下放 jar 包，因为 eclipse 是递归遍历的）
-- eclipse
---- plugins
---- features
---- site.xml // 以 dropins 的方式安装插件该文件最好删除，否则可能加载失败
```

# 使用 eclipse 插件安装向导

打开 eclipse，导航到 `help` -> `install new software`，在 `work with` 中输入下载好的插件安装包地址

- 输入的地址根目录下需要有插件的配置说明文件，比如 `site.xml`、`artifacts.xml`、`content.xml` 等，一般官方下载的插件离线安装包都有类似文件
- 输入的地址可以是一个本地目录，如 `file:/d:/subclipse-site-1.10.9/site-1.10.9/`，也可以是一个远程目录，如 `http://xxx.com/subclipse-1.6/`
- 也可以是一个压缩包，压缩包地址后面需要跟一个感叹号，如 `jar:file:/d:/subclipse-site-1.10.9.zip!/`，如果压缩包中有子目录，还需要根上子目录地址，如 `jar:file:/d:/subclipse-site-1.10.9.zip!/site-1.10.9`
- 压缩包可以是本地地址，也可以是远程地址，如 `jar:http://xxx.com/subclipse-site-1.10.9.zip!/site-1.10.9`

# 插件没有成功加载的处理

1. 删除 eclipse 根目录下的 `configuration\org.eclipse.update` 或者使用 `eclipse -clean` 重启 eclipse
2. 在 `configuration/config.ini` 文件中修改 `org.eclipse.update.reconcile=true` 或者加入一行 `osgi.checkconfiguration=true` 这样它会寻找并安装插件,加载成功后需要将修改过的配置还原回去

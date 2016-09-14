Google Chrome是由Google开发的一款设计简单、高效的Web浏览工具。  

Google Chrome的特点是简洁、快速。GoogleChrome支持多标签浏览，每个标签页面都在独立的“沙箱”内运行，在提高安全性的同时，一个标签页面的崩溃也不会导致其他标签页面被关闭。此外，Google Chrome基于更强大的JavaScript V8引擎，这是当前Web浏览器所无法实现的。   

* Chrome浏览器基于开源引擎WebKit、Blink，其中包含谷歌Gears。
* Chrome提供了浏览器扩展框架，可以制作与Adobe-AIR类似的混合应用。
* 包含V8 Javascript虚拟机，这个多线程的虚拟机可以加速Javascript的执行。
* 具备隐私浏览模式，可以让用户无需在本地机器上登录即可使用，这个功能与微软IE8中的Incognito类似。
* 浏览器将内置防止“网络钓鱼“及恶意软件功能。  

Chrome 官方网址为 <http://www.google.com/chrome>，在这个页面有下载、设置、帮助、论坛等的入口，在这个页面的下载中进行 Chrome 安装需要联网进行在线安装。    
  
Chrome 离线安装发布页为 <https://support.google.com/installer/answer/126299>，这个页面中包含了两个下载离线安装包的链接， <http://www.google.com/chrome/eula.html?standalone=1> 用于安装给自己使用； <http://www.google.com/chrome/eula.html?system=true&standalone=1> 用于安装给计算机上所有用户使用。  

<https://developer.chrome.com/home> 是 Chrome 面向开发者网站，里面有关于 Chrome 的控制台使用、网页调试、插件开发等相关内容。

# 常见问题

**错误** Error code: ssl_error_weak_server_ephemeral_dh_key    
**解决方法**
1. 低版本进入 `chrome://flags`，`支持的最低 SSL/TLS 版本` 改为 `TLS1.0`
1. 高版本：`右键` -> `目标` 加参数 `--cipher-suite-blacklist=0x0039,0x0033`

# 安装三方插件
Chrome 默认只允许从 Chrome Web Store 下载安装扩展程序， 而众多三方或离线下载的扩展程序默认无法安装，提示 ` 无法添加来自此网站的应用、扩展和用户脚本`
 
**解决方法如下**
1. 下载扩展程序到本地
1. 使用解压软件将扩展程序解压到文件夹
1. 打开 `chrome://extensions/`，启用 `开发者模式` ，选择 `加载正在开发的扩展程序...`，然后选择解压出来的文件夹

> 还有些插件提示 `Cannot load extension with file or directory name _metadata. Filenames starting with "_" are reserved for use by the system.` 可以将刚才解压出来的文件夹中的 `_metadata` 文件夹重命名为 `metadata`

# 不让google根据地区自动跳转网址
使用网址 <http://www.google.com/ncr>
> ncr = no contry redirect
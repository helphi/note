# 配置代理

文件 `~/.subversion/servers`

```bash
[global]
# http-proxy-exceptions = *.exception.com, www.internal-site.org
# http-proxy-host = defaultproxy.whatever.com
# http-proxy-port = 7000
# http-proxy-username = defaultusername
# http-proxy-password = defaultpassword
```

# 删除某个目录下所有的 `.svn` 文件夹

```bash
for /r %%d in (.) do if exist "%%d\.svn" rd /s /q "%%d\.svn"
```

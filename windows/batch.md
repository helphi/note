# 常用批处理

```bat
::将文件夹及其子文件夹里所有文件的文件名中的空格都替换成下划线
@echo off&setlocal enabledelayedexpansion
for /f "delims=" %%i in ('dir /a/b/s *.*') do (
cd /d "%%~pi"&set n=%%~nxi
set m=!n: =_!&ren "%%i" "!m!")
pause
::ping 一个网段的 ip
for /l %%i in (2,1,254) do ping -n 1 -w 1 192.168.1.%%i
::删除某个目录下所有的 .svn 文件夹
for /r %%d in (.) do if exist "%%d\.svn" rd /s /q "%%d\.svn"
::设置ip自动获取
netsh interface ip set address "LAN" dhcp
netsh interface ip set dns "LAN" dhcp
netsh interface ip set address "WLAN" dhcp
netsh interface ip set dns "WLAN" dhcp
ipconfig /flushdns
::设置静态ip
netsh interface ip set address "LAN" static 192.168.1.10 255.255.255.0 192.168.1.254
netsh interface ip set dns "LAN" static 192.168.1.100 primary no
netsh interface ip add dns "LAN" 192.168.1.101 index=2 no
ipconfig /flushdns
::设置代理
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Internet Settings" /v ProxyEnable /t REG_DWORD /d 1 /f
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Internet Settings" /v ProxyServer /d "127.0.0.1:8087" /f
ipconfig /flushdns
::禁用代理
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Internet Settings" /v ProxyEnable /t REG_DWORD /d 0 /f
```


# 常用命令

```bat
::根据 ip 找 mac 地址（arp 指令只能查找本网段的 mac 地址）
arp -d *
ping $ip
arp -a
::根据 ip 查主机名（linux 使用 nmblookup -A $ip）
nbtstat -A $ip
ping -a $ip
::直接跳转到 e 盘，即便当前命令行在其他盘符中
cd /d e:\dir
::将目录（包括其下所有子文件和子文件夹）所有者更改为管理员组（administrators）
takeown /a /r /d y /f ~~DIR~~ 
::将目录（包括其下所有子文件和子文件夹）NTFS权限修改为仅管理员组（administrators）完全控制
cacls ~~DIR~~ /t /g administrators:F
::在目录（包括其下所有子文件和子文件夹）NTFS权限上添加管理员组（administrators）完全控制
cacls ~~DIR~~ /t /e /g administrators:F 
::取消管理员组（administrators）完全控制权限
cacls ~~DIR~~ /t /e /r "administrators" 
::查看进程详情
wmic process list >1.txt && notepad.exe 1.txt 
::快速删除文件夹
del /f/s/q ~~DIR~~ > nul 
rd /s/q ~~DIR~~ 
::数学运算
set /A (10-5)*2+6/2 
::添加用户
net user ~~USERNAME~~ ~~PASSWORD~~ /add 
::设置CMD窗口背景色和字体颜色
color 1a 
color 0f
::设置CMD窗口标题
title ~~TITLE~~ 
::设置CMD窗口提示为 `#`，然后再还原回去
prompt # 
prompt $P$G
::CMD清屏
cls
::日期与时间
date
date /T
time
time /T
::start指令
start "My Shell"
start /d "C:\windows\system32"
start /min notepad
start /max mspaint
start /separate calc
start /shared write
start /low winword
start /normal iexplore.exe
start /high explorer.exe
start /realtime ...
start /abovenormal ..
start /belownormal mshearts.exe
start /wait tree
start /b
exit
::列出远程主机任务列表
tasklist /s ~~\\SERVERNAME~~ /u ~~USERNAME~~ /p ~~PASSWORD~~
::列出任务相关DLL信息
tasklist /M
::列出任务相关服务信息
tasklist /SVC
::列出任务相关运行时详细（verbose）
tasklist /V
::过滤任务列表中无响应的任务
tasklist /FI "status eq not responding"
::过滤任务列表中PID小于1000的任务
tasklist /FI "pid lt 1000"
::过滤任务列表中会话数为1的任务
tasklist /FI "session eq 1"
::过滤任务列表中占用内存大于100000k的任务
tasklist /FI "memusage gt 100000"
C:\>tasklist /FI "services ne explorer.exe"
C:\>tasklist /FI "username ne cybercrawler"
C:\>tasklist /FI "services ne themes" /FI "services ne server"
C:\>tasklist /FI "windowtitle eq Untitle*"
C:\>tasklist /FI "modules eq winsta.dll"
C:\>tasklist /S //productionserver /U administrator /P $3cr3t /FI "memusage gt 5000" /FI "windowtitle eq Untitled*"
C:\>taskkill /S 10.199.64.66 /U admin /P adminP4$$ /im soundmix.exe
C:\>taskkill /f /im userinit.exe
C:\>taskkill /f /PID 556
C:\>taskkill /f /im fun.exe /t
C:\>taskkill /S \\10.199.64.66 /U technocrawl /P 123@654 /IM remoteshell.exe /PID 1524 /T /PID 2415 /T /PID /T 995 /t /FI “memusage gt 20000” /T
C:\>label Root Drive
C:\>label D: Softwares
C:\>tree /A
C:\>tree /F
C:\>ver
C:\>type first.bat
@echo off
echo HELLO
pause
C:\>Convert D: /FS:NTFS
C:\>shutdown -s -t 10000 -c "This is a Temporary shutdown for updating the production softwares, and machines will be up soon"
C:\>shutdown -a
C:\>shutdown -l
C:\>shutdown -r
C:\>shutdown -I
C:\>shutdown /m \\10.199.64.71 -r -t 45 -c "This is a temporary reboot, for updating the production software’s, and machines will be up soon."
C:\>shutdown /f -l -t 00
C:\>at \\localhost 10AM "notepad.exe"
新加了一项作业，其作业 ID = 1
C:\>at
状态 ID     日期                    时间          命令行
-------------------------------------------------------------------------------
        1   明天                    10:00         notepad.exe
C:\>at \\localhost 10AM "notepad.exe"
新加了一项作业，其作业 ID = 1
 
C:\>at
状态 ID     日期                    时间          命令行
-------------------------------------------------------------------------------
        1   明天                    10:00         notepad.exe
 
C:\>at \\localhost 1
 
任务 ID:       1
状态:          OK
计划:          明天
时间:   10:00
交互:          No
命令:          notepad.exe
C:\>at 1 /delete
 
C:\>at
列表是空的。
C:\>at 5:11PM /interactive notepad
C:\>at 5:22PM /interactive /every:1,10,15,20,25 servermonitor.exe
C:\>at 5:22PM /interactive /next:M,T,TH,S,SU servermonitor.exe
C:\>at /delete /yes
```




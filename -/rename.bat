@echo off
setlocal enabledelayedexpansion
for /r %%i in (*.*) do (set a=%%~nxi
set n=!a: =_!
ren "%%i" "!n!"
)
pause
cd /d %1
for /r %%d in (.) do if exist "%%d\.svn" rd /s /q "%%d\.svn"

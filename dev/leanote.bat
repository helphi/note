set GOROOT=D:\dev\go1.3.1.windows-amd64\go
set GOPATH=d:%HOMEPATH%\gopath-1.3.1
set Path=%Path%;%GOROOT%\bin;%GOPATH%\bin;D:\dev\PortableGit-2.5.2.2-64-bit\cmd
echo http://localhost:9000/admin/abc123/demo@leanote.com/demo@leanote.com
revel run github.com/leanote/leanote

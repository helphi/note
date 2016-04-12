netsh interface ip set address "LAN" static 172.26.3.63 255.255.255.0 172.26.3.254
netsh interface ip set dns "LAN" static 172.17.18.18 primary no
netsh interface ip add dns "LAN" 172.17.42.201 index=2 no
ipconfig /flushdns
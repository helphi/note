netsh interface ip set address "LAN" static 192.168.10.66 255.255.255.0 192.168.10.254
netsh interface ip set dns "LAN" static 202.100.192.68 primary no
ipconfig /flushdns
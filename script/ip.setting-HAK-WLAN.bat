netsh interface ip set address "WLAN" static 192.168.10.30 255.255.255.0 192.168.10.254
netsh interface ip set dns "WLAN" static 202.100.192.68 primary no
ipconfig /flushdns
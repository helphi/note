netsh interface ip set address "LAN" dhcp
netsh interface ip set dns "LAN" dhcp
netsh interface ip set address "WLAN" dhcp
netsh interface ip set dns "WLAN" dhcp
ipconfig /flushdns

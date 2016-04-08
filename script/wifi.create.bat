echo netsh wlan set hostednetwork mode=disallow
echo netsh wlan show drivers
echo netsh wlan show hostednetwork
netsh wlan set hostednetwork mode=allow ssid=test key=123456789

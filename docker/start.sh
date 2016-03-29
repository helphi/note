http_proxy=http://172.17.18.84:8080 \
https_proxy=http://172.17.18.84:8080  \
no_proxy=hub.com docker  \
--registry-mirror=http://hub.com  \
--registry-mirror=https://6pzhi4th.mirror.aliyuncs.com  \
daemon  \
--bip=172.172.172.1/24  \
--dns=172.26.1.251  \
--insecure-registry=hub.com  \
-D 

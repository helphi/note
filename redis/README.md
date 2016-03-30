# ubuntu 上安装 redis

```sh
$ sudo apt-get install redis-server
$ sudo cp /etc/redis/redis.conf /etc/redis/redis.conf.default
$ sudo redis-server /etc/redis/redis.conf
$ redis-cli
redis> set foo bar
OK
redis> get foo
"bar"
```

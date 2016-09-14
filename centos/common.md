# 常用命令

```bash
scp file.tar.gz root@192.168.1.1:/home #文件网络复制
tar zcvf file.tar.gz file #压缩
tar zxvf file.tar.gz #解压
gzip file #使用 gzip 压缩
chmod a+w dir #让所有用户对文件夹有写入权限
chmod o-w dir #取消其他用户对文件夹的写入权限
chgrp group /opt/app #修改用户组
chown -R user:group /home/user #修改用户所有者，-R 表示包括子文件夹
ls -al /webdocs
ll -atr    #按时间反序
ll -aSr    #按大小反序
du -sh wstemp #查看目录 wstemp 的大小
du -sh * | sort -rn #查看当前目录下所有文件及文件夹大小并排序
find . -name '*.html' -type f -mmin -30  -ls # 查找当前目录下最近30分钟内修改过的.html文件的详细情况
find . -type f -mtime 5  # 查找当前目录下5天前天修改过的常规文件
find . -type f -mtime -5  # 查找当前目录下修改时间小于5天的常规文件
find . -type f -mtime +5 # 查找当前目录下修改时间大于5天的常规文件
find . -size +5000 #查找大于 5M 的文件
find / -name httpd.conf
grep -r --include="*.txt" "key" /opt/app  #抓取关键字时循环子目录且中所有txt文件
rz #上传文件，指令不存在可用 yum install lrzsz 安装
sz file.ext  #下载文件，指令不存在可用 yum install lrzsz 安装
crontab -l  #查看当前用户定时任务
crontab -e #编辑当前用户定时任务
crontab -l -u root    #查看root用户定时任务
ps -ef | grep STRING | grep -v grep | awk '{print $2}' | xargs kill -9 #杀掉指令中包含STRING的进程
kill -9 $( netstat -tlnp | grep PORT | awk '{print $7}' | awk -F '/' '{print $1}' ) #根据端口杀进程
echo  "123456"  |  passwd  --stdin jboss #将 jboss 用户的密码修改为 123456
service network restart #重启网络
/etc/init.d/network restart #重启网络
ifdown eth0 #关闭网络
ifup eth0 #启动网络
ifconfig eth0 down #关闭网络
ifconfig eth0 up #启动网络
man builtin #显示所有内建命令
连续按两次TAB #显示所有命令，包括内建和外部
alias    #显示命令别名
. ./file.sh    #和（source ./file.sh）一样，在当前shell中执行脚本
./file.sh    #另启一个shell执行脚本，需要该脚本有执行权限
!5    #执行第5个历史命令
!ser    #执行历史命令中以ser开头的命令
```

 


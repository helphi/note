# 环境变量的配置文件

- /etc/enviroment：系统环境变量，系统登录时第一个被读取
- /etc/profile：为系统的每个用户设置环境信息，当用户第一次登录时该文件被执行，并从 `/etc/profile.d` 目录的配置文件中搜集 shell 的设置
- /etc/bashrc：为每一个运行 bash shell 的用户执行此文件，当 bash shell 被打开时，该文件被读取
- ~/.bash_profile：每个用户都可使用该文件输入专用于自己使用的 shell 信息，当用户登录时该文件仅仅执行一次。默认情况下，它设置一些环境变量，执行用户的 .bashrc 文件。
- ~/.bashrc：该文件包含专用于你的 bash shell 的 bash 信息，当登录时以及每次打开新的 shell 时，该文件被读取
- ~/.bash_logout： 当每次退出 bash shell 时执行该文件

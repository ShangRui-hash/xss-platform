# XSSPlatform

前端: Vue+ElementUI+axios  
后端: Gin+MySQL+Redis  
# 演示站点：
https://8b1t.cn
# 软件截图
![avatar](https://img-blog.csdnimg.cn/20210828180007409.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)
![avatar](https://img-blog.csdnimg.cn/20210828180127793.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)
支持消息推送(当收到战利品时，服务端会自动向客户端推送消息)
![avatar](https://img-blog.csdnimg.cn/20210828183223575.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)
![avatar](https://img-blog.csdnimg.cn/20210828180222750.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5peg5Zyo5peg5LiN5Zyo,size_20,color_FFFFFF,t_70,g_se,x_16)


# 部署方法
先启动MySQL 和Redis  
要求MySQL 版本高于5.5 版本，因为5.5版本以及以下的版本一张表中不允许定义两个timestamp类型  

然后将back_end目录下的init.sql 文件导入mysql数据库  
```
mysql -u你的用户名 -p你的密码<init.sql
```
将redis_init.json 导入到redis中(需要先安装redis-dump工具)
```
cat redis_init.json | redis-load
```

redis-dump安装方法：
```
apt-get update
apt-get install ruby ruby-dev gcc
gem install redis-dump
```
因为项目使用了google的ReCAPTCHA 验证码，所以需要配置相关的秘钥  
注册账号：https://www.google.com/recaptcha/admin

将以下值导入到环境变量中：
```
#mysql  
export mysql_username=你的mysql用户名
export mysql_password=你的mysql密码  

#Google ReCAPTCHA
export reCAPTCHA_public_key=申请到的google验证码的公钥

#gin 生产环境下设置为release,将会自动加载生产环境下的配置文件 conf/config.production.json
export GIN_MODE=release
```
可以将以上语句添加到$HOME/.bashrc 脚本文件中,然后让shell进程读取该脚本文件即可
```
source $HOME/.bashrc
```
在./front_end 文件夹下创建 .env.local文件，配置google 验证码 浏览器端的秘钥
```
# 环境变量文件:在所有的环境中被载入，但会被 git 忽略
VUE_APP_RECAPTCHA_SITEKEY="你的秘钥"
```
修改./back_end/conf/config.json 
```
    "baseurl":"你申请的域名",
```
项目中默认使用了 https的测试证书。如果你在本地调试，不能使用https，所以需要关闭配置文件中的use_https选项。
如果你需要在服务器上使用https，可以将你的证书和秘钥文件的路径写入环境变量，程序自动会读取。如果没有写入环境变量，程序默认使用SSL文件下的https测试证书
```
vim $HOME/.bashrc
```
添加以下内容： 
```
#HTTPS
export key_file=./SSL/8b1t.cn.key #替换为你的key文件的路径
export cert_file=./SSL/8b1t.cn.pem #替换为你的pem文件的路径
```
让shell进程读取 $HOME/.bashrc
```
source $HOME/.bashrc
```

# TODO
1. 修改密码功能
2. 尝试实现实时远控
# PRODUCE 

### gorm 参考使用

https://segmentfault.com/a/1190000013216540

启动mysql

`docker run --name mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 mysql:5.7`

如果Sequel Pro 连接报错`Authentication plugin 'caching_sha2_password' cannot be loaded: dlopen(/usr/local/lib/plugin/caching_sha2_password.so, 2): image not found`, 则

`docker exec -it mysql bash`

`$mysql -uroot -p`

`$ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';`

### goland 无法解析go.mod 的包

开始建立一个go 项目， 使用go mod。一直都是OK的。然而在另一个机子导入、或者本机挪了个位置，就出现goland 无法识别第三方包的情况。

当前goland 第三方包或者go mod init xxx，使用 xxx/config 之类的也都全部无法识别。

找了半天看到[帖子](https://learnku.com/go/t/35849)中跟我描述一致，最后一个回复中，给出另一个链接[segmentfault](https://segmentfault.com/q/1010000020603338/)
提到、而且截图和我情况一致。

下面回复点醒了我:

> 如果是导入已有项目，记得启用下 Go Module，进入 Preference -> Go -> Go Modules，启用下就好。

也就是说，导入、或者挪了位置，要启用。

可是当前我的已经启用了呀Go - Go Modules(vgo) - Enable Go Modules (vgo)integration

这个checkbox 本来就是加载的。

此时我关闭他，然后apply; 再然后checkbox 打开他，再apply，这时就全部正常了。

### MYSQL

    docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
    
    docker exec -it mysql mysql -uroot -p
    
    create database novel charset utf8;
    
centos - /etc/my.cnf
将datadir=/var/lib/mysql 注释掉

[mysqld]
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock
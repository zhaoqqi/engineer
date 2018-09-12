### 安装shadowsocks客户端

#### Debian/Ubuntu
```bash
apt-get install python-pip
pip install shadowsocks
```

#### Centos
```bash
yum install python-setuptools && easy_install pip
pip install shadowsocks
```

#### 配置
```bash
vim /etc/shadowsocks.json

{
  "server":"my_server_ip",
  "local_address": "127.0.0.1",
  "local_port":1080,
  "server_port":my_server_port,
  "password":"my_password",
  "timeout":300,
  "method":"aes-256-cfb"
}
```   

- server: shadowsocks服务器IP
- my_server_port: shadowsocks服务器端口
- password: 改为自己的服务器密码
- method: 加密算法 

#### 启动
```bash
sslocal -c /etc/shadowsocks.json -d start //启动
sslocal -c /etc/shadowsocks.json -d stop //停止
```  

#### 开机自启动
将启动命令加入到用户.bashrc中   

#### 设置浏览器代理
以firefox浏览器为例：   
首选项 -> 高级 -> 网络 -> 连接 -> 设置   
手动配置代理：   
SOCKS主机：127.0.0.1   端口：1080   
不使用代理：localhost, 127.0.0.1   

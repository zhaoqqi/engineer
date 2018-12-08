## 打造 ubuntu 16.04 工作环境

#### Ubuntu16.04 的优点
- 屏幕亮度调节，每一级都有亮度变化，比14.04灵敏，好评；
- 系统设置 suspend 后，再次唤醒速度很快，好用；
- Ubuntu 自带的分屏功能，默认支持4个分屏的切换，好用；

#### shadowsocks 设置
- ss 客户端设置
[ubuntu下设置ss客户端](https://github.com/zhaoqqi/engineer/blob/master/tools/shadowsocks.md) 

通过系统网络设置实现：
进入代理设置 System settings > Network > Network Proxy
设置Method为Automatic
设置Configuration URL为autoproxy.pac文件的路径
```bash
$ apt-get install python-pip
$ pip install genpac
$ genpac -p "SOCKS5 127.0.0.1:1080" --gfwlist-proxy="SOCKS5 127.0.0.1:1080" --gfwlist-url=https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt --output="autoproxy.pac"
$ mv autoproxy.pac /home/zhaoqi/Documents/autoproxy.pac
```
也可以设置为手动模式：
```bash
进入代理设置 System settings > Network > Network Proxy
设置Method为Manual
Socks Host: 127.0.0.1, Port: 1080
```
[Linux下使用Shadowsocks和PAC玩转小飞机](http://lckiss.com/?p=2172) 

- firefox 设置
Preformance -> Network Setting -> 选中 Manual proxy configuration；
选择 SOCKS V5协议，SOCKS Host 127.0.0.1, port 1080； 
在 No Proxy for 中可以设置无需代理的访问地址，比如 *.cn、公司内网地址等等；

#### 连接 vpn
使用公司提供的 Linux VPN客户端即可。

#### thunderbird 邮件客户端
Ubuntu 自带的邮件客户端，简单设置服务器地址和端口、选择邮件协议后即可使用。
另外，可以通过过滤器分类邮件，跟 windows 下 foxmail 一样简单好用。

#### 浏览器
- firefox
目前默认使用的浏览器，网络代理设置方便。
- chrome
无法在浏览器中设置 SS 代理，需要使用命令行启动参数配置，这方面没有 firefox 方便。

#### goland
官网下载 Linux 版本安装包，解压后进入 bin 目录，直接执行 goland.sh 即可启动运行。

#### Markdown 编辑器
- Remarkable
- Typora

#### 社交软的
- 使用网页版微信
- 使用手机QQ……
- 不过，像我们公司使用 one piece 做内部通信软件，有 Linux 客户端，还是很幸福的，HOHO～

#### 主题设置
paper主题的安装
```bash
sudo add-apt-repository ppa:dyatlov-igor/materia-theme
sudo apt update
sudo apt install materia-gtk-theme
在软件管理中心安装 unity weak tools
安装成功后，打开unity weak tools进行相应设置即可
```

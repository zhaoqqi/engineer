## 打造 ubuntu 16.04 工作环境

#### Ubuntu16.04 的优点
- 屏幕亮度调节，每一级都有亮度变化，比14.04灵敏，好评；
- 系统设置 suspend 后，再次唤醒速度很快，好用；
- Ubuntu 自带的分屏功能，默认支持4个分屏的切换，好用；

#### shadowsocks 设置
- ss 客户端设置
[ubuntu下设置ss客户端](https://github.com/zhaoqqi/engineer/blob/master/tools/shadowsocks.md) 

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

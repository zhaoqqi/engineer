#### 5种网络模式

- bridge
- host
- container
- none

#### bridge 模式

bridge 模式也是 Docker 默认使用的网络模式。

Docker 在 docker daemon 运行的宿主机上创建一个网桥 docker0，并为该网桥分配 IP 地址，比如 172.17.0.1/24。

Docker 为宿主机上运行的每个容器创建一个网络设备 veth pair，一端作为容器的网卡 eth0，一端连接到网桥 docker0 之上，由网桥 docker0 为容器网卡 eth0 分配 IP 地址，比如 172.17.0.1 。如下图所示：

![](E:\git\engineer\opensource-project\docker\images\docker-network-bridge.png)

#### host 模式

在 host 模式下，容器共享宿主机的网络，容器没有自己的 IP 地址，只能通过宿主机的 IP 地址+端口映射的方式对外提供访问。好处是没有主机网桥会带来网络性能的提升。

![](E:\git\engineer\opensource-project\docker\images\docker-network-host.png)

#### container 模式

在 container 模式下，容器共享同一宿主机中另一个容器的网络命名空间。

使用如下命令启动容器：

```bash
# test_container 与 another_container 拥有相同的网络命名空间，拥有相同的IP地址
$ docker run --net=container:another_container ubuntu:14.04 test_container
```

#### none 模式

该模式将容器放置在它自己的网络命名空间中，但不进行任何网络配置。实际上，该模式关闭了容器的网络功能。

应用场景：

1. 容器不需要网络，例如执行只需要磁盘读写的任务；
2. 用户希望自定义网络；

```bash
$ docker run -d -P --net=none nginx:1.9.1 
$ docker ps
CONTAINER ID  IMAGE          COMMAND   CREATED
STATUS        PORTS          NAMES
d8c26d68037c  nginx:1.9.1    nginx -g  2 minutes ago
Up 2 minutes                 grave_perlman
$  docker inspect d8c26d68037c | grep IPAddress
"IPAddress": "",
"SecondaryIPAddresses": null,
```


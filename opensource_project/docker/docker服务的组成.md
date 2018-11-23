

#### OCI

Open Container Initiative，开放容器标准组织。

由开放容器标准组织（OCI）制定的容器运行时和镜像格式标准，称为OCI标准。

#### docker

docker 的发展可以大致分为两个阶段

第一阶段，docker 主要功能为容器的生命周期管理和镜像管理，仅是一个容器管理平台；

第二阶段，docker 增加了对容器存储、容器网络等的管理，变身成为了整个容器生态的管理平台。

![](E:\git\engineer\opensource-project\docker\images\arch.png)

#### docker client

可由用户直接执行的命令行工具，docker client 通过 HTTP 的方式与 docker daemon 服务交互，从而实现单机环境下的容器和镜像的管理操作。

#### docker engine

即 dockerd，docker daemon 进程。其封装了直接面向用户的容器管理操作命令，通过 HTTP server 的方式对外暴露服务。

#### containerd

为 docker 管理平台提供统一的抽象层，该抽象层的作用是处理系统调用、跨平台细节屏蔽。

containerd 是容器技术标准化的产物，为了能够兼容 OCI 标准，从 docker daemon 中将容器运行时及其管理功能剥离出来的。

containerd 通过 gRPC 向上为 docker daemon 提供接口。向下通过 containerd-shim 结合 runC 。

![](E:\git\engineer\opensource-project\docker\images\containerd.png)

#### docker-shim/containerd-shim

真实 docker 容器的运行载体，每启动一个容器都会启动一个 docker-shim/containerd-shim 作为容器运行的载体。

#### runC

OCI 定义了容器运行时标准，runC 就是 docker 依据开放容器格式标准（OCF, Open Container Format）开发的一种具体实现。

runC 是一段二进制脚本代码，与操作系统的 cgroup 和 namespace 交互，只负责启动/关闭容器、资源隔离/限制等功能。


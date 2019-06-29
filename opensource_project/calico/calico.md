
## Calico 学习
### 一、背景&目标
官网给出的Calico出现的目标为：Free and open source, Project Calico is designed to simplify, scale, and secure cloud networks 。简单翻译为：自由开源的Calico设计目标为为云端服务提供简化的、可扩展的、安全的网络服务。

### 二、优势&劣势？有哪些trade-off？
#### 优势
1. 简单
2. 可扩展
3. 安全
4. 完全基于路由，性能好
5. 具备网络隔离机制
6. 不封装网络报文，可以使用标准调试工具跟踪问题
7. 可以与Istio集成，用户可以配置强大的规则以描述pod如何发送和接受流量，提高安全性并控制网络

#### 劣势
1. 只提供bgp协议路由的方式，适用场景会有局限（比如人的偏好导致不会使用？）
2. Network policy 使用 iptables，网络性能会虽着规则条目增加线性下降

#### 设计上的trade-off，要什么不要什么，为什么？
使用路由和bgp协议打通节点间的网络

### 三、适用场景
1. 对集群网络性能有要求的
2. 对网络隔离功能有要求的

### 四、技术的组成和关键点
#### Felix
Calico agent，运行于每个 node 节点（网络节点），负责节点的路由管理。
#### iBgp client
internal BGP client，BGP协议客户端，iBGP负责同一子网内部的路由学习。
#### eBgp client
edge BGP client，BGP协议客户端，eBGP负责子网之间的路由学习。
#### Route reflector
负责全网路由学习，与每个iBGP连接，可以减少iBGP/eBGP之间的两两连接，减少路由学习路径复杂度。适合大规模集群使用。

### 五、技术的底层原理和关键点
#### IP路由协议
#### BGP协议

### 六、已有的实现和它的对比
1. flannel
简单
2. canal
集成了Flannel和Calico两者的能力
3. weave
可以对整个网络进行简单加密，但会增加不小的网络开销。
4. Calico


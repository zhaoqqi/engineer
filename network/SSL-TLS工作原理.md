### SSL/TLS工作原理——单向和双向认证原理

#### SSL/TLS介绍
SSL是 secure sockets layer，TLS是 transport layer security。TLS 是 SSL 的继任者。
作用就是在TCP层以上增加一层加密，以保证传输数据的安全。配置了 SSL/TLS 的服务，就被称为 HTTPS 服务。   

#### 成员

##### CA
证书授权中心
- CA本身是受信任的，各大浏览器默认信任的，国际认可的。
- CA给它信任的申请对象颁发证书，CA要确认申请者的合法身份，同时要收取费用，而且CA可以吊销申请者的证书。
- CA拥有 ca.crt 和 ca.key，分别是 CA 机构的根证书和私钥。

##### SSL Server
配置了 SSL/TLS 的服务端。   
server 生成服务端的公私钥 server.pub 和 server.key，私钥用来加密，私钥用来揭秘。   
server 使用 server.pub 生成请求文件 server.req，请求文件中包含 server 的信息，如域名、申请者、公钥等。   
server 将请求文件递交给 CA，CA 验证后，使用 ca.key 和请求文件 server.req 加密生成签名证书 server.crt。   

##### SSL Client
如果需要双向 SSL/TLS 认证，client 也需要上一部分 SSL Server 的操作：   
client 生成服务端的公私钥 client.pub 和 client.key，私钥用来加密，私钥用来揭秘。   
client 使用 client.pub 生成请求文件 client.req，client 请求文件中包含哪些内容呢？   
client 将请求文件递交给 CA，CA 验证后，使用 ca.key 和请求文件 client.req 加密生成签名证书 client.crt。  

#### 单向认证流程
单向认证指的是只有一个对象校验对端的证书合法性。通常都是 client 端来校验 server 端的合法性。   
client 端需要 ca.crt，server 端需要 server.crt 和 server.key。   
![Alt text](../opensource_project/kubernetes/https.png "HTTPS单向认证流程")
<i class="HTTPS单向认证流程"></i>   

1. SSL Client 向 SSL Server 发起 HTTPS 请求，同时要求验证 SSl Server的证书合法性；
2. SSL Server 将 server.crt 签名证书发送给 SSL Client；
3. SSL Client 使用 ca.crt 验证 server.crt 的合法性。验证通过则下一步，不通过则弹框提示风险；
4. SSL Client 生成 random key，并使用 server.crt 加密，将加密后的 random key 发送给 SSL Server；
5. SSL Server 使用 server.key 解密 random key，并使用双方协商的对称加密算法+random key，将 SSL Client 在第1步请求的处理结果返回；
6. SSL Client 使用 random key 解密 response content，最终获取到请求结果。

#### 双向认证流程
双向认证指的是互相认证，server 需要校验每个 client，client 也需要校验 server。   
server 需要 ca.crt，server.crt，server.pub   
client 需要 ca.crt，client.crt，client.pub   
双向认证对比单向认证，多了一步 server 验证 client 证书有效性的步骤。   


#### 认证证书格式
.crt 表示证书   
.key 表示私钥   
.req 表示请求文件   
.csr 也表示请求文件   
.pem 表示 pem 格式   
.der 表示 der 格式   
文件扩展名可以随意命名的，只是为了便于理解。   
所有证书、私钥都可以是 pem 和 der 格式的，取决于需求。   
pem 和 der 格式可以互转：
```bash
openssl x509 -in ca.crt -outform DER -out ca.der  //pem -> der
openssl x509 -inform der -in ca.der -out ca.pem    // der -> pem
```
pem 格式经过加密的文本文件，一般有以下几个开头结尾格式：
```bash
-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----
or:
-----BEGIN CERTIFICATE REQUEST-----
-----END CERTIFICATE REQUEST-----
or:
----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----
```
der格式: 经过加密的二进制文件。   

#### SSL/TLS 与 Openssl/mbedtls 的关系
SSL/TLS 是一种工作原理，Openssl/mbedtls 是该工作原理的实现，类似于 TCP/IP 协议与 Socket 之间的关系。   

#### 本地生成SSL（充当CA）
可以在本地使用 OpenSSL 工具建立一个 CA（ca.crt+ca.key），用这个 CA 给 server 和 client 颁发证书。   


#### 参考
[SSL/TLS 双向认证(一) -- SSL/TLS工作原理](https://blog.csdn.net/ustccw/article/details/76691248, "SSL/TLS 双向认证(一) -- SSL/TLS工作原理")   
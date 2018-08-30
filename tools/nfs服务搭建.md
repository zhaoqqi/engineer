#### server端
1. 安装 nfs-utils
```bash
yum install nfs-utils
```
2. 配置文件/etc/exports中增加client端的访问权限
```bash
/data    x.x.x.x1(rw,sync,no_root_squash,no_subtree_check)
/data    x.x.x.x2(rw,sync,no_root_squash,no_subtree_check)
```

#### client端nfs挂载
1. 安装 nfs-utils
```bash
yum install nfs-utils
```
2. 挂载命令
```bash
mount -o rw,nosuid,fg,hard,intr $server_ip:/dir /data
```
3. 设置开机自动挂载
```bash
文件 /etc/fstab
server_ip:/dir /data nfs defaults 0 0
```

## Linux cgroup

Linux CGroup 全称 Linux Control Group，是  Linux 内核的的一个功能，用来限制、控制、分离一个进程组群的资源（比如CPU、内存、磁盘IO等）。于2007年合并到2.6.24版本的内核中。

主要提供了如下功能：

- Resource limitation: 限制资源使用，比如内存使用的上限以及文件系统的缓存限制。
- Prioritization: 优先级控制，比如：CPU利用和磁盘IO吞吐。
- Accounting: 审计或统计，主要目的是为了计费。
- Control: 挂起/恢复执行进程。

在实践中，系统管理员一般会利用 CGroup 做下面这些事：

- 隔离一个进程集合（比如：Nginx 的所有进程），并限制它们所消费的资源，比如绑定 CPU 的核
- 为这组进程分配足够其使用的内存
- 为这组进程分配相应的网络带宽和磁盘存储限制
- 限制这组进程访问某些设备（通过设置设备的白名单）



来点感性认识，看看 CGroup 是怎么做到的。

首先，Linux 是使用 file system 来实现的 CGroup。在 Ubuntu 14.04 下，输入以下命令就可以看到 CGourp 已经为你 mount 好了。

```bash
zq@ubuntu:~$ mount -t cgroup
cgroup on /sys/fs/cgroup/cpuset type cgroup (rw,relatime,cpuset)
cgroup on /sys/fs/cgroup/cpu type cgroup (rw,relatime,cpu)
cgroup on /sys/fs/cgroup/cpuacct type cgroup (rw,relatime,cpuacct)
cgroup on /sys/fs/cgroup/memory type cgroup (rw,relatime,memory)
cgroup on /sys/fs/cgroup/devices type cgroup (rw,relatime,devices)
cgroup on /sys/fs/cgroup/freezer type cgroup (rw,relatime,freezer)
cgroup on /sys/fs/cgroup/blkio type cgroup (rw,relatime,blkio)
cgroup on /sys/fs/cgroup/net_prio type cgroup (rw,net_prio)
cgroup on /sys/fs/cgroup/net_cls type cgroup (rw,net_cls)
cgroup on /sys/fs/cgroup/perf_event type cgroup (rw,relatime,perf_event)
cgroup on /sys/fs/cgroup/hugetlb type cgroup (rw,relatime,hugetlb)
```

或者使用 lssubsys 命令：

```bash
$ lssubsys  -m
cpuset /sys/fs/cgroup/cpuset
cpu /sys/fs/cgroup/cpu
cpuacct /sys/fs/cgroup/cpuacct
memory /sys/fs/cgroup/memory
devices /sys/fs/cgroup/devices
freezer /sys/fs/cgroup/freezer
blkio /sys/fs/cgroup/blkio
net_cls /sys/fs/cgroup/net_cls
net_prio /sys/fs/cgroup/net_prio
perf_event /sys/fs/cgroup/perf_event
hugetlb /sys/fs/cgroup/hugetlb
```

在 /sys/fs 下有一个 cgroup  目录，该目录下还有很多子目录，比如：cpu，cpuset，memory，blkio……，这些目录都是 cgroup 的子系统，分别用于限制不同资源。

如果你没有看到上述目录，那么也可以自己动手 mount，下面给了一个例子：

```bash
mkdir cgroup
mount -t tmpfs cgroup_root ./cgroup
mkdir cgroup/cpuset
mount -t cgroup -ocpuset cpuset ./cgroup/cpuset/
mkdir cgroup/cpu
mount -t cgroup -ocpu cpu ./cgroup/cpu/
mkdir cgroup/memory
mount -t cgroup -omemory memory ./cgroup/memory/
```

你可以到其中一个子系统中，比如cpu，创建一个子目录。创建后你会发现，一旦子目录被创建，子目录下面又有很多子目录生产。

```bash
zq@ubuntu:/sys/fs/cgroup/cpu$ sudo mkdir zq
[sudo] password for zq:
hchen@ubuntu:/sys/fs/cgroup/cpu$ ls ./zq
cgroup.clone_children  cgroup.procs       cpu.cfs_quota_us  cpu.stat           tasks
cgroup.event_control   cpu.cfs_period_us  cpu.shares        notify_on_release
```

好了，让我们来看几个例子。



#### CPU 限制

假设我们有个非常吃 CPU 的程序，叫 deadloop，其源码如下：

```c
int main(void)
{
    int i = 0;
    for(;;) i++;
    return 0;
}
```

运行起来后，毫无疑问，CPU 被干到100%，下面是 top 的输出：

```bash
PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
3529 root      20   0    4196    736    656 R 99.6  0.1   0:23.13 deadloop
```

然后，我们在 /sys/fs/cgroup/cpu 下创建一个叫 zqtest 的 group。我们来设置一下这个 group 的 cpu 利用率：

```bash
zq@ubuntu:~# cat /sys/fs/cgroup/cpu/zqtest/cpu.cfs_quota_us
-1
root@ubuntu:~# echo 20000 > /sys/fs/cgroup/cpu/zqtest/cpu.cfs_quota_us
```

我们看到，这个进程的PID是3529，我们把这个进程加到这个cgroup中：

```bash
# echo 3529 >> /sys/fs/cgroup/cpu/haoel/tasks
```

然后，就会在top中看到CPU的利用立马下降成20%了。（前面我们设置的20000就是20%的意思）

```bash
PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
3529 root      20   0    4196    736    656 R 19.9  0.1   8:06.11 deadloop
```

#### 内存使用限制

我们再来看一个限制内存的例子（下面的代码是个死循环，其它不断的分配内存，每次512个字节，每次休息一秒）：

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>
int main(void)
{
    int size = 0;
    int chunk_size = 512;
    void *p = NULL;
    while(1) {
        if ((p = malloc(p, chunk_size)) == NULL) {
            printf("out of memory!!\n");
            break;
        }
        memset(p, 1, chunk_size);
        size += chunk_size;
        printf("[%d] - memory is allocated [%8d] bytes \n", getpid(), size);
        sleep(1);
    }
    return 0;
}
```

然后，在我们另外一边：

```bash
# 创建memory cgroup
$ mkdir /sys/fs/cgroup/memory/haoel
$ echo 64k > /sys/fs/cgroup/memory/haoel/memory.limit_in_bytes
# 把上面的进程的pid加入这个cgroup
$ echo [pid] > /sys/fs/cgroup/memory/haoel/tasks
```

你会看到，一会上面的进程就会因为内存问题被kill掉了。

#### 磁盘I/O限制

我们先看一下我们的硬盘IO，我们的模拟命令如下：（从/dev/sda1上读入数据，输出到/dev/null上）

```bash
sudo dd if=/dev/sda1 of=/dev/null
```

我们通过iotop命令我们可以看到相关的IO速度是55MB/s（虚拟机内）：

```bash
TID  PRIO  USER     DISK READ  DISK WRITE  SWAPIN     IO>    COMMAND
8128 be/4 root       55.74 M/s    0.00 B/s  0.00 % 85.65 % dd if=/de~=/dev/null...
```

然后，我们先创建一个blkio（块设备IO）的cgroup

```bash
mkdir /sys/fs/cgroup/blkio/zqtest
```

并把读IO限制到1MB/s，并把前面那个dd命令的pid放进去（注：8:0 是设备号，你可以通过ls -l /dev/sda1获得）：

```bash
root@ubuntu:~# echo '8:0 1048576'  > /sys/fs/cgroup/blkio/zqtest/blkio.throttle.read_bps_device
root@ubuntu:~# echo 8128 > /sys/fs/cgroup/blkio/zqtest/tasks
```

再用iotop命令，你马上就能看到读速度被限制到了1MB/s左右。

```bash
TID  PRIO  USER     DISK READ  DISK WRITE  SWAPIN     IO>    COMMAND
8128 be/4 root      973.20 K/s    0.00 B/s  0.00 % 94.41 % dd if=/de~=/dev/null...
```

#### CGroup 子系统

- blkio — 这个子系统为块设备设定输入/输出限制，比如物理设备（磁盘，固态硬盘，USB 等等）。
- cpu — 这个子系统使用调度程序提供对 CPU 的 cgroup 任务访问。
- cpuacct — 这个子系统自动生成 cgroup 中任务所使用的 CPU 报告。
- cpuset — 这个子系统为 cgroup 中的任务分配独立 CPU（在多核系统）和内存节点。
- devices — 这个子系统可允许或者拒绝 cgroup 中的任务访问设备。
- freezer — 这个子系统挂起或者恢复 cgroup 中的任务。
- memory — 这个子系统设定 cgroup 中任务使用的内存限制，并自动生成内存资源使用报告。
- net_cls — 这个子系统使用等级识别符（classid）标记网络数据包，可允许 Linux 流量控制程序（tc）识别从具体 cgroup 中生成的数据包。
- net_prio — 这个子系统用来设计网络流量的优先级。
- hugetlb — 这个子系统主要针对于HugeTLB系统进行限制，这是一个大页文件系统。

关于各个子系统的参数细节，以及更多的Linux CGroup的文档，你可以看看下面的文档：

[Redhat的官方文档]: https://access.redhat.com/documentation/zh-cn/red_hat_enterprise_linux/6/html-single/resource_management_guide/index#ch-Subsystems_and_Tunable_Parameters


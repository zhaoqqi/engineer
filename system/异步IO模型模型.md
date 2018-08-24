### 异步I/O模型

#### 《Unix网络编程》中的5中I/O模型
- 阻塞I/O
- 非阻塞I/O
- I/O的多路复用（select/poll）
- 信号驱动的I/O（SIGIO）
- 异步I/O（POSIX的aio_functions）

##### 阻塞I/O模型
用户态应用调用系统调用recvFrom从目标设备读取数据，如果此刻目标文件描述符未就绪，则阻塞于该系统调用。待对应设备准备好数据时，内核将数据从内核拷贝到用户空间，recvFrom系统调用返回成功。用户态应用即可从用户空间获取数据并处理。该模型在系统有多个I/O请求时，后续请求只能等待当前请求的I/O返回后才能执行。    
##### 非阻塞I/O模型
用户态应用调用系统调用recvFrom从目标设备读取数据，如果此刻目标文件描述符未就绪，则该系统调用返回EWOULDBLOCK异常信息。用户态应用不断轮询recvFrom系统调用，直到设备数据准备好，recvFrom将数据从内核拷贝到用户空间，返回成功。不停的轮询会浪费大量的CPU处理时间。    
##### I/O复用模型
使用系统调用select/poll，用户态应用阻塞于select/poll系统调用，而不是recvFrom。select/poll系统调用可以同时等待多个文件描述符就绪（select有数量上限，poll没有）。用户态应用通过内核返回的多个文件描述符表，来检查其中已就绪的并处理数据。该模型的缺点是，文件描述符列表每次都需要全量从内核拷贝到用户空间。并且，用户应用需要全量遍历该表后，才可以确认哪些描述符已就绪。select/poll均需要重复初始化文件描述符列表。
##### 信号驱动I/O模型
用户态应用调用recvFrom系统调用后立即返回，可以处理其他计算任务。待目标设备数据就绪后，内核发送SIGIO信号来通知用户态应用。用户态应用可以在信号处理函数中来处理已就绪的数据。
##### 异步I/O模型
告知内核启动某个操作，并在整个操作完成后通知我们。跟信号驱动I/O模型的区别是：信号驱动I/O是由内核通知我们何时可以启动一个I/O操作，异步I/O是由内核通知我们I/O操作何时完成。信号驱动I/O是，内核发出信号通知用户态应用此时可以开始调用系统调用recvFrom从目标设备读/写数据了。异步I/O是，内核通知时，用户态应用即可直接处理已就绪在用户空间的数据了。    
##### libevent
##### libuv
libevent :名气最大，应用最广泛，历史悠久的跨平台事件库；   
libev :较libevent而言，设计更简练，性能更好，但对Windows支持不够好；   
libuv :开发node的过程中需要一个跨平台的事件库，他们首选了libev，但又要支持Windows，故重新封装了一套，linux下用libev实现，Windows下用IOCP实现；   

##### c10k问题
一个网络服务器，同时有10k个client端保持与服务端的连接，并且不时会发送HTTP请求与服务端交互。这种场景，就是c10k的典型场景。   
与TPS/QPS达到10k不同，并发请求10k要求服务端可以并发处理请求的数量达到10k数量级，并且要求很快响应。    
c10k的问题是，同时有10k个client端连接。对于服务器来说，极端情况下同时会有10k个I/O请求到达。对于服务器所在的操作系统来说，频繁的进程/线程切换会产生高昂的代价。导致系统响应越来越慢，以致于系统彻底崩溃。    

##### c10k解决方案
1. 每个连接的处理分配一个进行/线程    
当连接数过大时，系统开销也随着线型增长。   
2. 一个进行/线程处理多个链接    
select/poll/epoll(异步I/O)    
3. libevent
4. libuv(python asyncio/async await + uvloop)


[聊聊C10K问题及解决方案](https://cloud.tencent.com/developer/article/1031629, "聊聊C10K问题及解决方案")    
[magicstack uvloop: Blazing fast Python networking](https://magic.io/blog/uvloop-blazing-fast-python-networking/, "magicstack uvloop: Blazing fast Python networking")    
[uvloop: Python极速网络互连](http://codingpy.com/article/uvloop-blazing-fast-networking-with-python, "uvloop: Python极速网络互连")   
[libuv design overview](http://docs.libuv.org/en/v1.x/design.html, "libuv design overview")   

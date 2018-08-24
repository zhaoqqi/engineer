### lock-free编程

#### 什么是lock-free编程
我个人的理解是，在编写多线程程序时，如果一个线程在获得资源锁以后，因为优先级低或者进行IO动作，而被当前cpu核心切换走。这时，如果其他等待此资源锁的线程均会因此而阻塞等待，则此多线程程序不是lock-free的。   
反过来讲，如果在多线程程序中，线程间不会因为获得资源锁的线程被cpu核心切换为非运行状态而block，则该程序就是lock-free的。   

![Alt text](is-lock-free.png "is lock-free")
<i class="lock-free"></i>   
按照上图所示，简单来说就是：多线程+线程访问共享内存+线程间不会互相block -> is lock-free   

#### 相关技术
要满足lock-free编程中的非阻塞条件，有一系列技术和方法可用。如原子操作（Atomic Operations）、内存栅栏（Memory Barrier）、避免ABA问题（Avoiding ABA Problem）。   

##### 读改写原子操作（Atomic Read-Modify_Write Operations）
原子操作（Atomic Operations）在操作内存时可以被看做是不可分割的，其他线程不会打断该操作，不存在部分完成的说法。在现代的cpu处理器上，很多操作已经被设计为原子的，例如对齐读（Aligned Read）和对齐写（Aligned Write）。   
Read-Modify-Write（RMW）操作的设计则让执行更复杂的事物操作变成了原子的，使得当有多个写入者对相同的内存进行修改时，保证一次只执行一个操作。   
RMW操作在不同的cpu家族中是通过不同的方法来实现的：   
- x86/64和Itanium架构通过Compare-And-Swap（CAS）方式来实现的；
- PowerPC、MIPS和ARM架构通过 Load-Link/Store-Conditionl（LL/SC）方式来实现；   

##### Compare-And-Swap循环（CAS Loops）
如何通过 CAS Loops 来完成对事物的原子处理呢？   
通常，开发人员会设计在一个循环中重复地执行 CAS 操作以试图完成一个事物操作。这个过程分为3步：   
1. 从指定的位置读取原始的值；
2. 根据读取到的值计算出新的值；
3. 检测如果指定内存位置还是原始的值，则使用新值写入该内存位置；   

##### ABA问题
1. 线程A从指定内存位置读取原始的值X；
2. 线程B更新该内存位置的值为Y，并完成其他计算或IO任务；
3. 线程B更新该内存位置的值为X；
4. 线程A使用原始值X计算得到新的值Z；
5. 线程A检测该内存位置的值为X，使用值Z写入该内存； 

##### 内存栅栏


参考资料：[Lock-Free编程](https://www.cnblogs.com/gaochundong/p/lock_free_programming.html#atomic_read_modify_write_operations, "Lock-Free编程")
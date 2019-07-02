Hashmap是一个用于存储Key-Value键值对的结合，每一个键值对也叫做一个Entry。这些个键值对（Entry）分散存储到一个数组当中，这个数组就是HashMap的主干。

HashMap数组每一个元素的初始值都是NULL。

| 0    | 1    | 2    | 3    | 4    | 5    | 6    | 7    |
| ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| NULL | NULL | NULL | NULL | NULL | NULL | NULL | NULL |

对于HashMap，最常用的是两个方法：**Get** 和 **Put**。

#### Put方法的原理

调用Put方法的时候发生了什么呢？

比如调用hashMap.put("apple", 0)，插入一个Key为"apple"的元素。这时候我们需要利用一个哈希函数来确定Entry的插入位置（index）：

index = Hash("apple")

最后假定计算出的index是2，那么结果如下：

| 0    | 1    | 2     | 3    | 4    | 5    | 6    | 7    |
| ---- | ---- | ----- | ---- | ---- | ---- | ---- | ---- |
| NULL | NULL | Entry | NULL | NULL | NULL | NULL | NULL |

但是，因为HashMap长度是有限的，当插入的Entry越来越多时，再完美的哈希函数也难免会出现index冲突的情况。例如下图，如果Entry6的index为2，则会产生冲突：

| 0      | 1    | 2      | 3      | 4    | 5      | 6      | 7    |
| ------ | ---- | ------ | ------ | ---- | ------ | ------ | ---- |
| Entry3 | NULL | Entry1 | Entry5 | NULL | Entry2 | Entry4 | NULL |

这时候我们该怎么办呢？我们可以利用**链表**来解决。

HashMap数组的每一个元素不只是一个Entry对象，也是一个链表的头节点。每一个Entry对象通过Next指针只想它下一个Entry节点。当新来的Entry映射到冲突的数组位置时，只需要插入到对应的链表即可：

| 0      | 1    | 2                | 3      | 4    | 5      | 6      | 7    |
| ------ | ---- | ---------------- | ------ | ---- | ------ | ------ | ---- |
| Entry3 | NULL | Entry6 -> Entry1 | Entry5 | NULL | Entry2 | Entry4 | NULL |

需要注意的是，新来的Entry节点插入链表时，使用的时**"头插法"**。

#### Get方法的原理

使用Get方法根据Key来查找Value的时候，发生了什么呢？

首先会把输入的Key做一次Hash映射，得到对应的index：

index = Hash("apple")

由于刚才所说的Hash冲突，同一个位置有可能匹配到多个Entry，这时候就要顺着对应链表的头节点，一个一个向下来查找。假设我们要找的Key时"apple"，即Entry1：

1. 计算得到index=2，找到链表头节点Entry6，但Entry6的Key是"banana"，顺序往下找。
2. 找到Key为"apple"的Entry2，结束。

之所以把Entry6放在头节点，是因为HashMap的发明者认为，**后插入的Entry被查找的可能性更大。**

#### 问题

1. HashMap默认的初始长度是多少？为什么这么规定？

   HashMap的默认初始长度是16，并且每次自动扩展或手动初始化时，长度必须时2的幂。

   之所以选择16，是为了服务与Key映射到index的Hash算法。如何才能实现一个尽量均匀分布的Hash函数呢？我们通过利用Key的HashCode值来做某种运行。为了实现高效的Hash算法，HashMap的发明者采用了位运算的方式。

   如何进行位运算呢？有如下的公式（Length是HashMap的长度）

   index = HashCode(Key) & (Length-1)

   下面我们以值为"book"的Key来演示整个过程：

   - 计算book的hashcode，结果为十进制的3029737，二进制的1011100011101011101001。
   - 假设HashMap长度默认是16，计算Length-1的结果为十进制15，二进制1111。
   - 把以上两数做**与运算**，101110001110101110 1001 & 1111 = 1001，十进制是9，所以index=9。

   所以，Hash算法最终得到的index结果，完全取决于Key的HashCode值的最后几位。因为长度为2的幂，Length-1以后的二级制值都是1，如2-1的二进制1，4-1的二进制11，8-1的二进制111。Length-1的二进制与HashCode进行**与运行**的结果，完全取决于HashCode的最后几位。只要输入的HashCode本身分布均匀，Hash算法的结果就是均匀的。

2. 高并发的情况下，为什么HashMap可能会出现死锁？（Java的HashMap是非线程安全的）

   简单来说，就是在HashMap存入新Entry产生冲突时，HashMap resize自身大小时，有哈希冲突的位置的链表形成了环型链表，造成infinite loop。Java官方推荐，在多线程高并发场景下使用ConcurrentHashmap。

   陈皓的博客CoolShell有这个问题的详细分析可参考

   https://coolshell.cn/articles/9606.html

1. 再Java8当中，HashMap的结构有什么样的优化？
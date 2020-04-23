### 解答开篇

问题：当你依次访问完一串页面 a-b-c 之后，点击浏览器的后退按钮，就可以查看之前浏览过的页面 b 和 a。当你后退到页面 a，点击前进按钮，就可以重新查看页面 b 和 c。但是，如果你后退到页面 b 后，点击了新的页面 d，那就无法再通过前进、后退功能查看页面 c 了。

如果你是 Chrome 浏览器的开发工程师，你会如何实现这个功能呢？

解答：

使用两个栈就，X 和 Y，把首次浏览的页面依次压入栈 X，当点击后退按钮时，再依次从战 X 中出栈，并将出栈的数据依次放入栈 Y。当我们点击前进按钮时，依次从栈 Y 中取出数据，放入栈 X。当栈 X 中没有数据时，就说明没有页面可以继续后退浏览了。当栈 Y 中没有数据，那就说明没有页面可以点击前进按钮浏览了。

### 课后思考

1. 我们在讲栈的应用时，讲到用函数调用栈来保存临时变量。为什么函数调用要用 “栈” 来保存临时变量呢？用其他数据结构不行吗？

   其实，我们不一定非要用栈来保存函数临时变量，只不过因为这个函数调用符合后进先出的特性，用栈这种数据结构来实现，是顺理成章的选择。

   从调用函数进入到被调用函数，对于数据来说，变化的是什么呢？是作用域。所以根本上，只要保证每进入一个新的函数，都是一个新的作用域就可以。而要实现这个，用栈就非常方便。在进入被调用函数的时候，分配一段栈空间给这个函数的变量，在函数结束的时候，将栈顶复位，正好回到调用函数的作用域内。

2. 我们都知道，JVM 内存管理中有个“堆栈”的概念。栈内存用来存储局部变量和方法调用，堆内存用来存储 Java 中的对象。那 JVM 里面的“栈”跟我们这里说的“栈”是不是一回事呢？如果不是，那它为什么又叫作“栈”呢？

   内存管理中的 ”堆栈“ 和数据结构的堆栈不是一个概念。内存空间在逻辑上分为三部分：代码区、静态数据区、动态数据区，而动态数据区又分为栈区和堆区。

   栈区：存储函数的形参、局部变量、返回值，由系统自动分配和回收。

   堆区：new一个对象时的引用或者地址存储在堆区，指向该对象存储在堆区中的真实数据。
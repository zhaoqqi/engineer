#### 插入排序算法

##### 描述
将数组的每个元素依次与其之前的元素比较，找到比当前元素小的元素M，将当前元素插入到M后面。再将X与M之间的元素都向右移动一位。   
排序过程中，当前索引左边的所有元素都是有序的，当索引到达数组的右端时，数组排序就完成了。   

##### 时间复杂度分析
对于随机排列的长度为N且主键不重复的数组，平均情况下插入排序需要 N平方/4 次比较以及N平方/4次交换。   
最坏情况下需要 N平方/2 次比较和 N平方/2 次交换。   
最好情况下需要 N-1 次比较和 0 次交换。

##### 特点
1. 对于有序数组进行排序时，插入排序能够立即发现每个元素都已经在合适的位置上，它的运行时间是线性的（N-1）。
2. 对于 部分有序 的数组，插入排序也是很有效的。   
部分有序包括：   
- 数组中每个元素距离它的最终位置均不远；
- 一个有序的大数组接一个小数组；
- 数组中只有几个元素的位置不正确；

##### 性能改进
如下改良版代码所示，将内循环的交换每个元素，改为将较大的元素向右移动。这样可以将访问数组的次数减半。


##### 算法实现
```java

public class Insertion {
    public static void sort(Comparable[] a) {
        int N = a.length;
        for (int i=1; i<N; i++) {
            for (int j=i; j>0 && a[j]<a[j-1]; j--) {
                exch(a, j, j-1)
            }
        }
    }

    //改良版-访问数组的次数减半
    public static void sort(Comparable[] a) {
        int N = a.length;
        for (int i=1; i<N; i++) {
            elem = a[i]
            index = 0
            for (int j=i; j>0 && a[j]<a[j-1]; j--) {
                a[j] = a[j-1]
                index = j
            }
            a[index] = elem
        }
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    private static void exch(Comparable[] a, int i, int j) {
        Comparable t = a[i];
        a[i] = a[j];
        a[j] = t;
    }

    private static void show(Comparable[] a) {
        for (int i = 0; i<a.length; i++) {
            StdOut.print(a[i] + " ");
        }
        StdOut.println();
    }

    public static boolean isSorted(Comparable[] a) {
        for (int i=1; i<a.length; i++) {
            if (a[i] > a[i+1]) return false;
        }
        return true;
    }

    public static void main(string[] args) {
        string[] a = In.readString();
        sort(a);
        assert isSorted(a);
        show(a);
    }
}
```

#### 选择排序算法

##### 描述
首先，找到数组中最小的元素，其次，将它与数组第一个元素交换位置（如果数组第一个元素就是最小的元素则自己跟自己交换）。   
再次，在剩下的元素中找到最小的元素，将它与数组的第二个元素交换位置。如此往复，直到将整个数组排序。

##### 时间复杂度分析
如果整个数组的大小是N，则整个排序过程要经过N次元素交换和 (N的平方)/2 次的元素比较。   
元素交换次数与数组大小N一致，这个理解起来很直观。   
元素的比较次数计算方法为 N+(N-1)+(N-2)+...+2+1 = (N的平方)/2 。   

##### 特点
不同于其他算法，排序算法具备以下两个特点：   
1. 与输入的顺序无关，哪怕输入的数组已经是排好序的，排序时间与随机排序的输入也是一致的（其他排序算法会根据输入顺序做优化）；
2. 数据移动次数最少；

##### 算法实现（java）

```java
public class Selection {
    public static void sort(Comparable[] a) {
        int N = a.length;
        for (int i=0; i<N; i++) {
            int min = i;
            for (int j=i+1; j<N; j++) {
                if (min < j) min = j;
            }
            exch(a, i, min);
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

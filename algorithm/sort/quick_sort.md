### 快速排序算法

#### 描述
快速排序是一种分治的排序算法。它将一个数组分成两个子数组，将两部分独立地排序。   
首先选择一个切分元素，将不大于切分元素的元素移动到左子数组，将不小于切分元素的元素移动到右子数组。   
然后将左子数字和右子数组使用同样的方法递归执行。用归纳法不难证明，如果左子数组和右子数组都是有序的，那么由左子数组、切分元素、右子数组组成的结果数组也一定是有序的。

#### 时间复杂度分析
快速排序算法的时间复杂度取决于切分元素的位置，如果刚好将数组切分为相等的两个子数组，则效果最佳。
平均时间复杂度为 NlogN   
最坏时间负责度为 N的平方/2   

#### 特点
快速排序算法是不稳定的排序算法。实现时要非常小心，才能避免低劣的性能。已经有无数例子显示许多错误都能致使它在实际中的性能只有平方级别。

#### 性能改进
1. 局部使用插入排序，完成小规模数组的排序。对于小规模数组来说，插入排序比快速排序快。
2. 使用三向切分快速排序算法。（具体算法后续补充）

#### 算法实现
```java

public class Quick {

    public static void sort(Comparable[] a) {
        StdRandom.shuffle(a);
        sort(a, 0, a.length-1);
    }

    private static void sort(Comparable[] a, int lo, int hi) {
        if (hi <= lo) return;
        int j = partition(a, lo, hi);
        sort(a, 0, j-1);
        sort(a, j+1, hi);
    }

    private static void partition(Comparable[] a, int lo, int hi) {
        int i = lo, j = hi + 1;
        Comparable v = a[lo];
        while (true) {
            while (less(a[++i], v)) if (i == hi) break;
            while (less(v, a[--j])) if (j == lo) break;
            if (i >= j) break;
            exch(a, i, j);
        }
        exch(a, lo, j);
        return j;
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
```
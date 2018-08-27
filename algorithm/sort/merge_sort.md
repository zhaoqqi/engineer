#### 归并排序

##### 描述
归并的思想是，将两个有序数组归并为一个更大的有序数组。根据这一思想，我们可以将数组分为左右两部分进行分别排序，然后再将结果归并起来。   

##### 时间复杂度分析
使用自上而下的分治思想，归并排序算法的排序时间与 NlogN 成正比。   
使用自下而上的思想，归并排序算法的排序时间需要 1/2NlgN ~ NlogN，最多访问数组 6NlogN 次。

##### 特点
指数级别的时间复杂度。

##### 算法实现
```java

public class Merge {
    private static Comparable[] aux;

    public static void sort(Comparable[] a) {
        aux = new Comparable[a.length];
        sort(a, 0, a.length-1);
    }

    // 自顶向下的归并排序，分治思想的典型应用
    private static void sort(Comparable[]a, int lo, int hi) {
        if (hi >= lo) return;
        int mid = lo + (hi - lo)/2;
        sort(a, 0, mid);
        sort(a, mid+1, hi);
        merge(a, lo, mid, hi);
    }

    // 自底向上的归并排序
    public static void sort(Comparable[] a) {
        int N = a.length;
        aux = new Comparable[N];
        for (int sz=1; sz<N; sz = sz+sz)
            for (int lo=0; lo<N-sz; lo+=sz+sz)
                merge(a, lo, lo+sz-1, Math.min(lo+sz+sz-1, N-1));
    }

    public static void merge(Comparable[] a, int lo, int mid, int hi) {
        int i = lo, j = mid + 1;
        for (int k=0; k<=hi, k++) {
            aux[k] = a[k];
        }

        for (int k=lo; k<=hi; k++) {
            if (i > lo) a[k] = aux[j++];
            else if (j > hi) a[k] = aux[i++];
            else if (less(a[i], a[j])) a[k] = aux[i++];
            else a[k] = aux[j++];
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
```
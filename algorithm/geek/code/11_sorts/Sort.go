package _11_sorts

/*
冒泡排序、插入排序、选择排序
 */

// 冒泡排序，a是数组，n是数组大小
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	// 惯常思维的双重循环处理
	/*
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}*/
	for i:=0; i<n; i++ {
		// 元素交换标志
		flag := true
		// 外围循环一次，数组最后一位就是已经有序的，所以内层循环每次少一次
		for j:=0; j<=n-i-1; j++ {
			if a[j] > j[j+1] {
				a[j], a[j+1] = a[j+1], a[i]
			}
		}
		// 没有元素交换时，提前退出
		if !flag {
			break
		}
	}
}

// 插入排序，a是数组，n是数组大小
func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i:=1; i<n; i++ {
		value := a[i]
		j := i - 1
		for ; j>=0; j-- {
			if a[j] > value {
				// 移动数据
				a[j] = a[j+1]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

// 选择排序，a是数组，n是数组大小
func SelectionSort(a []int, n int)  {
	if n <= 1 {
		return
	}
	for i:=0; i<n; i++ {
		minIndex := i
		for j:=i+1; j<n; j++ {
			// 寻找无序部分最小索引值
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		// 交换
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

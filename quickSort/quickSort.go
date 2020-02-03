package main
import (
	"fmt"
)

/**
通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。
*/
// 快速排序
// 说明
// 1. left 表示数组左边的下标
// 2. right 表示数组右边的下标

func QuickSort(array *[8]int, left int, right int) {
	if left < right {
		var i int = left
		var j int = right
		var temp int = array[left] 
		for ; i < j; {
			// 如果 array[j]  >= temp 就是符合条件的
			// 一定要加上 i < j ，不然加减的时候容易溢出
			for ; i < j && array[j] >= temp; {
				j-- // 前移
			}
			
			// 否则，将当前值放到前面
			if i < j {
				array[i] = array[j]
				i++
			}

			// 如果 array[i] <= temp 就是符合条件的
			for ; i < j && array[i] <= temp; {
				i++ // 后移
			}
			if i < j {
				array[j] = array[i]
				j--
			}
		}
		array[i] = temp
		QuickSort(array, left, i - 1)
		QuickSort(array, i + 1, right)
	}
}

func main() {
	arr := [8]int {-9, 78, 0, 23, -567, 0, 63, -5}
	fmt.Println("排序之前的数组 arr =", arr)
	QuickSort(&arr, 0, len(arr) - 1)
	fmt.Println("排序之后的数组 arr =", arr)
}
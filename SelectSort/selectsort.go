package main
import (
	"fmt"
)

/**
选择排序（Selection sort）是一种简单直观的排序算法。
它的工作原理是：第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，
然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。
以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。
*/
func SelectSort(arr *[5]int) {
	// 1. 先完成第一个最大值和 arr[0] 交换
	for i := 0; i < len(arr) - 1; i++ {
		max := arr[i]
		maxIndex := i
		// 遍历后面的进行比较
		for j := i + 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		// 交换
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}

func main() {

	// 定义一个数组，从小到大
	arr := [5]int{10, 35, 68, 52, 66}
	fmt.Println("选择排序之前：arr =", arr)
	SelectSort(&arr)
	fmt.Println("选择排序之后：arr =", arr)
}
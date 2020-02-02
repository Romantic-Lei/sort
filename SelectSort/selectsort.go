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
	// 外层循环，控制比较轮数
	for i := 0; i < len(arr) - 1; i++ {
		maxIndex := i // 假设当前值为最大值，保存其的索引
		// 遍历后面的进行比较
		for j := i + 1; j < len(arr); j++ {
			// 内层循环，控制每轮比较次数
			if arr[maxIndex] < arr[j] {
				maxIndex = j
			}
		}
		// 如果假设的最大值索引和比较之后的索引值不一致则需要交换
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
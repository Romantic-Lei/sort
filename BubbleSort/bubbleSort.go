package main
import (
	"fmt"
)

/**
它重复地走访过要排序的元素列，依次比较两个相邻的元素，
如果顺序（如从大到小、首字母从从Z到A）错误就把他们交换过来。
走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成
*/

// 冒泡排序
func BubbleSort(arr *[5] int) {
	fmt.Println("冒泡排序前：arr =", arr)
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j + 1] {
				temp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}
}

func main() {
	var arr [5]int = [5]int {4,3,5,6,1}
    BubbleSort(&arr)
    fmt.Println("main arr = ", arr) // 地址传递，此处的数组顺序也会被改变
}
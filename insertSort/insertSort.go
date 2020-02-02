package main
import (
	"fmt"
)

func InsertSort(arr *[7]int) {
	// 完成第一次,给第二个元素找到合适的位置并插入
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i] // 将当前值保存到临时变量中
		insertIndex := i - 1 // 当前值前面一位下标值
	
		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex-- // 索引前移
		}
		// 插入
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Printf("第%d次插入后%v\n", i, *arr)
	}
}

func main() {
	arr := [7]int{23, 0, 12, 56, 34, -1, 3}
	fmt.Println("原始数组是：arr =", arr)
	InsertSort(&arr)
	fmt.Println("排序之后的数组是：arr =", arr)
}
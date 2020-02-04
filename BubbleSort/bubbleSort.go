package main
import (
	"fmt"
	"math/rand"
	"time"
)

/**
它重复地走访过要排序的元素列，依次比较两个相邻的元素，
如果顺序（如从大到小、首字母从从Z到A）错误就把他们交换过来。
走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成
*/

// 冒泡排序
func BubbleSort(arr *[80000] int) {
	// fmt.Println("冒泡排序前：arr =", arr)
	// 外层循环，控制比较次数
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			// 内层循环，控制每轮比较次数
			if arr[j] > arr[j + 1] {
				temp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = temp
			}
		}
	}
}

func main() {
	// var arr [5]int = [5]int {4,3,5,6,1}
	var arr [80000]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		// func (r *Rand) Intn(n int) int
		// 返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	BubbleSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗费的时间 = %d 秒", end - start)
    // fmt.Println("冒泡排序后：arr = ", arr) // 地址传递，此处的数组顺序也会被改变
}
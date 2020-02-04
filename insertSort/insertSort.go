package main
import (
	"fmt"
	"math/rand"
	"time"
)

/**
插入排序法,插入排序的基本操作就是将一个数据插入到已经排好序的有序数据中，从而得到一个新的、个数加一的有序数据，
算法适用于少量数据的排序，时间复杂度为O(n^2)。是稳定的排序方法。
插入算法把要排序的数组分成两部分：第一部分包含了这个数组的所有元素，但将最后一个元素除外（让数组多一个空间才有插入的位置），
而第二部分就只包含这一个元素（即待插入元素）。在第一部分排序完成后，再将这个最后元素插入到已排好序的第一部分中。
*/
func InsertSort(arr *[80000]int) {
	// 完成第i次,给第i+1个元素找到合适的位置并插入
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i] // 将当前值保存到临时变量中
		insertIndex := i - 1 // 当前值前一位下标值
	
		// 从大到小
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex-- // 索引前移
		}
		// 插入
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertVal
		}
		// fmt.Printf("第%d次插入后%v\n", i, *arr)
	}
}

func main() {
	// arr := [7]int{23, 0, 12, 56, 34, -1, 3}
	// fmt.Println("原始数组是：arr =", arr)
	var arr [80000]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		// func (r *Rand) Intn(n int) int
		// 返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("插入排序耗费的时间 = %d 秒", end - start)
	// fmt.Println("排序之后的数组是：arr =", arr)
}
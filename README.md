# sort
###排序

###冒泡排序
它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果顺序（如从大到小、首字母从从Z到A）错误就把他们交换过来。走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素列已经排序完成

#####冒泡排序比较完相邻两个元素之后，符合比较条件就去交换

###选择排序
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是：第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。

#####选择排序每一轮比较之后，选出最大或者最小的值，一轮只交换一次

###插入排序

插入排序（Insertion Sorting）的基本思想是：把 n 个待排序的元素看成为一个有序表和一个无序表，开始时有序表中只包含一个元素，无序表中包含有 n-1 个元素，排序过程中每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码进行比较，将它插入到有序表中的适当位置，使之成为新的有序表。

#####插入排序每轮都是将前面的当做有序数组，后面的值和前面数组的值进行比较，直到找到符合条件的位置



### 快速排序

通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列

#####快速排序是以某一个数为标准，如果最终序列为从小到大，那么每一次比较，大于它的排在右边，小于它的排在左边，然后不断递归完成



###耗时比较

以本机器随机产生80000数据进行排序

冒泡排序(15秒)  **>** 选择排序(5秒) **>** 插入排序(2秒) **>** 快速排序(8000000数据, 1秒)



 
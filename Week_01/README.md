# 学习笔记
###  数据结构与算法总览
#### 数据结构
* 一维  
基础：数组array（string），链表 linked list  
高级：栈 stack，队列 queue，双端队列 deque，集合 set，映射 map（hash or map），等等  
* 二维：  
基础：树 tree，图graph
高级：二叉搜索树 binary search tree （red-black tree，AVL），堆 heap，并查集 disjoint set，字典树 Trie，等等  
* 特殊：  
位运算 Bitwise，布隆过滤器 BloomFilter  
缓存 LRU Cache  

#### 算法总览
所有的算法基石  
* if-else，switch --> branch 分支结构  
* for，while loop --> iteration 迭代（循环）
* 递归 Recursion（分治 Divide & Conquer，回溯 Backtrace）  
化繁为简，找到重复单元
* 搜索 Search 深度优先Depth first search，广度优先 Breadth first search ，A*，等等
* 动态规划 Dynamic Programming
* 二分查找 Binary Search
* 贪心 Greedy
* 数学 Math，几何 Geometry

### 数组、链表、跳表
#### 基本实现和特性
* 数组（array）  
在内存开辟了一段连续的内存地址，可以随机的访问任何一个元素，访问时间复杂度为O(1)  
新增、删除操作时间复杂度则是O(n)，因为数组的新增删除操作是要对某个索引后的所有元素进行位移  
数组实现，会涉及到数组的复制，挪动，新增空间，所以效率较低  
   * 时间复杂度：  
   prepend ：正常情况下是O(n)，但是可以优化到O(1)，实现方式：申请较大的空间，在数组的最开始预留一部分的空间，然后在epend操作时，把头指针前移一个位置即可  
   append：O(1)  
   lookup：O(1)  
   insert：O(n)  
   delete：O(n)
* 链表（linked list）  
一般定义为class 
链表中一个节点有两个成员变量，一个成员变量是value，另一个就成员就是指针，指针指向下一个节点，头指针用head标识，尾指针用tail表示，最后一个指针指向None（Null）  
单链表：只有一个next指针  
双链表：有next指针和prev（previous）指针，一个指向前，一个指向后  
循环链表：next指针指到前面的某一个节点  
查询操作时间复杂度是O(n)，因为它在内存中并不是一段连续的内存，而是指针指向下一个节点，所以查询的时候必须循环依次寻找，而新增或修改操作则是将原来的指针指向新的节点地址就好，时间复杂度是O(1)  
   * 时间复杂度  
   prepend ：O(1)  
   append：O(1)  
   lookup：O(n)  
   insert：O(1)  
   delete：O(1)
   
* 跳表（skip list）  
注意：只能用于元素有序的情况  
跳表（skip list）对标的是平衡树（AVL Tree）和二分查找，是一种插入、删除、搜索都是O(logn)的数据结构，1989年出现  
优势：原理简单、容易实现、方便扩展、效率高，因此有些热门的项目中替代平衡树，如Redis，LevelDB等  
实现思想：升维思想 + 空间换时间
实现方式：在原有的有序链表中增加一级索引、二级索引...一级索引，举个简单例子：每个指针指向2的1次幂节点（next的next），二级索引每个指针指向2的2次幂节点（next的next的next的next）  
   时间复杂度分析 
   n是节点的个数，以2为多级索引的跳跃比例，则  
   n/2、n/4、n/8...，第k级索引的节点个数就是n/(2^k)，这里要保证n >= 2（两个节点的链表），且n/(2^k) > 1，不然没什么意义  
   假设索引有h级，最高级的索引有2个节点，则n/2^h = 2，从而求得h = log2(n) - 1  
* 现实中的跳表形态  
由于在实际操作过程中，由于节点的增加或删除，会导致索引不是完全工整的，有些地方索引会跨几步，有些地方索引会少跨几步，且维护成本高，新增或删除需要把索引都更新，因为存在多级索引的情况，所以空间复杂度为O(n)  
   * 时间复杂度
   prepend ：O(1)  
   append：O(1)  
   lookup：O(logn)  
   insert：O(logn)  
   delete：O(logn)

### 栈和队列
stack：栈，后进先出 last in first out LIFO  插入、删除都是O(1)，查询是O(n)  
queue：队列，先进先出，first in first out FIFO  插入、删除都是O(1)，查询是O(n)  

deque：Double-End Queue 双端队列，插入、删除都是O(1)，查询是O(n)  
Priority queue：优先队列，插入操作是O(1)，取出操作是O(nlogn)，按照元素的优先级取出  
底层具体实现的数据结构较为多样和复杂：heap、BST（binary search tree）、treap， 
```go
package main

import (
	"errors"
	"fmt"
)


//使用数组模拟栈的实现
type Stack struct {
	MaxTop int //表示最大栈存放个数
	Top int //顶栈
	arr [5]int  //数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error)  {
	//判断栈是否满
	if this.Top == this.MaxTop -1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	//
	this.Top++
	//入栈
	this.arr[this.Top] = val
	return
}

func (this *Stack) List() {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i:=this.Top; i>=0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	//先取值，再top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}
```

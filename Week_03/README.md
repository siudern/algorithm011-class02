# 学习笔记
### 泛型递归、树的递归

#### 递归的实现、特性以及思维要点

递归 Recursion  
递归 - 循环  
通过函数体来进行的循环  

递归的写法：  
1、首先名明确中止条件  
2、处理当前逻辑  
3、下探到下一层  
4、清理当前层  

```go
//递归伪代码
func rescursion(level, param1, param2, ...) {
    //中止条件
    if level > maxLevel {
    	process_result
        return

    }
    //处理当前逻辑
    process(level, data...)
    //下探到下一层
    rescursion(level+1, p1, p2, ...)
    //清理当前层
}
```

三个思维要点：  
1、抵制人肉递归  
2、找最近重复性  
3、数学归纳法思维  


### 分治、回溯
#### 分治
本质上是递归，特殊的递归  
```go
//分治伪代码
func divideConquer(problem, param1, param2, ...) {
    //中止条件
    if problem == nil {
        print_result
        return
    }
    // 准备数据，处理当前层逻辑
    data = prepare_data(problem)
    subproblems = splitProblem(problem, data)
    //下探到下一层，处理子问题
    subResult1 = divideConquer(subproblems[0], p1, p2, ...)
    subResult2 = divideConquer(subproblems[1], p1, p2, ...)
    subResult3 = divideConquer(subproblems[2], p1, p2, ...)
    ...
    //处理和生成最终结果
    result = processResult(subResult1, subResult2, subResult3, ...)
    //清理当前层状态
}
```

#### 回溯
回溯法采用试错的思想，它尝试分步的去解决一个问题，在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能分步解答再次尝试寻找问题的答案。   
回溯法通常用最简单的递归方法来实现，在反复重复上述的步骤后可能出现两种情况：  
* 找到一个可能存在的正确答案  
* 在尝试了所有可能的分步方法后宣告该问题无解  
在最坏情况下，回溯法的时间复杂度是指数级

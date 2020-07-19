# 学习笔记

搜索：
* 每个节点都要访问一次  
* 每个节点仅访问一次  


## 深度优先和广度优先

### 深度优先
对于节点的访问顺序不限  
- 深度优先  depth first search  
- 广度优先  breadth first search  
- 优先级优先  


深度优先的递归写法：  
```python
visited = set()
def dfs(node):
    if node in visited:
        # already visited
        return
    visited.add(node)
    # process current node
    # ... logic here
    dfs(node.left)
    dfs(node.right)

def dfs2(node, visited):
    visited.add(node)
    # process current node  here.
    # ...
    for next_node in node.children():
        if not next_node in visited:
            dfs2(next_node, visited)


def dfs3(node, visited):
    if node in visited:
        # already visited, 中止条件
        return
    visited.add(node)

    # process current node here.
    # ...
    pass

    for next_node in node.children():
        if not next_node in visited:
            dfs(next_node, visited)
```

深度优先的非递归写法
```python
class Search:
    def generate_related_nodes(self, node):
        pass

    def dfs(self, tree):
        if tree.root is None:
            return []

        visited, stack = [], [tree.root]

        while stack:
            node = stack.pop()
            visited.add(node)
            # process current node here
            # ... 
            pass
            nodes = self.generate_related_nodes(node)
            stack.push(nodes)
         # other processing work
```

### 广度优先
```python
visited = []
class Search:
    def generate_related_nodes(self, node):
        pass

    def bfs(self, graph, start, end):
        queue = []
        queue.append([start])
        visited.add(start)
        
        while queue:
            node = queue.pop()
            visited.add(node)
    
            # process logic
            # ...
            pass
            nodes = self.generate_related_nodes(node)
            queue.push(nodes)
    
         # other processing work
```

### 贪心算法 Greedy
贪心算法是一种在每一步选择都才去在当前状态下最好或最优解（即最有利）的选择，从而希望使结果也是全局最优解的算法。  
贪心算法与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能  

贪心：当下做局部最优解  
回溯：能够回退
动态规划：最优判断+回退  

贪心算法可以解决一些最优化问题，如：求图中的最小生成树，求哈夫曼编码等，然而对于工程和生活中的问题，贪心一般不能得到所要求的答案。  
一旦一个问题可以通过贪心算法来解决，那么贪心算法一般是解决这个问题的最好办法。由于贪心算法的高效性以及其所求得的答案比较接近最优结果，贪心算法也可以做辅助算法或者直接解决一些要求结果不是特别精确的问题。  

贪心算法的适用场景
问题能够分解为子问题  
子问题最优解能递推到最终问题的最优解，这种子问题的最优解称为最优子结构

### 二分查找
首先，假设表中元素是按升序排列，将表中间位置记录的关键字与查找关键字比较，如果两者相等，则查找成功；否则利用中间位置记录将表分成前、后两个子表，如果中间位置记录的关键字大于查找关键字，则进一步查找前一子表，否则进一步查找后一子表。重复以上过程，直到找到满足条件的记录，使查找成功，或直到子表不存在为止，此时查找不成功  
能用二分查找的条件：  
1. 目标函数单调性（单调递增或递减），即排序后的  
2. 存在上下界（bounded）  
3. 能够通过索引访问（index accessible）  

代码模板
```go
func BinarySearch(n []int, target int) int {
    length := len(n)   
    low := 0   
    high := length - 1
    for low <= high {
         mid := (low + high) / 2      
         if n[mid] > target {
             high = mid - 1
         } else if n[mid] < target {
             low = mid + 1      
         } else {         
             return mid      
         }   
    }
    return -1
} 
```


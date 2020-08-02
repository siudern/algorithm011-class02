学习笔记

# 动态规划  Dyanmic Programming

动态规划的特点：
分治 + 最优子结构

动态规、递归、分治，没有本质上的区别（关键看有无最优子结构）  
共性：找到重复子问题  
差异性：最优子结构、中途可以淘汰次优解  

状态转移方程 DP
```python
opt[i, j] = opt[i+1, j] + opt[i, j+1]
# 完整逻辑：
if a[i, j] = "空地"
    opt[i, j] = opt[i+1, j] + opt[i, j+1]
else:
    opt[i, j] = 0
```

动态规划关键点  
1. 最优子结构 opt[n] = best_of(opt[n-1], opt[n-2], ...)  
2. 存储中间状态opt[i]  
3. 递推公式（状态转移方程，或者叫DP方程）  
  fib：opt[i] = opt[n-1] + opt[n-2]  
  二维路径：opt[i,j] = opt[i+1, j] +opt[i, j+1] 且判断opt[i, j]是否为空  


动态规划小结  
1. 打破自己的思维惯性，形成机器思维  
2. 理解复杂逻辑的关系  

    
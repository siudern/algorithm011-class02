学习笔记
# 字典树和并查集

## 字典树 Trie
1、字典树的树结构  
* 字典树，即Trie树，又称单词查找树或键树，是一种属性结构。典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计  
* 优点：最大限度减少无所谓的字符串比较，查询效率比哈希高  
* 多叉树  

2、字典树的核心思想  
* 空间换时间  
* 利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的

3、字典树的基本性质  
* 节点本身不存完整单词  
* 从根节点到某一节点，路径上锁经过的字符连起来，为该节点对应的字符串  
* 每个节点的所有子节点路径代表的字符都不相同  
```python
class Trie(object):
    def __init__(self):
        self.root = {}
        self.end_of_word = "#"

    def insert(self, word):
        node = self.root
        for char in word:
            node = node.setdefault(char, {})
        node[self.end_of_word] = self.end_of_word

    def search(self, word):
        node = self.root
        for char in word:
            if char not in node:
                return False
            node = node[char]
        return self.end_of_word in node

    def startswith(self, prefix):
        node = self.root
        for char in prefix:
            if char not in node:
                return False
        return True
```

### 并查集 Disjoint Set
使用场景：  
组团、配对问题  
group or not  

基本操作：  
* make Set(s)：建立一个新的并查集，其中包含S个单元素集合。
* union Set（x, y)：把元素x和元素y所在的集合合并，要求x和y所在的集合不相交，如果相交则不合并。
* find(x)：找到元素x所在的集合的代表，该操作也可以用于判断两个元素是否位于同一个集合，只要将他们各自的代表比较一下就可以了

### 高级搜索
* 剪枝
* 双向BFS
* 启发式搜索（A*)

初级搜索：
1. 朴素搜索
2. 优化方式：不重复（fibonacci）、剪枝（生成括号问题）
3. 搜索方向：
   * DFS
   * BFS
   双向搜索、启发式搜索

启发式函数：h(n)，它用来评价哪些节点最优希望的是一个我们要找的节点，h(n)会返回一个非负整数，也可以认为是从节点n的目标节点路径的估计成本  
启发式函数是一种告知搜索方向的方法，它提供了一种明智的方法来猜测哪个邻居节点会导向一个目标  

### 高级树、AVL树和红黑树

#### AVL:
Balance Factor（平衡因子）：  
是它左子树的高度减去它右子树的高度（有时相反）。  
balance factor {-1,0,1}  
旋转操作
1. 左旋
2. 右旋
3. 左右旋
4. 右左旋

总结
1. 平衡二叉树搜索
2. 每个节点存balance factor = {-1,0,1}
3. 四种旋转操作
不足：节点需要存储额外信息，且调整次数频繁  

#### 红黑树
红黑树是一种近似平衡的二叉搜索树，他能够确保任何一个节点的左右子树的高度差小于两倍。具体来说，红黑树是满足如下条件的二叉搜索树
* 每个节点要么是红色，要么是黑色
* 根节点是黑色
* 每个叶子节点（NIL节点，空节点）是黑色。
* 不能有相邻的两个红色节点
* 从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点





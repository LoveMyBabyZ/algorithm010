# 学习笔记
## 广度优先
广度优先搜索（Breadth-First-Search），我们平常都简称 BFS。直观地讲，它其实就是一种“地毯式”层层推进的搜索策略，即先查找离起始顶点最近的，然后是次近的，依次往外搜索

```
def BFS(graph, start, end): 
     queue = [] 
     queue.append([start]) 
     visited.add(start) 
    while queue: 
     node = queue.pop() 
     visited.add(node) 
     process(node) 
     nodes = generate_related_nodes(node) 
     queue.push(nodes)
```
## 深度优先
深度优先搜索用的是一种比较著名的算法思想，回溯思想。这种思想解决问题的过程，非常适合用递归来实现

```

visited = set() 
def dfs(node, visited): 
    if node in visited: # terminator 
    # already visited 
    return 
     visited.add(node) 
    # process current node here. 
    ...
    for next_node in node.children(): 
    if not next_node in visited: 
     dfs(next_node, visited)
  
```
## 贪心算法
贪心算法是一种在每一步选择中都采取在当前状态下最好或最优（即最有
利）的选择，从而希望导致结果是全局最好或最优的算法

## 二分查找

1. 目标函数单调性（单调递增或者递减）

2. 存在上下界（bounded）

3. 能够通过索引访问（index accessible)

### 代码模板

```
left, right = 0, len(array) - 1
   while left <= right:
   mid = (left + right) / 2
  if array[mid] == target: # find the target!! 
  break or return result 
   elif array[mid] < target:``
   left = mid + 1
  else:
   right = mid - 1
   
```

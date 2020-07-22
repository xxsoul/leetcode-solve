# 2020-07

## 目录
| 题目 | 文件 |
| ---- | ---- |
| [不同的二叉搜索树II](#不同的二叉搜索树II) | 0721.go |
| [旋转数组的最小数字](#旋转数组的最小数字) | 0722.go |

## 题目
### 不同的二叉搜索树II
给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。</br>
示例：</br>
输入：3
输出：
```
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
```
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：
```

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```
提示：0 <= n <= 8
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

### 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  </br>
示例 1：</br>
输入：[3,4,5,1,2]</br>
输出：1</br>
示例 2：</br>
输入：[2,2,2,0,1]</br>
输出：0</br>
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
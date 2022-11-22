package _21010

import (
	"fmt"
	"testing"
)

/*
对于下标 iii，下雨后水能到达的最大高度等于下标 iii 两边的最大高度的最小值，下标 iii 处能接的雨水量等于下标 iii 处的水能到达的最大高度减去 height[i]\textit{height}[i]height[i]。

朴素的做法是对于数组 height\textit{height}height 中的每个元素，分别向左和向右扫描并记录左边和右边的最大高度，然后计算每个下标位置能接的雨水量。假设数组 height\textit{height}height 的长度为 nnn，该做法需要对每个下标位置使用 O(n)O(n)O(n) 的时间向两边扫描并得到最大高度，因此总时间复杂度是 O(n2)O(n^2)O(n
2
 )。

上述做法的时间复杂度较高是因为需要对每个下标位置都向两边扫描。如果已经知道每个位置两边的最大高度，则可以在 O(n)O(n)O(n) 的时间内得到能接的雨水总量。使用动态规划的方法，可以在 O(n)O(n)O(n) 的时间内预处理得到每个位置两边的最大高度。

创建两个长度为 nnn 的数组 leftMax\textit{leftMax}leftMax 和 rightMax\textit{rightMax}rightMax。对于 0≤i<n0 \le i<n0≤i<n，leftMax[i]\textit{leftMax}[i]leftMax[i] 表示下标 iii 及其左边的位置中，height\textit{height}height 的最大高度，rightMax[i]\textit{rightMax}[i]rightMax[i] 表示下标 iii 及其右边的位置中，height\textit{height}height 的最大高度。

显然，leftMax[0]=height[0]\textit{leftMax}[0]=\textit{height}[0]leftMax[0]=height[0]，rightMax[n−1]=height[n−1]\textit{rightMax}[n-1]=\textit{height}[n-1]rightMax[n−1]=height[n−1]。两个数组的其余元素的计算如下：

当 1≤i≤n−11 \le i \le n-11≤i≤n−1 时，leftMax[i]=max⁡(leftMax[i−1],height[i])\textit{leftMax}[i]=\max(\textit{leftMax}[i-1], \textit{height}[i])leftMax[i]=max(leftMax[i−1],height[i])；

当 0≤i≤n−20 \le i \le n-20≤i≤n−2 时，rightMax[i]=max⁡(rightMax[i+1],height[i])\textit{rightMax}[i]=\max(\textit{rightMax}[i+1], \textit{height}[i])rightMax[i]=max(rightMax[i+1],height[i])。

因此可以正向遍历数组 height\textit{height}height 得到数组 leftMax\textit{leftMax}leftMax 的每个元素值，反向遍历数组 height\textit{height}height 得到数组 rightMax\textit{rightMax}rightMax 的每个元素值。

在得到数组 leftMax\textit{leftMax}leftMax 和 rightMax\textit{rightMax}rightMax 的每个元素值之后，对于 0≤i<n0 \le i<n0≤i<n，下标 iii 处能接的雨水量等于 min⁡(leftMax[i],rightMax[i])−height[i]\min(\textit{leftMax}[i],\textit{rightMax}[i])-\textit{height}[i]min(leftMax[i],rightMax[i])−height[i]。遍历每个下标位置即可得到能接的雨水总量。



作者：力扣官方题解
链接：https://leetcode.cn/problems/trapping-rain-water/solutions/692342/jie-yu-shui-by-leetcode-solution-tuvc/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func TestRain(t *testing.T) {
	arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	//fmt.Println(trap(arr))
	fmt.Println(trapNew(arr))
}

func trap(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

func trapNew(height []int) (ans int){
	n := len(height)
	if n == 0 {
		return 0
	}



	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i - 1], height[i])
	}

	rightMax[n - 1] = height[n - 1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i + 1], height[i])
	}

	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


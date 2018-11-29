package twosum

import "errors"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的 两个 整数。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9

所以返回 [0, 1]
 */

/*
暴力解法 最快为60ms

时间复杂度：O(n^2)， 对于每个元素，我们试图通过遍历数组的其余部分
来寻找它所对应的目标元素，这将耗费 O(n) 的时间。因此时间复杂度为 O(n^2)

空间复杂度：O(1)
)。
*/

func twosumOne(nums []int, target int) []int {
	var i,j int
	length := len(nums)
	for i = 0 ; i < length - 1 ; i++ {
		for j = i + 1; j < length; j ++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	panic(errors.New("No two sum solution"))
}


/*
两遍哈希表 最快为8ms

通过以空间换取速度的方式，我们可以将查找时间从 O(n) 降低到 O(1)。
哈希表正是为此目的而构建的，它支持以 近似 恒定的时间进行快速查找。我用“近似”来描述，
是因为一旦出现冲突，查找用时可能会退化到 O(n)。但只要你仔细地挑选哈希函数，
在哈希表中进行查找的用时应当被摊销为 O(1)。

一个简单的实现使用了两次迭代。在第一次迭代中，我们将每个元素的值和它的索引添加到表中。
然后，在第二次迭代中，我们将检查每个元素所对应的目标元素（target - nums[i]）
是否存在于表中。注意，该目标元素不能是 nums[i] 本身！

时间复杂度：O(n)， 我们把包含有 n 个元素的列表遍历两次。
由于哈希表将查找时间缩短到 O(1) ，所以时间复杂度为 O(n)。

空间复杂度：O(n)， 所需的额外空间取决于哈希表中存储的元素数量，该表中存储了 n 个元素。
 */
func twosumTwo(nums []int, target int) []int {
	map1 := make(map[int]int)
	length := len(nums)
	for i := 0; i < length; i++ {
		map1[nums[i]] = i
	}

	for j := 0; j < len(nums); j++ {
		elem := target - nums[j]
		if _, ok := map1[elem];ok && map1[elem] != j {
			return []int{j,map1[elem]}
		}
	}
	panic(errors.New("No two sum solution"))
}


/*
一遍哈希表 最快为8 ms

事实证明，我们可以一次完成。在进行迭代并将元素插入到表中的同时，
我们还会回过头来检查表中是否已经存在当前元素所对应的目标元素。
如果它存在，那我们已经找到了对应解，并立即将其返回。

时间复杂度：O(n)， 我们只遍历了包含有 n 个元素的列表一次。在表中进行的每次查找只花费 O(1) 的时间。

空间复杂度：O(n)， 所需的额外空间取决于哈希表中存储的元素数量，该表最多需要存储 n 个元素。
 */
func twosumThree(nums []int, target int) []int {
	map1 := make(map[int]int)

	for i:= 0;i < len(nums); i++ {
		elem := target - nums[i]
		if _, ok := map1[elem]; ok {
			return []int{map1[elem],i}
		}
		map1[nums[i]] = i
	}
	panic(errors.New("No two sum solution"))
}
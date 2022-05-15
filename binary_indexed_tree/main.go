package main

import "fmt"

type NumArray struct {
	bit  BIT
	nums []int
}

func Constructor(nums []int) NumArray {
	b := BIT{
		tree: make([]int, len(nums)+1),
	}
	for i, v := range nums {
		b.Update(i, v)
	}
	return NumArray{b, nums}
}

func (this *NumArray) Update(index int, val int) {
	this.bit.Update(index, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.bit.Query(right) - this.bit.Query(left-1)
}

type BIT struct {
	tree []int
}

func (this *BIT) Update(id int, delta int) {
	id++
	for id < len(this.tree) {
		this.tree[id] += delta
		id += -id & id
	}
}

func (this *BIT) Query(id int) int {
	id++
	s := 0
	for id > 0 {
		s += this.tree[id]
		id -= -id & id
	}
	return s
}

func main() {
	arr := []int{1, 3, 5}
	numArray := Constructor(arr)
	s1 := numArray.SumRange(0, 2)
	fmt.Println(s1)
	s2 := numArray.SumRange(0, 2)
	numArray.Update(1, 2)
	fmt.Println(s2)
}

func maximumCandies(candies []int, k int64) int {

	var sum int64
	for i := 0; i < len(candies); i++ {
		sum += int64(candies[i])
	}
	l := 0
	r := int(sum / k)
	for l < r {
		mid := l + (r-l+1)/2
		if canArraive(candies, k, mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l

}

func canArraive(candies []int, k int64, count int) bool {
	var cnt int64
	for _, v := range candies {
		cnt += int64(v / count)
		if cnt >= k {
			return true
		}
	}
	return false
}

func projectionArea(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ret := 0
	max := 0
	for i := 0; i < m; i++ {
		max = 0
		for j := 0; j < n; j++ {
			max = maa(max, grid[i][j])
		}
		ret += max
	}
	for i := 0; i < n; i++ {
		max = 0
		for j := 0; j < m; j++ {
			max = maa(max, grid[j][i])
			ret += min(1, grid[j][i])
		}
		ret += max
	}
	return ret

}

func maa(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

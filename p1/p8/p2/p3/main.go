package main

type Node struct {
	pre *Node
	aft *Node
	val int
}

func findTheWinner(n int, k int) int {
	if n == 1 {
		return n
	}
	head := Node{
		val: 1,
	}
	tail := Node{
		val: n,
	}
	//建环
	head.pre = &tail
	tail.aft = &head
	befo := &head
	for i := 2; i < n; i++ {
		cur := Node{
			val: i,
		}
		cur.pre = befo
		befo.aft = &cur
		befo = &cur
	}
	befo.aft = &tail
	tail.pre = befo
	//遍历
	tmp := head
	for cnt := 0; cnt < n-1; cnt++ {
		for i := 1; i < k; i++ {
			tmp = *tmp.aft
		}
		p := tmp.pre
		l := tmp.aft
		p.aft = l
		l.pre = p
		tmp = *l
	}
	return tmp.val

}

func main() {
	n, k := 5, 2
	println(findTheWinner(n, k))
}

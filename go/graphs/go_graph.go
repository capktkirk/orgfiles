package main

func BFS(tree *Tree) []int {
	queue := []*Tree{}
	queue = append(queue, tree)
	result := []int{}
	return BFSUtil(queue, result)
}

func BFSUtil(queue []*Tree, res []int) []int {
	if len(queue) == 0 {
		return res
	}
	res = append(res, queue[0].Value)
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	return BFSUtil(queue[1:], res)
}

func main() {

}

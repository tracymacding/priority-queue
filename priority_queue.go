package priority_queue

import (
	"container/heap"
)

type Node struct {
	Key      interface{}
	Value    interface{}
	Priority int
	Index    int
}

func NewNode(key interface{}, v interface{}, priority int) *Node {
	return &Node{
		Key:      key,
		Value:    v,
		Priority: priority,
		Index:    -1,
	}
}

func (n *Node) GetKey() interface{} {
	return n.Key
}

func (n *Node) GetValue() interface{} {
	return n.Value
}

func (n *Node) GetIndex() int {
	return n.Index
}

func (n *Node) UpdatePrio(newPrio int) {
	n.Priority = newPrio
}

type Nodes []*Node

func (nodes Nodes) Len() int { return len(nodes) }

func (nodes Nodes) Less(i, j int) bool { return nodes[i].Priority < nodes[j].Priority }

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
	nodes[i].Index = i
	nodes[j].Index = j
}

func (nodes *Nodes) Push(v interface{}) {
	node := v.(*Node)
	node.Index = len(*nodes)
	*nodes = append(*nodes, node)
}

func (nodes *Nodes) Pop() interface{} {
	old := *nodes
	size := len(old)
	node := old[size-1]
	// for safety
	node.Index = -1
	*nodes = old[0 : size-1]
	return node
}

type PriorityQueue struct {
	nodes Nodes
}

func (pq *PriorityQueue) Push(n *Node) {
	heap.Push(&(pq.nodes), n)
}

func (pq *PriorityQueue) Pop() *Node {
	if len(pq.nodes) <= 0 {
		return nil
	}
	n := heap.Pop(&(pq.nodes))
	return n.(*Node)
}

func (pq *PriorityQueue) Remove(index int) {
	if index < 0 || index >= len(pq.nodes) {
		return
	}
	heap.Remove(&(pq.nodes), index)
}

func (pq *PriorityQueue) Length() int {
	return len(pq.nodes)
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{nodes: make(Nodes, 0, 1024)}
	heap.Init(&(pq.nodes))
	return pq
}

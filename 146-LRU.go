//Design a data structure that follows the constraints of a Least Recently Used
//(LRU) cache.
//
// Implement the LRUCache class:
//
//
// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//
// int get(int key) Return the value of the key if the key exists, otherwise
//return -1.
// void put(int key, int value) Update the value of the key if the key exists.
//Otherwise, add the key-value pair to the cache. If the number of keys exceeds
//the capacity from this operation, evict the least recently used key.
//
//
// The functions get and put must each run in O(1) average time complexity.
//
//
// Example 1:
//
//
//Input
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//Output
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//Explanation
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // cache is {1=1}
//lRUCache.put(2, 2); // cache is {1=1, 2=2}
//lRUCache.get(1);    // return 1
//lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
//lRUCache.get(2);    // returns -1 (not found)
//lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
//lRUCache.get(1);    // return -1 (not found)
//lRUCache.get(3);    // return 3
//lRUCache.get(4);    // return 4
//
//
//
// Constraints:
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10⁴
// 0 <= value <= 10⁵
// At most 2 * 10⁵ calls will be made to get and put.
//
//
// Related Topics Hash Table Linked List Design Doubly-Linked List 👍 14496 👎 5
//75

package main

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	Head     *Node
	Tail     *Node
	CacheMap map[int]*Node
	Size     int
}

type Node struct {
	Key      int
	Val      int
	PreNode  *Node
	NextNode *Node
}

func LRUConstructor(capacity int) LRUCache {
	cacheMap := LRUCache{
		Head:     &Node{},
		Tail:     &Node{},
		CacheMap: make(map[int]*Node, capacity),
		Size:     capacity,
	}
	cacheMap.Head.NextNode = cacheMap.Tail
	cacheMap.Tail.PreNode = cacheMap.Head
	return cacheMap
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.CacheMap[key]; !ok {
		return -1
	}
	node := this.CacheMap[key]
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.CacheMap[key]; !ok {
		node := &Node{
			Key: key,
			Val: value,
		}
		this.CacheMap[key] = node
		this.addToHead(node)

		if len(this.CacheMap) > this.Size {

			delNode := this.removeTail()
			delete(this.CacheMap, delNode.Key)
		}
	} else {
		node := this.CacheMap[key]
		node.Val = value
		this.moveToHead(node)
	}

	return
}

func (this *LRUCache) addToHead(node *Node) {
	node.PreNode = this.Head
	node.NextNode = this.Head.NextNode
	this.Head.NextNode.PreNode = node
	this.Head.NextNode = node

}

func (this *LRUCache) removeNode(node *Node) {

	node.PreNode.NextNode = node.NextNode
	node.NextNode.PreNode = node.PreNode
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.Tail.PreNode
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

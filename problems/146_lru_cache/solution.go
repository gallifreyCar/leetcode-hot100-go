package lrucache

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	dummy    *Node //铏氭嫙澶磋妭鐐?
	tail     *Node //铏氭嫙灏捐妭鐐?
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {

	dummy := &Node{}
	tail := &Node{}
	dummy.next = tail
	tail.prev = dummy

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		dummy:    dummy,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

// 鏂板厓绱犳坊鍔犲埌澶磋妭鐐?
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.dummy
	node.next = this.dummy.next
	this.dummy.next.prev = node
	this.dummy.next = node
}

// 绉婚櫎鑺傜偣
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 绉婚櫎灏捐妭鐐?瑕佽繑鍥炵粰鍝堝笇琛ㄥ垹闄?
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 绉诲姩鍒板ご鑺傜偣
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	//瀛樺湪
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.addToHead(node)

	if len(this.cache) > this.capacity {
		removeTail := this.removeTail()
		delete(this.cache, removeTail.key)
	}

}

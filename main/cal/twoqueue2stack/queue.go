package twoqueue2stack

/*
队列的特性较为单一，基本操作即初始化、获取大小、添加元素、移除元素等。
最重要的特性就是满足先进先出
*/
type Queue struct {
	element []interface{}
}

func New() *Queue {
	return new(Queue)
}

//加入队列
func (this *Queue) Put(data interface{}) {
	this.element = append(this.element, data)
}

//pop出队列
func (this *Queue) Pop() interface{} {
	if this.empty() {
		return nil
	}

	data := this.element[0]
	this.element = this.element[1:]
	return data
}

func (this *Queue) empty() bool{
	return len(this.element) == 0
}

//获得队列的长度
func (this *Queue) GetSize() int {
	return len(this.element)
}

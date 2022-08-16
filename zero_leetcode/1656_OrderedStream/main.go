package _656_OrderedStream

//leetcode submit region begin(Prohibit modification and deletion)
type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n+1),
		ptr:    1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	if idKey < 1 {
		return []string{}
	}
	var res []string
	this.stream[idKey] = value
	if idKey == this.ptr {
		for idKey < len(this.stream) && len(this.stream[idKey]) > 0 {
			res = append(res, this.stream[idKey])
			idKey++
		}
		this.ptr = idKey
	}
	return res
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

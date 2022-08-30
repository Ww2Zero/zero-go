package _32_findItinerary

import (
	"fmt"
	"sort"
)

/*建立每个起点所有能到达终点的映射,这个映射值便是要遍历的集合
  为了避免起点到终点的重复访问,需要为每个机场设置一个标志位,来标记是否被访问过。
  如：A -> B, B -> A, 接下来A -> C, 但是仍然可以有A -> B,为了避免这种情况，就需要在B 设置个标志位,代表已经访问    过, 当然这个访问仅代表A -> B访问过了,至于其他到B的依旧可以访问
*/

//结点用于存机场名和访问标志
type fromTo struct {
	dst    string
	visted bool
}
type fromToSlice []fromTo

//Len 使用Sort排序, 必须实现其对应的Interface接口中的三个方法
func (f fromToSlice) Len() int {
	return len(f)
}
func (f fromToSlice) Less(i, j int) bool {
	return f[i].dst < f[j].dst //按机场名的字典序排序
}
func (f fromToSlice) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func findItinerary(tickets [][]string) (res []string) {
	var backTracking func(string) bool
	m := make(map[string]fromToSlice)
	//建立映射关系
	for _, v := range tickets {
		if _, ok := m[v[0]]; ok { //如果该起点存在,直接追加 [注]其实不论存在与否,都可以直接追加,这样分开写效率高
			m[v[0]] = append(m[v[0]], fromTo{dst: v[1]})
		} else { //不存在直接赋值
			m[v[0]] = fromToSlice{fromTo{dst: v[1]}}
		}
	}
	//对每一个映射值按字典序排序
	for _, v := range m {
		sort.Sort(v)
	}
	fmt.Println(m)
	//回溯算法模板
	res = append(res, "JFK")
	backTracking = func(startIndex string) bool {
		if len(res) == len(tickets)+1 {
			return true
		}

		for i := 0; i < len(m[startIndex]); i++ {
			if m[startIndex][i].visted == true {
				continue
			}
			res = append(res, m[startIndex][i].dst)
			fmt.Println(res)
			m[startIndex][i].visted = true
			if backTracking(m[startIndex][i].dst) == true {
				return true
			}
			m[startIndex][i].visted = false
			res = res[:len(res)-1]
		}
		return false
	}
	backTracking("JFK")
	return
}

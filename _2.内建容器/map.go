package main

import "fmt"

/*
	Map

*/
//寻找最长不含有重复字符的子串
func lengthByString(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		//fmt.Println(i,ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
		//fmt.Println(lastOccurred)
	}
	return maxLength
}
func main() {
	m := map[string]string{
		"name":    "xu",
		"course":  "golang",
		"site":    "fekerxu",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m, m2, m3)     //map[course:golang name:xu quality:notbad site:fekerxu] map[] map[]
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
		/*	无序打印
			name xu
			course golang
			site fekerxu
			quality notbad
		*/
	}
	fmt.Println("Getting values") //以键取值
	courseName, ok := m["course"]
	fmt.Println(courseName, ok) //golang true
	if courseName, ok := m["cours"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist") //key does not exist
	}
	fmt.Println("删除值")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name") //xu true
	name, ok = m["name"]
	fmt.Println(name, ok) //false
	/*
		创建:make(map[string]int)
		获取元素:m[key]
		key不存在，获得Value类型的初始值
		用value,ok := m[key]来判断是否存在key
		用delete删除一个key
		使用range 遍历key,或者遍历key,value对
		不保证遍历顺序，如需顺序，需手动对key排序
		map使用哈希表 必须可以比较相等
		除了slice,map,function的内建类型都可以作为key
		struct类型不包含上述字段，也可作为key
	*/
	//
	fmt.Println(lengthByString("abc"))    //3
	fmt.Println(lengthByString("bbb"))    //1
	fmt.Println(lengthByString(""))       //0
	fmt.Println(lengthByString("abcabc")) //3
}

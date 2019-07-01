package main

import "fmt"

/*
	指针
*/
//值传递
func swap(c, d int) {
	fmt.Printf("传入后c %p d %p \n", &c, &d) //传入后a 0xc000062098 b 0xc0000620b0
	c, d =  d,c
	fmt.Printf("交换后c %p d %p \n", &c, &d) //交换后a 0xc000062098 b 0xc0000620b0
	fmt.Println("交换后值不影响a,b", c, d)              //交换后 4 3 地址未改变只是值改变了 也就是说变量a,b分配的内存空间没有改变，只是在这个内存空间内存储的值改变了
}

//引用传递
func swapPointe(c, d *int) {
	fmt.Printf("传入后c %p d %p \n", c, d)
	*c, *d = *d, *c
	fmt.Println("引用交换后c,d", *c, *d)
	fmt.Println("引用交换后c,d", c, d)
	/*
		引用交换后a 0xc00006c080 b 0xc00006c088 初始地址未改变，但是地址指向的值发生交换
		也就是a地址原本是0xc00006c080 ，b地址原本是0xc00006c088
	*/

}
func main() {
	a, b := 3, 4
	//初始a,b地址 a 0xc000062080 b 0xc000062088
	fmt.Printf("初始a %p b %p \n", &a, &b)
	swap(a, b)
	fmt.Println(a, b) //3,4值传递 函数内部a,b改变了 外部a，b 不会改变
	fmt.Printf("初始a %p b %p \n", &a, &b)
	swapPointe(&a, &b)//传址 a 0xc000062080 b 0xc000062088
	fmt.Println(a, b) //4,3 拷贝的地址
	fmt.Println(&a, &b) //4,3 拷贝的地址
}

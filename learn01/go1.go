package main

//

import "time"
import "fmt"

/*
	import {
		"time"
		"fmt"
	}
*/
func main() { //左{ 一定在方法名称旁边 不能换行
	fmt.Println("hello word")
	time.Sleep(1)

	// =====================================变量神明===============================================
	/*
		var a int                        //1默认值0
		var a1 int = 1111111111111111100 //2初始化值
		var c = "120"                    //3可以省去数据类型，自动匹配
		e := 12.3                        //4省去var关键字 自动匹配  只能用在函数体内申明
		var ccd string = "99999"

		var xx, yy, zz int = 11, 22, 33 //多个变量声明  按顺序赋值
		fmt.Println("xx=", xx, "yy=", yy, "zz=", zz)
		var kk, ll = "ppp", 22
		fmt.Println("kk=", kk, "ll=", ll)

		var ( //多行变量声明
			cc string = "222"
			dd bool   = false
		)
		fmt.Println("cc=", cc, "dd=", dd)

		fmt.Println("a->", a, 1)
		// 格式化输出
		fmt.Printf("a1->%s,type of a1 =%T\n", a1, a1)
		fmt.Printf("c->%s,type of c =%T\n", c, c)
		fmt.Printf("ccd->%s,type of ccd =%T\n", ccd, ccd)
		fmt.Printf("e->%s,type of e =%T\n", e, e)
	*/
	//==================================常量=========================================
	/*
		//用来定义枚举
		const (
			//可以在const添加一个关键字iota,每行iota会累加1， 第一行默认为0

			BEIJIN   = iota //iota=0
			SHNAGHAI        //iota=1
			NANJING         //iota=2
		)
		fmt.Println("BEIJIN=", BEIJIN)
		fmt.Println("SHNAGHAI=", SHNAGHAI)
		fmt.Println("NANJING=", NANJING)
		const (// 多个常量 用（）
			a,b=iota+1,iota+2  //iota=0,a=0+1,b=0+2
			c,d					//iota=1,c=1+1,d=1+2
			e,f					//iota=2,e=2+1,f=2+2

			g,h=iota*2,iota*3   //iota=3,g=3*2,h=3*3
			i,j					//iota=4,i=4*2,j=4*3
		)
		const (
			k=iota   //此时 ioto=0
			l
		)
		fmt.Println("a=", a)
		fmt.Println("b=", b)
		fmt.Println("c=", c)
		fmt.Println("d=", d)
		fmt.Println("e=", e)
		fmt.Println("f=", f)
		fmt.Println("g=", g)
		fmt.Println("h=", h)
		fmt.Println("i=", i)
		fmt.Println("j=", j)
		fmt.Println("k=", k)
		fmt.Println("l=", l)

		//var length0 = 100
		const length = 100 //只读属性
		fmt.Printf("length->%s,type of length =%T\n", length, length)
		//length=30// 不允许修改

	*/
	//	=======================================函数===================================================

	c := sum(111, 222)
	fmt.Println("sum:", c)

	d, e := fool(111, 222)
	fmt.Println("fool 1:", d)
	fmt.Println("fool 2:", e)

	j, h := fool2(111, 222)
	fmt.Println("fool2 1:", j)
	fmt.Println("fool2 2:", h)

	i, g := fool3(111, 222)
	fmt.Println("fool3 1:", i)
	fmt.Println("fool3 2:", g)

}

// int 返回值匿名 没有不加
func sum(a int, b int) int {
	fmt.Println("a->", a, " b->", b)
	c := a + b
	return c
}

// 多返回值匿名
func fool(a int, b int) (int, int) {
	fmt.Println("a->", a, " b->", b)
	c := a + b
	return c, 222
}

// 多返回值 有名称的返回值
func fool2(a int, b int) (a1 int, a2 int) {
	fmt.Println("a->", a, " b->", b)

	a1 = 1000
	a2 = 2000
	return
	//way2
	//c := a + b
	//return c,222
}

// 多返回值 返回类型如果一样，可以简写
func fool3(a int, b int) (a1, a2 int) {
	fmt.Println("a->", a, " b->", b)
	c := a + b
	return c, 222
}

//

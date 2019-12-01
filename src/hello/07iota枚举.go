package main

import "fmt"

func main() {
	//1. iota常量自动生成器,每隔一行,自动累加1
	//2. iota给常量赋值使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	//3. iota遇到const,重置为0
	const d = iota
	fmt.Println(d)

	//4. 可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)

	//5. 如果是同一行,值都一样

	const (
		i    = iota
		j, k = iota, iota
		x    = iota
		y    = iota
	)
	fmt.Println(i, j, k, x, y)

	//6. 每一行的常量个数要相同的情况下,才可以省略掉iota赋值
	const (
		p, q = iota, iota
		r, s
		u, v
	)
	fmt.Println(p, q, r, s, u, v)

	const (
		h, g = iota, iota
		m, n
		w = iota
		o = iota
		l = iota
	)
	fmt.Println(h, g, m, n, w, o, l)
}

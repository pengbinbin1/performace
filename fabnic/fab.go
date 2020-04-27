package fabnic

func Fab1(n int) int {
	if n <= 2 {
		return 1
	}

	return Fab1(n-1) + Fab1(n-2)
}

func Fab2(n int) int {
	if n <= 2 {
		return 1
	}

	fa1 := 1
	fa2 := 1
	var res int
	for i := 3; i <= n; i++ {
		res = fa1 + fa2
		fa1 = fa2
		fa2 = res
	}
	return res
}

/*
func main() {
	log.Println("Enter main thread")

	t := time.Now()
		res := fab1(45)
		log.Println("fab1 result:", res)
		log.Println("fab1 cost time:", time.Since(t))

	log.Println("call fab2")
	t2 := time.Now()
	res2 := fab2(70)
	log.Println("fab2 result:", res2)
	log.Println("fab2 cost time:", time.Since(t2))
}
*/

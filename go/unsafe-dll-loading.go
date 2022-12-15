package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
)

func main() {
	// ruleid: unsafe-dll-loading
	test1 := syscall.NewLazyDLL("a_insecure.dll")
	fmt.Println(test1.Load())

	// ruleid: unsafe-dll-loading
	syscall.LoadDLL("b_insecure.dll")


	// ok: unsafe-dll-loading
	test3 := windows.NewLazySystemDLL("c_secure.dll")
	test3.Load()

	// ruleid: unsafe-dll-loading
	test4 := windows.NewLazyDLL("d_insecure.dll")
	fmt.Println(test4.Load())

	// ok: unsafe-dll-loading
	test5 := windows.NewLazyDLL("e_secure.dll")
	test5.System = true
	fmt.Println(test5.Load())

	// ruleid: unsafe-dll-loading
	test6 := windows.NewLazyDLL("f_insecure.dll")
	fmt.Println(test6.Load())
	test6.System = true

	// ruleid: unsafe-dll-loading
	test6 = windows.NewLazyDLL("g_insecure.dll")
	test6.Load()
	test6.System = true

	// ok: unsafe-dll-loading
	test1337 := windows.NewLazyDLL("h_secure.dll")
	test1337.System = true
	test1337.Load()

	// ok: unsafe-dll-loading
	test1338 := windows.NewLazyDLL("i_secure.dll")
	test1338.System = true

	// ruleid: unsafe-dll-loading
	test7 := windows.NewLazyDLL("j_insecure.dll")
	fmt.Println(test7.Load())
	test7.System = true

	// ok: unsafe-dll-loading
	fmt.Println(windows.NewLazyDLL("k_secure.dll").Name)

	// ruleid: unsafe-dll-loading
	b := windows.NewLazyDLL("l_insecure.dll")

	// ruleid: unsafe-dll-loading
	b = windows.NewLazyDLL("m_insecure.dll")
	b.Load()

	// ok: unsafe-dll-loading
	test8 := windows.NewLazyDLL("n_secure.dll")
	fmt.Println("Hello World!")
	test8.System = true
	fmt.Println(test8.Load())

	// ruleid: unsafe-dll-loading
	test9 := windows.NewLazyDLL("o_insecure.dll")
	fmt.Println("Hello World!")
	fmt.Println(test9.Load())

	// ruleid: unsafe-dll-loading
	test10 := windows.NewLazyDLL("p_insecure.dll")
	fmt.Println("Hello World!")
	fmt.Println(test10.Load())
	test10.System = true

	// ruleid: unsafe-dll-loading
	test0, _ := syscall.LoadLibrary("q_insecure.dll")
	fmt.Println(test0)

	// ruleid: unsafe-dll-loading
	windows.LoadDLL("r_insecure.dll")

	// ruleid: unsafe-dll-loading
	windows.MustLoadDLL("s_insecure.dll")
}
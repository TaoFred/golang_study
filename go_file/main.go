package main

import "fmt"

func main() {

	// r := gin.Default()

	// r.GET("/AI711-H11使用手册.pdf", func(c *gin.Context) {

	// 	filePath := "D:/IntegrityData_test/DocumentData/default-bucket/default/ECS-700/Chinese/03系统硬件/IO模块使用手册/AI711-H11使用手册.pdf"
	// 	c.File(filePath)
	// })
	// r.Run(":8080")
	a := "fefe%"
	fmt.Printf("a: %v\n", a)

	a = "fefe%%"
	fmt.Printf("a: %v\n", a)

	a = "fefe" + "%"
	fmt.Printf("a: %v\n", a)

	a = "fefe" + "%d"
	fmt.Printf("a: %v\n", a)
}

/*
 * @Author: GG
 * @Date: 2022-09-15 10:58:45
 * @LastEditTime: 2022-09-15 11:03:23
 * @LastEditors: GG
 * @Description:
 * @FilePath: \gin-blog\demo.go
 *
 */
package main

import (
	"fmt"
	"gin-blog/models"
)

func main1() {
	postList := make([]models.Post, 10)
	fmt.Printf("postList: %v\n", postList)
	fmt.Printf("postList: %T\n", postList)
	fmt.Printf("len(postList): %v\n", len(postList))

	postList1 := []models.Post{}
	fmt.Printf("postList1: %v\n", postList1)
	fmt.Printf("postList1: %T\n", postList1)
	fmt.Printf("len(postList): %v\n", len(postList1))

	var postList2 []models.Post
	fmt.Printf("postList2: %v\n", postList2)
	fmt.Printf("postList2: %T\n", postList2)
	fmt.Printf("len(postList): %v\n", len(postList2))
}

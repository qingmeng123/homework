/*******
* @Author:qingmeng
* @Description:排序
* @File:sort
* @Date2021/10/19
 */

package main

import (
	"fmt"
	"math/rand"
)

func my_sort(enum[] int)  {
	for i:=0;i< len(enum)-1;i++{
		for j:=0;j< len(enum)-i-1;j++{
			if enum[j]<enum[j+1]{
				enum[j],enum[j+1]=enum[j+1],enum[j]
			}
		}
	}
}
func main() {
	var s[] int
	for i:=0;i<100;i++{
		s=append(s,rand.Intn(150))
	}
	my_sort(s)
	fmt.Println(s)
}

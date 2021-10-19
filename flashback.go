/*******
* @Author:qingmeng
* @Description:字符串倒叙返回字符串
* @File:flashback
* @Date2021/10/18
 */

package main

import "fmt"

func back(s string) string {
	enum:=[]byte(s)
	var (
		res string
		data []byte
	)
	n:= len(enum)
	for i:=0;i<n;i++{
		data=append(data,enum[n-i-1])
	}
	res=string(data)
	return res
}
func main() {
	var s string
	fmt.Println("请输入一串一串字符串：")
	fmt.Scan(&s)
	s=back(s)
	fmt.Println("倒叙后：",s)

}

/*******
* @Author:qingmeng
* @Description:基本运算
* @File:counter
* @Date2021/10/18
 */

package main

import "fmt"

func count(a float64,cmd string,b float64 )(float64,bool) {
	res:=0.0
	flag:=true
	switch cmd {
	case"+":
		res=a+b
	case "-":
		res=a-b
	case "*":
		res=a*b
	case "/":
		if b!=0{
			res=a/b
		}else{
			flag=false
			fmt.Println("错误，除数为0")
		}
	default:
		flag=false
		fmt.Println("运算符输入错误")
	}
	return res,flag
}
func main() {
	var(
		a,b,res float64
		cmd string
		flag bool
	)
	fmt.Println("请输入需要计算的内容（a cmd b) a为0时退出程序:")

	for i:=0;;i++{
		fmt.Scan(&a)
		switch  {
		case a!=0:
			fmt.Scan(&cmd,&b)
			res,flag=count(a,cmd,b)
			if flag{
				fmt.Println(a,cmd,b,"=",res)
			}
		default:
			fmt.Println("计算结束")
			break
		}
	}



}
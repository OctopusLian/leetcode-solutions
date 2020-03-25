package main

import (
	"fmt"
	//"strconv"
)

func main() {
	//1,输入
	fmt.Println("pls input 3 numbers a,b,c (0<a,b,c<10 and a!=b!=c):")
	var x1, x2, x3 int
	fmt.Scanf(":%d %d %d", &x1, &x2, &x3)
	//var str1, str2, str3 string

	//2,转为字符串
	//x1_str := strconv.Itoa(x1)
	//x2_str := strconv.Itoa(x2)
	//x3_str := strconv.Itoa(x3)

	//3,判断和做任务
	flag1 := 0
	flag2 := 0
	flag3 := 0

	for i := 1; i <= 100; i++ {
		var num string
		// i_str := strconv.Itoa(i)

		/*
					if(num.find(str1)<3)
			        {
			            //cout<<num<<endl;
			            cout<<"Fizz"<<endl;
			            continue;
			        }

		*/

		if i%1 == 0 {
			flag1 = 1
		}
		if i%2 == 0 {
			flag2 = 2
		}
		if i%3 == 0 {
			flag3 = 4
		}

		//cpp位运算
		//int res=flag1|flag2|flag3;//标识位
		res := flag1 | flag2 | flag3
		flag1 = 0
		flag2 = 0
		flag3 = 0

		switch res {
		case 1:
			fmt.Println("Fizz")
			break //不要忘了break
		case 2:
			fmt.Println("Buzz")
			break
		case 4:
			fmt.Println("Whizz")
			break
		case 3:
			fmt.Println("FizzBuzz")
			break
		case 5:
			fmt.Println("FizzWhizz")
			break
		case 6:
			fmt.Println("BuzzWhizz")
			break
		case 7:
			fmt.Println("FizzBuzzWhizz")
			break
		default:
			fmt.Println(num)
			break
		}

	}

	//return 0
}
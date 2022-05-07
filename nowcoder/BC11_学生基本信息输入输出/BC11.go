package main

import (
	"fmt"
)

func main() {
	var sno int                                                                                                     //学号
	var c_score, math_score, english_score float32                                                                  //三科成绩,注意有小数
	fmt.Scanf("%d;%f,%f,%f", &sno, &c_score, &math_score, &english_score)                                           //输入学号和成绩
	fmt.Printf("The each subject score of  No. %d is %.2f, %.2f, %.2f.\n", sno, c_score, math_score, english_score) //输出
}

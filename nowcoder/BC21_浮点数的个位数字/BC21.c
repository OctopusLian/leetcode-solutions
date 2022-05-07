#include <stdio.h>
int main(){
    float n;
    int i;
    scanf("%f", &n);  //输入浮点数
    i = (int)n;  //只取整数部分
    i = i % 10;  //取10的余数
    printf("%d", i);
}
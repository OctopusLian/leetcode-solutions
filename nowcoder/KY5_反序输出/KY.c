#include<stdio.h>
int main() {
    char a[4];  //声明一个长度为4的数组
    while(scanf("%s",a)!=EOF){
        for(int i=3;i>=0;i--)
        printf("%c",a[i]); //倒序输出
    }
}
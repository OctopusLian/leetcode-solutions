#include<stdio.h>
int main() {
    int a;
    int res;
    scanf("%d",&a);
    for (;a > 0;){
        res = a %10;
        a = a / 10;
        printf("%d",res); //不要换行
    }
}
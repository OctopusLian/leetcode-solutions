#include<stdio.h>
int main()
{
    int i,a;
    scanf("%d",&a);  //输入
    for(i=0;a!=1;i++)
        if(a%2==0)
            a/=2;  //偶数
        else
            a=(3*a+1)/2;  //奇数
        printf("%d\n",i);
        return 0;
}
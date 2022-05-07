#include<stdio.h>
int common(int a,int b);
int main()
{
    int x,y;
    while(scanf("%d%d",&x,&y)!=EOF)  //输入包含多组数据，每组数据包含两个正整数x和y
    {
        printf("%d\n",common(x,y));  //递归解法
    }
    return 0;
}
int common(int a,int b)
{
    //二分法
    if(a==b)
    {
        return a;  //公共父节点
    }
    else if(a>b){
        return common(a/2,b);
    }
    else{
        return common(a,b/2);
    }
}
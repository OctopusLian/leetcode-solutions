#include<stdio.h>
#include<math.h>
#pragma warning( disable : 4996)
 
char skew[1000] = { 0 };
int i = 0,b;
int num = 0;//num为10进制后
int main()
{
     
    while (scanf("%s", &skew) != EOF)
    {
        i = 0;
        num = 0;//每循环一次重置
        while (skew[i] != '\0')//i用来计算数组长度
        {
            i++;
        }
        b = i;
        for (int k = 0; k < i; k++)
        {
            num += (skew[k] - 48) * (pow(2, b) - 1);
            b--;
        }
        printf("%d\n", num);
 
    }
     
    return 0;
}
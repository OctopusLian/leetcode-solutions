#include<stdio.h>
int main()
{
    int n;
    while(scanf("%d", &n) != EOF)
    {
        if(n >= 90 && n <= 100)
            printf("Perfect\n");
    }
}
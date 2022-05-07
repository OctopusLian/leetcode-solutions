#include<stdio.h>
int main()
{
    char c, x;
    while(scanf("%c", &c) == 1)
    {
        if(c >= 'a' && c <= 'z')
        {
            x = c - 32;
            printf("%c\n",x);
        }
        if(c >= 'A' && c <= 'Z')
        {
            x = c + 32;
            printf("%c\n",x);
        }
          
    }
    return 0;
}
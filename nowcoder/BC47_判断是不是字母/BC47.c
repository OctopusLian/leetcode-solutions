#include<stdio.h>
int main()
{
    char c;
    while(scanf("%c\n", &c) == 1)
    {
        if((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'))
        {
            printf("%c is an alphabet.\n", c);
        }
        else
        {
            printf("%c is not an alphabet.\n", c);
        }
              
    }
    return 0;
 }

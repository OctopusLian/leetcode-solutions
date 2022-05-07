#include <stdio.h>
  
int main()
{
    int score;
    while(scanf("%d", &score) != EOF)
    {
        if(score >= 60)
            printf("Pass\n");
        else
            printf("Fail\n");
    }
}
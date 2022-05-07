#include <stdio.h>
#include <stdlib.h>
#define N  25
int v[N];
int f[N];
int main()
{
    int n;
    while(scanf("%d",&n)!=EOF)
    {
        for(int i=1;i<=n;i++)
        {
            scanf("%d",&v[i]);
        }
        f[0]=1;
        for(int i=1;i<=n;i++)
        {
            for(int j=40;j>=v[i];j--)
            {
                f[j]=f[j]+f[j-v[i]];
            }
        }
        printf("%d\n",f[40]);
    }
 
    return 0;
}
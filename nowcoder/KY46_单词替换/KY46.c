#include <stdio.h>
#include <string.h>
int main()
{
    char s[10000],a[101],b[101],c[1000][100];
    int i,j,k;
    while(gets(s))
    {
        gets(a),gets(b);
        for(i=0,j=0,k=0;i<strlen(s);i++)
        {
            if(s[i]!=' ')    c[j][k++]=s[i];
            else             c[j][k+1]='\0',j++,k=0;
        }
        for(i=0,j++;i<j;i++)
        {
            if(strcmp(c[i],a)==0)
            {
                if(i==0)  printf("%s",b);
                else printf(" %s",b);
            }
            else
            {
                if(i==0) printf("%s",c[i]);
                else printf(" %s",c[i]);
            }
        }
        printf("\n");
    }
    return 0;
}
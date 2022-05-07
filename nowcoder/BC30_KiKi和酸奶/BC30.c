#include<stdio.h>
int main() {
    int n,h,m;
    scanf("%d %d %d",&n,&h,&m);
    if (m % h == 0) {
        printf("%d\n",n - m / h);
    } else {
        printf("%d\n",n - m / h - 1);
    }
    
    return 0;
}
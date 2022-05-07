#include<stdio.h>
int main(){
    int n;
    while(scanf("%d\n",&n) != EOF) {
        printf("%d\n",1 << n);
    }
    
    return 0;
}
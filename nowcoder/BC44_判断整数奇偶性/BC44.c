#include<stdio.h>
int main(){
    int num;
    while(scanf("%d\n",&num)!=EOF){
        if (num%2)
            printf("Odd\n");
        else
            printf("Even\n");
    }
    return 0;
}
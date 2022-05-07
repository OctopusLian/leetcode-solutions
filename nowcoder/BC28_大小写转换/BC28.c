#include<stdio.h>
int main(){
    char ch;
    while(scanf("%c",&ch) != EOF){
        printf("%c\n",ch+32);  //大写转小写
        getchar();
    }
}
#include <stdio.h>
#include <ctype.h>
  
int main(){
    char ch;
    while(scanf("%c",&ch)!=EOF){
        if(isalpha(ch)){
            printf("YES\n");
        } else {
            printf("NO\n");
        }
        getchar();
    }
}
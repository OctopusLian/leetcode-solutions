#include <stdio.h>
int main(){
    int seconds;
    scanf("%d", &seconds);
    printf("%d %d %d", seconds/3600, seconds%3600/60, seconds%60);
     
    return 0;
}
#include <stdio.h>
int main(){
    int N, min = 1000000, max = -1000000, temp;
    while (scanf("%d", &N) != EOF){
        for (int i = 0; i < N; ++i){
            scanf("%d", &temp);
            max = (temp > max) ? temp : max;
            min = (temp < min) ? temp : min;
        }
        printf("%d %d\n", max, min);
    }
    return 0;
}
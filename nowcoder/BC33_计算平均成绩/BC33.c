#include<stdio.h>
int main()
{
    int a, b, c, d, e;
    double sum, avg;
    scanf("%d%d%d%d%d", &a, &b, &c, &d, &e);
    sum = a + b + c + d + e;
    avg = sum / 5;
    printf("%.1f", avg);
}
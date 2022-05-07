#include<stdio.h>
#define pi 3.1415927
int main(){
    float r,v;
    scanf("%f",&r);
    v = (4 * pi * r * r * r) / 3;
    printf("%.3f",v);
}
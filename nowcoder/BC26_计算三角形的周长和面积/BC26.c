#include<stdio.h>
int main() {
    float a,b,c;
    float cir,area;
    scanf("%f %f %f",&a,&b,&c);
    cir = a + b + c;
    /*海伦公式p=(a+b+c)/2,s=sqrt(p(p-a)(p-b)(p-c));*/
    float p=(a+b+c)/2.00;
    area = sqrt(p*(p-a)*(p-b)*(p-c));
    printf("circumference=%.2f area=%.2f",cir,area);
}
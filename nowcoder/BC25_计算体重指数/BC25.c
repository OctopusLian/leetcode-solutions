#include <stdio.h>
int main()
{
    float a,b,BMI;
    scanf("%f %f",&a,&b);
    b=b/100;
    printf("%.2f",a/(b*b));
    return 0;
}
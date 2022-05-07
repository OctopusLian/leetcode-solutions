#include<stdio.h>
int main(){
    float height,weight,bmi;
    scanf("%f%f",&height,&weight);
    bmi=height/(weight*weight);
    if(bmi>18.5&&bmi<23.9){
        printf("Normal");
    }else{
        printf("Abnormal");
    }
      
    return 0;
}
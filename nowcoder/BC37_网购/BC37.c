#include <stdio.h>
  
int main()
{
    float money;
    int month,day,zk;
    scanf("%f %d %d %d",&money,&month,&day,&zk);
    float sm;
    if(month==11&&day==11&&zk==1){
        sm=money*0.7-50.0;
    }else if(month==11&&day==11&&zk==0){
         sm=money*0.7;
    }else if(month==12&&day==12&&zk==1){
         sm=money*0.8-50.0;
    }else if(month==12&&day==12&&zk==0){
         sm=money*0.8;
    }else {
        sm=money;
    }
    if(sm<0){
        sm=0;
    }
    printf("%.2f\n",sm);
}

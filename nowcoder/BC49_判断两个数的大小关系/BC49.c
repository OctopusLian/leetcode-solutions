#include<stdio.h>
int main(){
    int a,b;
    while (scanf("%d %d",&a,&b)!=EOF){
        if(a==b){
        printf ("%d=%d\n",a,b); }
       else if (a>b){
           printf("%d>%d\n",a,b);}
            else if (a<b){
           printf("%d<%d\n",a,b);}
            
    }
  
    return 0;
}

#include<stdio.h>
int main(){
    int n,a[100]={0},i,j,t;
    scanf("%d",&n);
    for(i=0;i<n;i++){
        scanf("%d",&a[i]);
    }
    for(i=0;i<n;i++){
        for(j=0;j<n;j++){
            if(a[j]<a[j+1]){
                t=a[j];
                a[j]=a[j+1];
                a[j+1]=t;
            }
        }
    }
    for(i=0;i<5;i++){
        printf("%d ",a[i]);
    }
}

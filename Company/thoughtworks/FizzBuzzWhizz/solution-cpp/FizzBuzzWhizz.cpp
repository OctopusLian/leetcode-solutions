#include <iostream>
#include <sstream>
using namespace std;

int main()
{
    cout<<"pls input 3 numbers a,b,c (0<a,b,c<10 and a!=b!=c):"<<endl;
    int x1,x2,x3;
    cin>>x1>>x2>>x3;

    string str1,str2,str3;
    ostringstream convert;
    //转为字符串
    convert<<x1;
    str1=convert.str();
    convert.str("");//清空convert
    convert<<x2;
    str2=convert.str();
    convert.str("");//清空convert
    convert<<x3;
    str3=convert.str();

    //cout<<str1<<str2<<str3<<endl;
    //cout<<x1<<x2<<x3<<endl;

    int flag1=0;
    int flag2=0;
    int flag3=0;

    for(int i=1; i<=100; i++)
    {
        //转为字符串
        string num;
        ostringstream convert;
        convert<<i;
        num=convert.str();
        //cout<<num<<endl;

        if(num.find(str1)<3)
        {
            //cout<<num<<endl;
            cout<<"Fizz"<<endl;
            continue;
        }
        if(i%x1==0)
        {
            flag1=1;
        }
        if(i%x2==0)
        {
            flag2=2;
        }
        if(i%x3==0)
        {
            flag3=4;
        }
        int res=flag1|flag2|flag3;//标识位
        flag1=0;
        flag2=0;
        flag3=0;
        //cout<<"res : "<<res<<endl;
        switch(res)
        {
        case 1:
            cout<<"Fizz"<<endl;
            break;//不要忘了break
        case 2:
            cout<<"Buzz"<<endl;
            break;
        case 4:
            cout<<"Whizz"<<endl;
            break;
        case 3:
            cout<<"FizzBuzz"<<endl;
            break;
        case 5:
            cout<<"FizzWhizz"<<endl;
            break;
        case 6:
            cout<<"BuzzWhizz"<<endl;
            break;
        case 7:
            cout<<"FizzBuzzWhizz"<<endl;
            break;
        default:
            cout<<num<<endl;
            break;
        }

    }

    return 0;

}
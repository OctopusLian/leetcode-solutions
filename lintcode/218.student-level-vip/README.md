## 描述  

```
实现一个学生的类 Student. 包含如下的成员函数和方法：

    两个公有成员 name 和 score。分别是字符串类型和整数类型。
    一个构造函数，接受一个参数 name。
    一个公有成员函数 getLevel() 返回学生的成绩等级（一个字符）。

等级表如下：

    A: score >= 90
    B: score >= 80 and < 90
    C: score >= 60 and < 80
    D: score < 60

您在真实的面试中是否遇到过这个题？  
样例

样例 1：

Java:
    Student student = new Student("Zuck");
    student.score = 10;
    student.getLevel(); // should be 'D'
    student.score = 60;
    student.getLevel(); // should be 'C'

样例 2：

Python:
    student = Student("Zuck")
    student.score = 10
    student.getLevel() # should be 'D'
    student.score = 60
    student.getLevel() # should be 'C'


```
# Conference TrackManagement

```
You are planning a big programming conference and have receivedmany proposals which have passed the initial screen process but you're havingtrouble fitting them into the time constraints of the day -- there are so manypossibilities! So you write a program to do it for you.

    The conference has multiple tracks each of which has a morning and afternoon session.
    Each session contains multiple talks.
    Morning sessions begin at 9am and must finish before 12 noon, for lunch.
    Afternoon sessions begin at 1pm and must finish in time for the networking event.
    The networking event can start no earlier than 4:00 and no later than 5:00.
    No talk title has numbers in it.
    All talk lengths are either in minutes (not hours) or lightning (5 minutes).
    Presenters will be very punctual; there needs to be no gap between sessions.


Note that depending on how you choose to complete this problem,your solution may give a different ordering or combination of talks intotracks. This is acceptable; you don’t need to exactly duplicate the sampleoutput given here.


Test input:

Writing Fast Tests Against Enterprise Rails 60min

Overdoing it in Python 45min

Lua for the Masses 30min

Ruby Errors from Mismatched Gem Versions 45min

Common Ruby Errors 45min

Rails for Python Developers lightning

Communicating Over Distance 60min

Accounting-Driven Development 45min

Woah 30min

Sit Down and Write 30min

Pair Programming vs Noise 45min

Rails Magic 60min

Ruby on Rails: Why We Should Move On 60min

Clojure Ate Scala (on my project) 45min

Programming in the Boondocks of Seattle 30min

Ruby vs. Clojure for Back-End Development 30min

Ruby on Rails Legacy App Maintenance 60min

A World Without HackerNews 30min

User Interface CSS in Rails Apps 30min

 

Test output: 

Track 1:

09:00AM Writing Fast Tests Against Enterprise Rails 60min

10:00AM Overdoing it in Python 45min

10:45AM Lua for the Masses 30min

11:15AM Ruby Errors from Mismatched Gem Versions 45min

12:00PM Lunch

01:00PM Ruby on Rails: Why We Should Move On 60min

02:00PM Common Ruby Errors 45min

02:45PM Pair Programming vs Noise 45min

03:30PM Programming in the Boondocks of Seattle 30min

04:00PM Ruby vs. Clojure for Back-End Development 30min

04:30PM User Interface CSS in Rails Apps 30min

05:00PM Networking Event

 

Track 2:

09:00AM Communicating Over Distance 60min

10:00AM Rails Magic 60min

11:00AM Woah 30min

11:30AM Sit Down and Write 30min

12:00PM Lunch

01:00PM Accounting-Driven Development 45min

01:45PM Clojure Ate Scala (on my project) 45min

02:30PM A World Without HackerNews 30min

03:00PM Ruby on Rails Legacy App Maintenance 60min

04:00PM Rails for Python Developers lightning

05:00PM Networking Event
```

## 中文翻译  

```
您正在计划召开一次大型编程会议，并且收到了许多提案，这些提案已经通过了最初的筛选过程，但是您很难将它们适应一天中的时间限制-如此之多！因此，您编写了一个程序来为您做。

    会议有多个轨道，每个轨道都有一个上午和下午的会议。
    每个会话包含多个对话。
    上午会议从上午9点开始，必须在中午12点之前结束，以供午餐。
    下午的会议从下午1点开始，必须在网络活动之前及时结束。
    联网活动的开始时间不得早于4:00，且不得晚于5:00。
    谈话标题中没有数字。
    所有通话时间以分钟（而非小时）或闪电（5分钟）为单位。
    演讲者将非常准时。会议之间应该没有间隙。

请注意，根据您选择解决此问题的方式，您的解决方案可能会将谈话的顺序或轨迹组合不同。这是可以接受的；您无需完全重复此处给出的样本输出。


测试输入：

编写针对Enterprise Rails的快速测试60分钟

在Python中过度处理45分钟

大众欢呼30分钟

宝石版本不匹配导致的Ruby错误45分钟

常见的Ruby错误45分钟

适用于Python开发人员的Rails闪电

远距离通讯60分钟

会计驱动开发45分钟

哇30分钟

坐下来写30分钟

结对编程与噪声45分钟

Rails魔术60分钟

Ruby on Rails：为什么我们应该继续前进60分钟

Clojure Ate Scala（在我的项目中）45分钟

在西雅图Boockcks编程30分钟

Ruby vs. Clojure用于后端开发30分钟

Ruby on Rails旧版应用维护60分钟

没有黑客新闻的世界30分钟

Rails Apps中的用户界面CSS 30分钟

 

测试输出：

音轨1：

09:00 AM针对Enterprise Rails编写快速测试60分钟

10:00 AM在Python中处理它45分钟

10:45 AM Lua为大众服务30分钟

11:15 AM宝石版本不匹配导致的Ruby错误45分钟

12:00 PM午餐

01:00 PM Ruby on Rails：为什么我们应该继续前进60分钟

02:00 PM常见的Ruby错误45分钟

02:45 PM配对编程与噪音45分钟

03:30 PM在西雅图Boockcks编程30分钟

04:00 PM Ruby与Clojure进行后端开发30分钟

04:30 PM Rails Apps中的用户界面CSS 30分钟

05:00 PM社交活动

 

音轨2：

09:00 AM全程沟通60分钟

10:00 AM Rails Magic 60分钟

11:00 AM哇30分钟

11:30 AM坐下来写30分钟

12:00 PM午餐

01:00 PM会计驱动开发45分钟

01:45 PM Clojure Ate Scala（在我的项目中）45分钟

02:30 PM没有HackerNews的世界30分钟

03:00 PM Ruby on Rails旧版应用维护60分钟

04:00 PM适用于Python开发人员的Rails闪电

05:00 PM社交活动
```
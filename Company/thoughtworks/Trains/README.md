# Problem one: Trains

```
The local commuter railroad services a number of towns inKiwiland.  Because of monetary concerns, all of the tracks are 'one-way.' That is, a route from Kaitaia to Invercargill does not imply theexistence of a route from Invercargill to Kaitaia.  In fact, even if bothof these routes do happen to exist, they are distinct and are not necessarilythe same distance!

The purpose of this problem is to help the railroad provide itscustomers with information about the routes.  In particular, you willcompute the distance along a certain route, the number of different routesbetween two towns, and the shortest route between two towns.

Input: A directed graph where a node represents a town and an edge represents aroute between two towns.  The weighting of the edge represents thedistance between the two towns.  A given route will never appear more thanonce, and for a given route, the starting and ending town will not be the sametown.

Output:For test input 1 through 5, if no such route exists, output 'NO SUCH ROUTE'. Otherwise, follow the route as given; do not make any extra stops! For example, the first problem means to start at city A, then traveldirectly to city B (a distance of 5), then directly to city C (a distance of4).

1.    The distance of the route A-B-C.

2.    The distance of the route A-D.

3.    The distance of the route A-D-C.

4.    The distance of the routeA-E-B-C-D.

5.    The distance of the route A-E-D.

6.    The number of trips starting at Cand ending at C with a maximum of 3 stops.  In the sample data below,there are two such trips: C-D-C (2 stops). and C-E-B-C (3 stops).

7.    The number of trips starting at Aand ending at C with exactly 4 stops.  In the sample data below, there arethree such trips: A to C (via B,C,D); A to C (via D,C,D); and A to C (viaD,E,B).

8.    The length of the shortest route(in terms of distance to travel) from A to C.

9.    The length of the shortest route(in terms of distance to travel) from B to B.

10. The number of different routesfrom C to C with a distance of less than 30.  In the sample data, thetrips are: CDC, CEBC, CEBCDC, CDCEBC, CDEBC, CEBCEBC, CEBCEBCEBC.

 

Test Input:

For the test input, the towns are named using the first fewletters of the alphabet from A to D.  A route between two towns (A to B)with a distance of 5 is represented as AB5.

Graph: AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7

Expected Output:

Output #1: 9

Output #2: 5

Output #3: 13

Output #4: 22

Output #5: NO SUCH ROUTE

Output #6: 2

Output #7: 3

Output #8: 9

Output #9: 9

Output #10: 7
```

## 中文翻译  

```
当地的通勤铁路为基维兰的许多城镇提供服务。由于金钱的考虑，所有途径都是“单向的”。也就是说，从凯塔亚到因弗卡吉尔的路线并不意味着存在从因弗卡吉尔到凯塔亚的路线。实际上，即使这两个路径确实都存在，它们也是截然不同的，并且不一定是相同的距离！

该问题的目的是帮助铁路为其客户提供有关路线的信息。特别是，您将计算沿特定路线的距离，两个镇之间的不同路线的数量以及两个镇之间的最短路线。

输入：有向图，其中一个节点代表一个镇，一条边代表两个镇之间的路线。边缘的权重表示两个镇之间的距离。给定的路线永远不会出现超过一次，并且对于给定的路线，起点和终点的城镇将不是同一镇。

输出：对于测试输入1至5，如果不存在这样的路由，则输出'NO SUCH ROUTE'。否则，请遵循给定的路线；不要停下来！例如，第一个问题意味着从城市A开始，然后直接前往城市B（距离为5），然后直接前往城市C（距离为4）。

1.路线A-B-C的距离。

2.路线A-D的距离。

3.路线A-D-C的距离。

4.路线A-E-B-C-D的距离。

5.路线A-E-D的距离。

6.从Cand开始的旅行次数，以C结尾，最多3站。在下面的示例数据中，有两次这样的行程：C-D-C（2站）。和C-E-B-C（3站）。

7.从Aand到C的旅程，恰好有4个站点。在下面的示例数据中，有三个行程：A到C（通过B，C，D）； A至C（通过D，C，D）；和A到C（通过D，E，B）。

8.从A到C的最短路线的长度（以行进距离为单位）。

9.从B到B的最短路线的长度（就行进距离而言）。

10.从C到C的距离小于30的不同路线数。在样本数据中，行程为：CDC，CEBC，CEBCDC，CDCEBC，CDEBC，CEBCEBC，CEBCEBCEBC。

 

测试输入：

对于测试输入，使用字母从A到D的前几个字母来命名城镇。两个城镇（A到B）之间距离为5的路线表示为AB5。

图表：AB5，BC4，CD8，DC8，DE6，AD5，CE2，EB3，AE7

预期产量：

输出＃1：9

输出＃2：5

输出＃3：13

输出＃4：22

输出＃5：没有此类路由

输出＃6：2

输出＃7：3

输出＃8：9

输出＃9：9

输出＃10：7
```
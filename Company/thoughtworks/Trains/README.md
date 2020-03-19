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
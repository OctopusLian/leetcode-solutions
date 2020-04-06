下面是我给出的答案。
先说一些思路吧，其实这道题主要用了图的相关知识。整个题有十个小问题，大概用了图的遍历，深度优先算法之类的知识。

思路大概是这个样子：

    1~5：int routeDistance(City[] route)
        传入的cities是城市数组，比如对于第1个问题A-B-C，传入的数组是[0, 1, 2]、第2个问题A-D，传入的数组是[0, 3]，以此类推。如果发现某两个城市之间不连通，就直接抛异常，表示路径不存在（NO SUCH ROUTE）。
    6~7：int numberOfTrips(City[] route)
        比如第6题，它要求从C到C，中间最多经过3个城市，也就是说要么是 C-?-C，要么是 C-?-?-C。中间的问号代表城市不确定。?可以使ANY，也可以是OPTIONAL。遇到ANY和OPTIONAL的时候做特殊处理。如果不是Any和Optional，而是一般的城市，那就直接判断两者之间是否是连通的。
    8~9：int shortestRouteDistance(City origin, City dest)
        求最短路路径。 用了DFS算法。
    10：int numberOfRoutes(City origin, City dest, int maxDistance)
        用了简单的递归算法。

至于一些细节的话，在设计类的时候，我使用了枚举类，主要是让各个城市索引的可读性更好，比如A代表A城市的索引0，比直接写0要好。final，static，枚举覆盖等等这些知识点我都有注意。

- City.java  

```java
package main;

public enum City {
	// main use of overriding enumerations
	A("A", 0),
	B("B", 1),
	C("C", 2), 
	D("D", 3),
	E("E", 4),
	// represents any of a city
	ANY("", -1),
	// represents a city or no city 
	OPTIONAL("", -2);

	private int index;
	private String name;

	private City(String name, int index) {
		this.name = name;
		this.index = index;
	}

	public String getName() {
		return name;
	}

	/**
	 * >= 0: normal cities.
	 * <  0: else status.
	 */
	public int getIndex() {
		return index;
	}
}
```

- CityNetwork.java  

```java
package main;

import static main.City.*;

public class CityNetwork {
	private int[][] network;

	public CityNetwork(int[][] network) {
		this.network = network;
	}

	/**
	 * Return total distance of the given route. Throw an Exception if no path.
	 * 
	 * @param route the given route.
	 * @return total distance of the route.
	 * @throws Exception if no path.
	 */
	public int routeDistance(City[] route) throws Exception {
		int distance = 0;

		// get the 1st city's index
		int idx1 = route[0].getIndex();
		int idx2;
		for (int i = 1; i < route.length; i++) {
			// get next city's index
			idx2 = route[i].getIndex();
			// if there is no route between 2 cities, then throw an exception
			if (network[idx1][idx2] == 0) {
				throw new Exception();
			}
			// add distance of 2 cities to total distance
			distance += network[idx1][idx2];
			// change index-1 to index-2 and go forward to next city
			idx1 = idx2;
		}

		return distance;
	}

	/**
	 * Return the number of trips of the given route.
	 * 
	 * @param route the given route.
	 * @return the number of trips of the route.
	 */
	public int numberOfTrips(City[] route) {
		return numberOfTrips(route, 0);
	}

	/**
	 * Return the distance of the shortest route between the given cities.
	 * 
	 * @param origin the origin city (start city).
	 * @param dest the destination city.
	 * @return the distance of the shortest route between origin and destination.
	 */
	public int shortestRouteDistance(City origin, City dest) {
		boolean[][] flags = new boolean[network.length][network.length];
		int shortest = shortestRouteDistance(origin.getIndex(), dest.getIndex(), flags, 0, 0, origin == dest);
		return shortest;
	}

	/**
	 * Return the number of routes between the given cities which distance is less than maxDistance.
	 * 
	 * @param origin the origin city.
	 * @param dest the destination city.
	 * @param maxDistance the max distance of the found routes.
	 * @return the number of routes.
	 */
	public int numberOfRoutes(City origin, City dest, int maxDistance) {
		int number = numberOfRoutes(origin.getIndex(), dest.getIndex(), maxDistance, 0, origin == dest);
		return number;
	}

	private int numberOfTrips(City[] route, int offset) {
		// we are at the last city on the route.
		if (offset == route.length - 1) return 1;

		// try to go forward to next city.
		City start = route[offset];
		City next = route[offset + 1];
		int number = 0;
		int idx1, idx2;
		if (start.getIndex() < 0) {
			// if start is a special city (Any or Optional), try every city on the network graph.
			for (City city : City.values()) {
				if (city.getIndex() >= 0) {
					route[offset] = city;
					number += numberOfTrips(route, offset);
				}
			}
			// Optional city can be skipped
			if (start == OPTIONAL) {
				number += numberOfTrips(route, offset + 1);
			}
			route[offset] = start;
		} else {
			// start is a normal city
			idx1 = start.getIndex();
			// check whether next is a special city or a normal city.
			if (next.getIndex() < 0) {
				for (City city : City.values()) {
					if (city.getIndex() >= 0) {
						route[offset + 1] = city;
						number += numberOfTrips(route, offset);
					}
				}
				// Optional city can be skipped
				if (next == OPTIONAL) {
					route[offset + 1] = start;
					number += numberOfTrips(route, offset + 1);
				}
				route[offset + 1] = next;
			} else {
				idx2 = next.getIndex();
				number = network[idx1][idx2] == 0 ? 0 :
						numberOfTrips(route, offset + 1);
			}
		}

		return number;
	}
	
	/**
	 * a conventional DFS algorithm
	 * 
	 * I think Dijkstra can be improved. That's is a better idea, but i didn't realize it, I'll continue to learn.
	 * 
	 * @param origin
	 * @param dest
	 * @param flags
	 * @param shortest
	 * @param distance
	 * @param sameCity
	 * @return
	 */
	private int shortestRouteDistance(int origin, int dest, boolean[][] flags, int shortest, int distance, boolean sameCity) {
		if (!sameCity && origin == dest) return distance;

		int currentDistance;
		for (int next = 0; next < network.length; next++) {
			if (!flags[origin][next] && network[origin][next] != 0) {
				flags[origin][next] = true;
				currentDistance = shortestRouteDistance(next, dest, flags, 0, distance + network[origin][next], false);
				if (currentDistance != 0 && (shortest == 0 || currentDistance < shortest)) {
					shortest = currentDistance;
				}
				flags[origin][next] = false;
			}
		}

		return shortest;
	}

	/**
	 * @param origin
	 * @param dest
	 * @param maxDistance
	 * @param distance
	 * @param sameCity
	 * @return
	 */
	private int numberOfRoutes(int origin, int dest, final int maxDistance, int distance, boolean sameCity) {
		if (distance >= maxDistance) return 0;

		int number = 0;
		if (!sameCity && origin == dest && distance < maxDistance) {
			number++;
		}
		for (int next = 0; next < network.length; next++) {
			if (network[origin][next] != 0) {
				number += numberOfRoutes(next, dest, maxDistance, distance + network[origin][next], false);
			}
		}

		return number;
	}
}
```

- HomeWork.java  

```java
package main;

import static main.City.*;

public class HomeWork {
	// all case routes
	private static final City[][] ROUTES = {
			// case1
			{A, B, C},
			// case2
			{A, D},
			// case3
			{A, D, C},
			// case4
			{A, E, B, C, D},
			// case5
			{A, E, D},
			// case6: maximum of 3 stops.
			{C, OPTIONAL, OPTIONAL, C},
			// case7: exactly 4 stops.
			{A, ANY, ANY, ANY, C},
			// case8
			{A, C},
			// case9
			{B, B},
			// case10
			{C, C}
	};

	// use to constructors. (adjacnecy matrix)
	private static final int[][] CITY_NETWORK = {
            {0, 5, 0, 5, 7},
            {0, 0, 4, 0, 0},
            {0, 0, 0, 8, 2},
            {0, 0, 8, 0, 6},
            {0, 3, 0, 0, 0}
	};

	public static void main(String[] args) {
		CityNetwork cityNetwork = new CityNetwork(CITY_NETWORK);

		int i = 0;
		// case1~5. The distance of route
		for (; i < 5; i++) {
			try {
				int distance = cityNetwork.routeDistance(ROUTES[i]);
				System.out.printf("Output #%d: %d\n", i + 1, distance);
			} catch (Exception e) {
				System.out.printf("Output #%d: NO SUCH ROUTE\n", i);
			}
		}

		// case6~7. The number of trips
		for (; i < 7; i++) {
			int number = cityNetwork.numberOfTrips(ROUTES[i]);
			System.out.printf("Output #%d: %d\n", i + 1, number);
		}

		// case8~9. The length of the shortest route
		for (; i < 9; i++) {
			int distance = cityNetwork.shortestRouteDistance(ROUTES[i][0], ROUTES[i][1]);
			System.out.printf("Output #%d: %d\n", i + 1, distance);
		}

		// case10. The number of different routes with a distance of less than 30
		int number = cityNetwork.numberOfRoutes(ROUTES[i][0], ROUTES[i][1], 30);
		System.out.printf("Output #%d: %d\n", ++i, number);
	}
}
```
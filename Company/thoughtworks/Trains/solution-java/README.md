# Solution  

## Java  

```java
    public class Graphmain {

    public int[][] map = { 
            {-1, 5, -1, 5, 7}, 
            {-1, -1, 4, -1, -1},
            {-1, -1, -1, 8, 2},
            {-1, -1, 8, -1, 6},
            {-1, 3, -1, -1, -1},
    };
    
    public void dfs(String end, String path, int maxLength)
    {
        if (path.length() - 1 > maxLength) return;
        
        if ( path.length() > 1 && path.endsWith(end) ) {
            System.out.println(path + ", " + path.length());
        }
        
        char lastChar = path.charAt(path.length()-1);
        int lastNodeIndex = lastChar - 'A';
        
        for ( int i=0; i<map[lastNodeIndex].length; i++ )
        {
            char newChar = (char)(i + 'A');
            if ( map[lastNodeIndex][i] > 0) {
                dfs(end, path + newChar, maxLength);
            }
        }
    }
    
    public static void main(String[] args) {
        Graphmain g = new Graphmain();
        
        g.dfs("C", "C", 3);
    }
}



    public class Graphmain {

    public int[][] map = { 
            { -1, 5, -1, 5, 7 }, 
            { -1, -1, 4, -1, -1 },
            { -1, -1, -1, 8, 2 }, 
            { -1, -1, 8, -1, 6 }, 
            { -1, 3, -1, -1, -1 }
    };
    public void bfs(String start, String end, int hops) {
        String lastRoute = start;
        for (int hop = 0; hop < hops; hop++) {
            String route = "";
            for (int i = 0; i < lastRoute.length(); i++) {
                char c = lastRoute.charAt(i);
                int id = c - 'A';
                for (int j = 0; j < map[id].length; j++) {
                    if (map[id][j] > 0)
                        route = route + (char) (j + 'A');
                }
            }
//          System.out.println(lastRoute);
            lastRoute = route;
        }
//      System.out.println(lastRoute);
        System.out.println(lastRoute.split(end).length - 1);
    }
    public static void main(String[] args) {
        Graphmain g = new Graphmain();
        g.bfs("A", "C", 4);
    }
}


    public class Graphmain {

    public int[][] map = { 
            {-1, 5, -1, 5, 7}, 
            {-1, -1, 4, -1, -1},
            {-1, -1, -1, 8, 2},
            {-1, -1, 8, -1, 6},
            {-1, 3, -1, -1, -1},
    };
    
    public void dfs(String end, String path, int cost) {
        if (path.endsWith(end) && cost < bestCost && cost > 0) {
            bestPath = path;
            bestCost = cost;
            return;
        }
        char lastChar = path.charAt(path.length() - 1);
        int lastNodeIndex = lastChar - 'A';
        for (int i = 0; i < map[lastNodeIndex].length; i++) {
            char newChar = (char) (i + 'A');
            int newCost = map[lastNodeIndex][i];
            if (newCost > 0) {
                if (path.indexOf(newChar) > 0)
                    continue;
                dfs(end, path + newChar, cost + newCost);
            }
        }
    }
    public String bestPath = "";
    public int bestCost = Integer.MAX_VALUE;
    public static void main(String[] args) {
        Graphmain g = new Graphmain();
        g.dfs("C", "A", 0); // 8 
//      g.dfs("B", "B", 0); // 9
        System.out.println("Best Path: " + g.bestPath + "/nCost: " + g.bestCost);
    }
}


    public class Graphmain {

    public int[][] map = { 
            {-1, 5, -1, 5, 7}, 
            {-1, -1, 4, -1, -1},
            {-1, -1, -1, 8, 2},
            {-1, -1, 8, -1, 6},
            {-1, 3, -1, -1, -1},
    };
    
    public void dfs(String end, String path, int cost) {
        if (cost >= 30)
            return;
        if (cost > 0 && path.endsWith(end)) {
            System.out.println(path + ", " + cost);
        }
        char lastChar = path.charAt(path.length() - 1);
        int lastNodeIndex = lastChar - 'A';
        for (int i = 0; i < map[lastNodeIndex].length; i++) {
            char newChar = (char) (i + 'A');
            int newCost = map[lastNodeIndex][i];
            if (newCost > 0) {
                dfs(end, path + newChar, cost + newCost);
            }
        }
    }
    public static void main(String[] args) {
        Graphmain g = new Graphmain();
        g.dfs("C", "C", 0);
    }
}
```



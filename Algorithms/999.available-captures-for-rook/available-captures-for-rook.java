class Solution {
    public int numRookCaptures(char[][] board) {
        // 因为题目已经明确给出 board.length == board[i].length == 8，所以不做输入检查
        // 定义方向数组，可以认为是四个方向向量，在棋盘问题上是常见的做法
        int[][] directions = {{-1, 0}, {1, 0}, {0, 1}, {0, -1}};

        for (int i = 0; i < 8; i++) {
            for (int j = 0; j < 8; j++) {
                
                if (board[i][j] == 'R') {
                    int res = 0;
                    for (int[] direction : directions) {
                        if (burnout(board, i, j, direction)) {
                            res++;
                        }
                    }
                    return res;
                }
            }
        }
        // 代码不会走到这里，返回 0 或者抛出异常均可
        return 0;
    }

    /**
     * burnout 横冲直撞的意思（来自欧路词典）
     *
     * @param board     输入棋盘
     * @param x         当前白象位置的横坐标
     * @param y         当前白象位置的纵坐标
     * @param direction 方向向量
     * @return 消灭一个 p，就返回 true
     */
    private boolean burnout(char[][] board, int x, int y, int[] direction) {
        int i = x;
        int j = y;
        while (inArea(i, j)) {
            // 是友军，路被堵死，直接返回
            if (board[i][j] == 'B') {
                break;
            }

            // 是敌军，拿下一血（不知道一血这个词是不是这么用的）
            if (board[i][j] == 'p') {
                return true;
            }

            i += direction[0];
            j += direction[1];
        }
        return false;
    }

    /**
     * @param i 当前位置横坐标
     * @param j 当前位置纵坐标
     * @return 是否在棋盘有效范围内
     */
    private boolean inArea(int i, int j) {
        return i >= 0 && i < 8 && j >= 0 && j < 8;
    }
}

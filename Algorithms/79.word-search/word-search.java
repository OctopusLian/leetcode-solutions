class Solution {
    public boolean exist(char[][] board, String word) {
        if (board == null || board.length == 0 || board[0].length == 0) return false;
        if (word == null || word.length() == 0) return false;

        boolean[][] visited = new boolean[board.length][board[0].length];
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (dfs(board, word, visited, i, j, 0)) {
                    return true;
                }
            }
        }

        return false;
    }
    public boolean dfs(char[][] board, String word,
                               boolean[][] visited,
                               int row, int col,
                               int wi) {
        // out of index
        if (row < 0 || row > board.length - 1 ||
            col < 0 || col > board[0].length - 1) {
            return false;
        }

        if (!visited[row][col] && board[row][col] == word.charAt(wi)) {
            // return instantly
            if (wi == word.length() - 1) return true;
            // traverse unvisited row and col
            visited[row][col] = true;
            boolean down = dfs(board, word, visited, row + 1, col, wi + 1);
            boolean right = dfs(board, word, visited, row, col + 1, wi + 1);
            boolean up = dfs(board, word, visited, row - 1, col, wi + 1);
            boolean left = dfs(board, word, visited, row, col - 1, wi + 1);
            // reset with false if none of above is true
            visited[row][col] = up || down || left || right;
            return up || down || left || right;
        }

        return false;
    }
}

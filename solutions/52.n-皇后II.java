class Solution {

  private boolean[] dg;
  private boolean[] udg;
  private boolean[] col;
  private int res = 0;

  private void dfs(int level, int n) {
    if(level == n) {
      res++;
      return;
    }

    for(int i = 0; i < n; i++) {
      if(!col[i] && !dg[i + level] && !udg[level - i + 100]) {
        col[i] = dg[i + level] = udg[level - i + 100] = true;
        dfs(level + 1, n);
        col[i] = dg[i + level] = udg[level - i + 100] = false;
      }
    }
  }

  public int totalNQueens(int n) {
    dg = new boolean[n*2];
    udg = new boolean[n*2];
    col = new boolean[n];
    dfs(0, n);
    return res;
  }
}
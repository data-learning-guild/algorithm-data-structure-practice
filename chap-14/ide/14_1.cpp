/*
問題: https://atcoder.jp/contests/dp/tasks/dp_g
解説: https://qiita.com/drken/items/03c7db44ccd27820ea0d#g-%E5%95%8F%E9%A1%8C---longest-path

Input:
4 5
1 2
1 3
3 2
2 4
3 4

Output:
3

最長路を求める

解説が分かりやすい。この問題のように DP の更新順序が非自明な場合は、メモ化再帰を使うとよいらしい（ただし別解もある）

dp[v]に、ノードvを始点としたときの有向パスの長さの最大値を記録していく
*/

#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>
using namespace std;

template <typename T>
bool chmax(T &a, const T& b) {
  if (a < b) { a = b; return true; } return false;
}

vector<vector<int>> G;
vector<int> dp; // memoという名前でもよいかも

// ノードvを始点としたときの最大値をreturnする
int rec(int v) {
  // メモを使うことで、計算量の爆発を防ぐ
  if (dp[v] != -1) return dp[v];

  int res = 0;
  for (auto next_v : G[v]) {
    chmax(res, rec(next_v) + 1); // next_vを始点とした場合に比べて最長路は1長くなるので1を加算している
    // 単純に「next_vがあるならresに1を加算」とすればよい気もしたが、計算量が増えそう
  }
  return dp[v] = res;
}

int main() {
  int N, M;
  cin >> N >> M;
  G.assign(N, vector<int>());

  for (int i = 0; i < M; i++) {
    int a, b;
    cin >> a >> b;
    a--, b--;
    G[a].push_back(b);
  }

  dp.assign(N, -1);

  int res = 0;
  // 全ノードについて一通り更新
  for (int v = 0; v < N; v++) {
    chmax(res, rec(v));
  }
  cout << res << endl;
}

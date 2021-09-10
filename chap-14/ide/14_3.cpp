/*
問題: https://atcoder.jp/contests/abc132/tasks/abc132_e
解説: https://drken1215.hatenablog.com/entry/2019/07/01/111500

Input:
4 4
1 2
2 3
3 4
4 1
1 3

Output:
2
1 から出発して、1 回目のけんけんぱでは 1→2→3→4、2 回目のけんけんぱでは 4→1→2→3 と移動することで頂点 3 に辿り着ける

方針としては、S から T への路であって、長さが 3 の倍数であるもののうち最短なものを求め、その長さが 3 かけるいくつであるか出力する。解説コードは読みにくかったので、自分で整理した

dp[v][n]に、インデックスvの頂点に3m + n (n = 0, 1, 2)で達するための最短路長を記録する
*/

#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
#include <iostream>
#include <vector>
#include <queue>
using namespace std;

vector<vector<int>> G;
// vector<bool> discovered; // discoverされた頂点も訪問するので、管理する意味がない。この点が、最も基本的な幅優先探索との違い
vector<vector<int>> dp;

int bfs(int s, int t) {
  queue<int> que;
  dp[s][0] = 0;
  que.push(s);

  while (!que.empty()) {
    int v = que.front();
    que.pop();
    // cout << "v: " << v << endl;

    // vが0のとき、next_vは1, 2
    for (auto next_v : G[v]) {

      rep(i, 3) {
        // いま検討している弧の始点に到達することが可能だったら、次にその弧の終点に行くことができる。この条件と、弧の終点にまだ数値が書き込まれていないという条件の積が成り立ったら、最短路長を書き込むことができる
        if (dp[v][i] != -1 && dp[next_v][(i + 1) % 3] == -1) {
          dp[next_v][(i + 1) % 3] = dp[v][i] + 1;
          que.push(next_v);
        }
      }
    }
  }

  if (dp[t][0] != -1) return dp[t][0] / 3;
  else return dp[t][0];
}

int main() {
  int N, M;
  cin >> N >> M;
  G.assign(N, vector<int>());
  dp.assign(N, vector<int>(3, -1));

  for (int i = 0; i < M; i++) {
    int a, b;
    cin >> a >> b;
    a--, b--;
    G[a].push_back(b);
  }

  int S, T;
  cin >> S >> T;
  S--, T--;
  cout << bfs(S, T) << endl;
}

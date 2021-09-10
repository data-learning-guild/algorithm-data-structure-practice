/*
問題: https://atcoder.jp/contests/abc061/tasks/abc061_d
解説: https://drken1215.hatenablog.com/entry/2019/02/16/075900

Input:
3 3
1 2 4
2 3 3
1 3 5

Output:
7

Input 2:
6 5
1 2 -1000000000
2 3 -1000000000
3 4 -1000000000
4 5 -1000000000
5 6 -1000000000

Output 2:
-5000000000

最長路の長さを求める。始点（1）から到達できる正閉路がある場合は、infと出力する

有向閉路を含むグラフなので、ベルマン・フォード法を使う

提出結果は一部WAとなった。解答でなぜ2Nが使われているかなど理解できなかったので、そのあたりが原因かもしれない
https://atcoder.jp/contests/abc061/submissions/25648584
*/

#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
#include <iostream>
#include <vector>
using namespace std;
using Edge = pair<int, int>; // 弧の終点と重みをペアにする。重みは最大で10^9なので、intで間に合う

const long long INF = 1LL << 60;
vector<vector<Edge>> G;

template <typename T>
bool chmax(T &a, const T& b) {
  if (a < b) { a = b; return true; }
  return false;
}

int main() {
  int N, M; // Nが頂点、Mが辺の数
  cin >> N >> M;
  G.assign(N, vector<Edge>());

  rep(_, M) {
    int a, b, c;
    cin >> a >> b >> c;
    a--, b--;
    G[a].push_back(Edge(b, c));
  }

  bool has_positive_cycle = false;
  vector<long long> dist(N, -INF);
  dist[0] = 0;

  rep(iter, N) {
    // cout << "iter: " << iter << endl;
    rep(v, N) {
      // 一つ前のiterationの時点で始点から到達できていない頂点からの緩和は行わない
      if (dist[v] == -INF) continue;
      // cout << "v: " << v << endl;
      for (auto e : G[v]) {
        // e.firstがvから出た弧の終点、e.secondが弧の重みを指す
        if (chmax(dist[e.first], dist[v] + e.second)) {
          // ここで更新が発生していて、かつそれが最後のiterationだったら、閉路があることが分かる。証明は、p.253
          if (iter == N - 1) has_positive_cycle = true;
        }
      }
    }
  }

  if (!has_positive_cycle) cout << dist[N - 1] << endl;
  else cout << "inf" << endl;
}

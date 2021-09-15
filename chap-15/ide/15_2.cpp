/*
問題: https://atcoder.jp/contests/jag2012autumn/tasks/icpc2012autumn_c

通常の最小全域木のようにコストの総和ではなく、コストの中央値を最小化し、その最小値を出力する

という問題だが、中央値を最小化すれば総和も最小化されるという定理があるらしく、結果としてクラスカル法で解ける

Input:
2 1
1 2 5
4 6
1 2 1
1 3 2
1 4 3
2 3 4
2 4 5
3 4 6

Output:
5
2
*/

#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

struct UnionFind {
  vector<int> par, siz;

  UnionFind(int n) : par(n, -1), siz(n, 1) {}

  int root(int x) {
    if (par[x] == -1) return x;
    return par[x] = root(par[x]);
  }

  bool issame(int x, int y) {
    return root(x) == root(y);
  }

  bool merge(int x, int y) {
    x = root(x); y = root(y);
    if (x == y) return false;
    if (siz[x] > siz[y]) swap(x, y);
    siz[x] += siz[y];
    par[y] = x;
    return true;
  }
};

int main() {
  int N, M;
  while (cin >> N >> M, N) {
    // weight, from, toの順（weightを最初に置いておくと、sortがしやすくなる）
    using Edge = pair<int, pair<int, int>>;
    vector<Edge> edges(M);
    rep(i, M) {
      cin >> edges[i].second.first >> edges[i].second.second >> edges[i].first;
      edges[i].second.first--;
      edges[i].second.second--;
    }

    sort(edges.begin(), edges.end());

    UnionFind uf(N);
    int edge_conter = 0; // 辺の本数
    int res = -1;
    // 次のfor文は|E|回繰り返しを行うが、その中のuf.issame(u, v)などの計算量がO(log|V|)なので（この理由についてはUnion-Findに関するp.185を参照）、全体で計算量はO(|E|log|V|)となる
    for (auto e : edges) {
      int w = e.first;
      int u = e.second.first, v = e.second.second;

      // uとvが同じグループに属するとき、この二頂点をつなぐとサイクルができてしまうのでスキップ
      if (uf.issame(u, v)) continue;
      // そうでなければ、ufに辺u, vを追加
      uf.merge(u, v);
      edge_conter++;

      if (edge_conter == (N + 1) / 2) {
        res = w;
        break;
      }
    }
    cout << res << endl;
  }
}

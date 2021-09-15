/*
問題; https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?lang=jp&id=1350
解説: https://drken1215.hatenablog.com/entry/2020/11/04/221100
提出: https://judge.u-aizu.ac.jp/onlinejudge/review.jsp?rid=5876981#1

ありうる最小全域木すべてに共通して含まれる辺を求める

ありうる最小全域木の辺の集合を求めておいて、最後にそれらの積集合を出すことも考えたが、これだと効率が悪い。そこでまず最小全域木Tを求め、そこに含まれる辺eの1本1本について、最小全域木に必要かどうか次の方法で調べる。すなわち、G - eの最小全域木とTの重みを比べて、Tの重みの方が小さければ、eは必要な辺だと判断する

Input:
4 4
1 2 3
1 3 3
2 3 3
2 4 3

Output:
1 3
（no alternative bridgeの本数は1であり、その重みの和は3）
*/

#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;
using Edge = pair<int, pair<int, int>>;
// Edgeの定義は、以下のページにあるように構造体を使って行うとedge.fromなど可読性の高い書き方ができるので、この方法に慣れたほうがよさそう
// http://algoogle.hadrori.jp/icpc/icpc2014tokyof/

const int INF = 1 << 30;
int N, M;

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
    if (siz[x] < siz[y]) swap(x, y);
    siz[x] += siz[y];
    par[y] = x;
    return true;
  }
};

// 頂点数Nと辺の集合edgesが与えられたうえで、edges_usedに使用した辺を記録しつつ最小全域木を求め、その重みを返す
int kruskal(vector<Edge> edges, vector<Edge> &edges_used, Edge e_to_skip = Edge(-1, make_pair(-1, -1))) {
  UnionFind uf(N);
  int res = 0; // 返り値とする重みの和
  for (auto e : edges) {
    if (e == e_to_skip) {
      // cout << "edge starting from " << e.second.first << " is skipped" << endl;
      continue;
    }
    int w = e.first;
    int u = e.second.first, v = e.second.second;
    // uとvが同じグループに属するとき、この二頂点をつなぐとサイクルができてしまうのでスキップ
    if (uf.issame(u, v)) continue;
    // そうでなければ、ufとedges_usedに辺u, vを追加
    uf.merge(u, v);
    edges_used.push_back(Edge(w, make_pair(u, v)));
    res += w;
  }
  // 以下のrepは、min_st_weightを求める際には不要な操作
  if (e_to_skip.first != -1) {
    // 省く辺eが橋だった場合、全域木を作ることができない。そのような場合、eは必要だとみなすためINFを返す
    rep(i, N) {
      if (!uf.issame(0, i)) return INF;
    }
  }
  return res;
}

int main() {
  cin >> N >> M;
  vector<Edge> edges(M), edges_used;

  rep(i, M) {
    int u, v, w;
    cin >> u >> v >> w;
    u--, v--;
    edges[i] = Edge(w, make_pair(u, v));
  }

  sort(edges.begin(), edges.end()); // これを関数kruskalの中に入れてしまうと、TLEとなる
  int min_st_weight = kruskal(edges, edges_used);
  // cout << "min_st_weight: " << min_st_weight << endl;
  int no_alt_count = 0;
  int no_alt_weight = 0;

  vector<Edge> edges_used_2; // このvectorを使うことはないが、関数kruskalをシンプルなままエラーなく動作させるために用意する
  // このfor文はN - 1回繰り返しを行うが、中のkruskalの計算量がO(M log N)なので？全体としてO(MN log N)となる
  for (auto e : edges_used) {
    // cout << e.second.first << " -> " << e.second.second << endl;

    // eを使わずに最小全域木を作り、その重み総和を求める
    int weight = kruskal(edges, edges_used_2, e);
    // cout << "weight: " << weight << endl;
    if (weight > min_st_weight) {
      no_alt_count++;
      no_alt_weight += e.first;
    }
  }

  cout << no_alt_count << " " << no_alt_weight << endl;
}

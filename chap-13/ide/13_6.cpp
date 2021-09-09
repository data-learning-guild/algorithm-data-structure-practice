/*
解説: https://qiita.com/drken/items/a803d4fc4a727e02f7ba
関連問題: https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=2891

Input:
4 4
0 1
0 2
1 2
1 3
（解説と異なり、0-indexedであることを想定）

Output:
No

Input:
4 4
0 1
1 2
1 3
2 0

Output:
Yes

有向サイクルが含まれるかどうか判定し、含まれるならば一つ求める

サイクルを含むことは、頂点 v から出発した探索が、v から行くことのできる全頂点の探索を終えるよりも先に（すなわち帰りがけの状態になる前に）、v に戻って来ることと同じ。なので、探索を進めてdiscoveredかつfinishedでない頂点に行きがけにぶつかったら、サイクルがあったと判定する

判定のみを行うコードは13_6_simple.cppを参照

合っているかどうか、きちんと確認していない
*/

#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>
#include <stack>
#include <set>
using namespace std;
using Graph = vector<vector<int>>;

vector<bool> discovered, finished; // finishedはなくても「discoveredがtrueかつstackに含まれない」という条件で判定できると思いきや、stackに含まれるかどうか調べるのは面倒そうなので、結局finishedを用意したほうがよい
int v_in_circle = -1; // posだと変数名として分かりにくいのでv_in_circleとする
stack<int> hist; // 訪問履歴
set<int> circle;

void dfs(const Graph &G, int v) { // , int p
  discovered[v] = true;
  hist.push(v);
  // cout << "v: " << v << endl;
  for (auto next_v : G[v]) {
    // cout << "next_v: " << next_v << endl;
    if (finished[next_v]) continue;

    if (discovered[next_v] && !finished[next_v]) {
      v_in_circle = next_v;
      return;
    }

    dfs(G, next_v);

    if (v_in_circle != -1) {
      // 以下の方法でサイクルを表示しようとすると、There is a circleが二度出てきてしまう
      // while (hist.top() != v_in_circle) {
      //   cout << hist.top() << " ";
      //   hist.pop();
      // }
      // cout << hist.top() << endl;
      return;
    }
  }
  finished[v] = true;
}

int main() {
  int N, M;
  cin >> N >> M;

  // グラフ入力受取
  Graph G(N);
  for (int i = 0; i < M; ++i) {
    int a, b;
    cin >> a >> b;
    G[a].push_back(b);
  }

  discovered.assign(N, false);
  finished.assign(N, false);
  dfs(G, 0);

  if (v_in_circle == -1) {
    cout << "There is no circle" << endl;
  }
  else {
    cout << "There is a circle" << endl;
    while (hist.top() != v_in_circle) {
      cout << hist.top() << " ";
      hist.pop();
    }
    cout << hist.top() << endl;
  }
}

// // サイクル復元のための情報
// int pos = -1; // サイクル中に含まれる頂点 pos
// stack<int> hist; // 訪問履歴

// void dfs(const Graph &G, int v, int p) {
//   discovered[v] = true;
//   hist.push(v);
//   for (auto nv : G[v]) {
//     if (nv == p) continue; // 逆流を禁止する

//     // 完全終了した頂点はスルー
//     if (finished[nv]) continue;

//     // サイクルを検出
//     if (discovered[nv] && !finished[nv]) {
//       pos = nv;
//       return;
//     }

//     // 再帰的に探索
//     dfs(G, nv, v);

//     // サイクル検出したならば真っ直ぐに抜けていく
//     if (pos != -1) return;
//   }
//   hist.pop();
//   finished[v] = true;
// }

// int main() {
//   // 頂点数 (サイクルを一つ含むグラフなので辺数は N で確定)
//   int N; cin >> N;

//   // グラフ入力受取
//   Graph G(N);
//   for (int i = 0; i < N; ++i) {
//     int a, b;
//     cin >> a >> b;
//     --a, --b; // 頂点番号が 1-indexed で与えられるので 0-indexed にする
//     G[a].push_back(b);
//     G[b].push_back(a);
//   }

//   // 探索
//   discovered.assign(N, false), finished.assign(N, false);
//   pos = -1;
//   dfs(G, 0, -1);

//   // サイクルを復元
//   set<int> cycle;
//   while (!hist.empty()) {
//     int t = hist.top();
//     cycle.insert(t);
//     hist.pop();
//     if (t == pos) break;
//   }

//   // クエリに答える
//   int Q; cin >> Q;
//   for (int _ = 0; _ < Q; ++_) {
//     int a, b;
//     cin >> a >> b;
//     --a, --b;
//     if (cycle.count(a) && cycle.count(b)) cout << 2 << endl;
//     else cout << 1 << endl;
//   }
// }

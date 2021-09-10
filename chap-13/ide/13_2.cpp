/*
s-tパスが存在するかどうか判定する。頂点sを始点としたグラフ探索を行い、その過程で頂点tを訪問したかどうかを調べる

code_13_4とcode_13_1を参考に
*/

#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>
#include <queue>
using namespace std;
using Graph = vector<vector<int>>;

// 頂点数と辺数、s と t
int N, M, s, t;
vector<bool> discovered;
queue<int> todo;

void bfs(const Graph &G, int v) {
  // int N = G.size(); // 最初にNを宣言しておけば、ここで改めて代入する必要はない
  discovered[s] = true;
  todo.push(s);

  while (!todo.empty()) {
    int v = todo.front();
    todo.pop();

    for (int x : G[v]) {
      if (discovered[x]) continue;

      discovered[x] = true;
      todo.push(x);
    }
  }
}

int main() {
  cin >> N >> M >> s >> t;
  // グラフ入力受取
  Graph G(N);
  for (int i = 0; i < M; ++i) {
      int a, b;
      cin >> a >> b;
      G[a].push_back(b);
  }

  // テスト用
  // N = 4;
  // M = 4;
  // s = 0;
  // t = 3;
  // Graph G = {
  //   {1, 2},
  //   {2, 3},
  //   {},
  //   {}
  // };

  // 頂点 s を始点とした探索
  discovered.assign(N, false);
  bfs(G, s);

  if (discovered[t]) cout << "Yes" << endl;
  else cout << "No" << endl;
}

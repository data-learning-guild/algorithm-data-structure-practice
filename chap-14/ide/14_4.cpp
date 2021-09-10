/*
問題: https://atcoder.jp/contests/arc005/tasks/arc005_3
解説: https://drken1215.hatenablog.com/entry/2021/07/30/024800

Input 1:
4 5
s####
....#
#####
#...g

Output 2:
YES
(1,2), (2,2), (3,2) のいずれかの塀を壊すことで、ゴールに到達できる

Input 2:
1 10
s..####..g

Output 2:
NO

方針としては、各マスを頂点とした有向グラフにおいて、
- 通路マスに向かう辺の重みは0
- 壁マスに向かう辺の重みは1
として最短路長を求め、それが2以下であればYESと出力する。重みが非負であるような重み付きグラフなので、ダイクストラ法を使う

迷路の扱いについては13_4を参考に。解説のコードは計算量を落とす工夫が加えられているので、まずは解説を読まずに自分で考えたほうが書きやすかった

dist[i][j] に、マス (i, j) への最短路長を記録していく

以下のコードでACは出せたが、p.256で解説されているようにusedのような配列を用意した方が計算量が抑えられるかも。そうしないとダイクストラ法とは呼べない？ 動的計画法との違いとして、更新が起こった場合のみその続きを探索をする、という点はあるが（ifの条件にchminを入れている箇所を参照）
*/

#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
#include <iostream>
#include <vector>
#include <queue>
using namespace std;

const int INF = 1000; // 街の長さは500x500なので、路長は最大で998

template <typename T>
bool chmin(T &a, const T& b) {
  if (a > b) { a = b; return true; } return false;
}

int main() {
  int H, W;
  cin >> H >> W;
  vector<string> maze(H); // 可読性を上げるうえでは、stringの配列ではなく、二次元配列で隣接リスト表現で扱うのもよいかもしれない。それをすると計算量がかさむので、今回はこれで問題ない
  rep(i, H) cin >> maze[i];

  // スタートとゴールのマスを割り出す
  int sx = -1, sy = -1, gx = -1, gy = -1;
  for (int i = 0; i < H; i++) {
    for (int j = 0; j < W; j++) {
      if (maze[i][j] == 's') sx = i, sy = j;
      if (maze[i][j] == 'g') gx = i, gy = j;
    }
  }

  using Node = pair<int, int>;
  queue<Node> que;
  vector<vector<int>> move = {{1, 0}, {0, 1}, {-1, 0}, {0, -1}};

  // 初期条件
  que.push({sx, sy});
  vector<vector<int>> dist(H, vector<int>(W, INF));
  dist[sx][sy] = 0;

  while (!que.empty()) {
    auto [x, y] = que.front();
    que.pop();
    // cout << "x: " << x << ", y: " << y << endl;

    for (int dir = 0; dir < 4; ++dir) {
      int nx = x + move[dir][0];
      int ny = y + move[dir][1];

      if (nx < 0 || nx >= H || ny < 0 || ny >= W) continue;
      // cout << "[nx][ny]: " << dist[nx][ny] << ", " << "[x][y]: " << dist[x][y] << endl;

      // ここがポイント
      if (maze[nx][ny] == '#') {
        // 緩和が実行された場合のみ、そのルートの続きを探索したいので、chminをif節の条件として使う
        if (chmin(dist[nx][ny], dist[x][y] + 1)) que.push(Node(nx, ny));
      }
      else {
        if (chmin(dist[nx][ny], dist[x][y])) que.push(Node(nx, ny));
      }
    }
  }

  if (dist[gx][gy] <= 2) cout << "YES" << endl; // けんちょん本の設問に答えるためには、単にdist[gx][gy]を返せばよい
  else cout << "NO" << endl;
}

// 失敗作（掘り方によっては違う解答が出てしまう）
// 13.2のサンプルコードをベースにつくる
#include <iostream>
#include <vector>
using namespace std;
using Graph = vector<vector<int>>;

// 深さ優先探索
vector<int> seen;
void dfs(const Graph &G, int v) {
    seen[v] = v; // v に頂点番号を入れて訪問済にする

    // v から行ける各頂点 next_v について
    for (auto next_v : G[v]) {
        seen[v] = min(seen[v], seen[next_v]);
        if (seen[next_v] != 999) continue; // next_v が探索済ならば探索しない
        dfs(G, next_v); // 再帰的に探索
    }
}

int main() {
    // 頂点数と辺数
    int N, M;
    cin >> N >> M;

    // グラフ入力受取
    Graph G(N);
    for (int i = 0; i < M; ++i) {
        int a, b;
        cin >> a >> b;
        G[a].push_back(b);
        G[b].push_back(a);
    }

    // 探索
    seen.assign(N, 999); // 初期状態では全頂点が未訪問
    for (int v = 0; v < N; ++v) {
        if (seen[v] != 999) continue; // すでに訪問済みなら探索しない
        dfs(G, v);
    }
  
    // 集計
    int res = 0;
    for (int x = 0; x < N; x++) {
        if (seen[x] == x) res++;
    }
    cout << res << endl;
}

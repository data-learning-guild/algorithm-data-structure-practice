// わからなかったので解答をいったん丸写し
// https://github.com/drken1215/book_algorithm_solution/blob/master/solutions/chap05.md

#include <iostream>
#include <vector>
using namespace std;

int main() {
    // 入力
    int N, W;
    cin >> N >> W;
    vector<int> a(N);
    for (int i = 0; i < N; ++i) cin >> a[i];

    // DPテーブル作成
    vector<vector<bool>> dp(N+1, vector<bool>(W+1, false));
  
    // 注： dpは0始まり、配列の○個目は1始まりなので、その点に気をつけること
  
    // 0個目までの数字で0をつくるのは当然可能
    dp[0][0] = true;
    for (int i = 0; i < N; ++i) {
        for (int j = 0; j <= W; ++j) {
            // i-1個目までの数字でjを作れれば、i個目までの数字でjを作ることは可能（使わないだけなので）
            if (dp[i][j]) dp[i+1][j] = true;
            // i個目までの数字でjを作れれば、i個目までの数字でjにa[i]を足したものも作ることが可能
            // (※j+a[i]がWを超えると配列のサイズをオーバーしてしまうため、判定・代入はしないように）
            if (dp[i+1][j] && j+a[i] <= W) dp[i+1][j+a[i]] = true;
        }
    }

    // 答え
    if (dp[N][W]) cout << "Yes" << endl;
    else cout << "No" << endl;
}

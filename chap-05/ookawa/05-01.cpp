#include <iostream>
#include <vector>
using namespace std;

/* 設計の指針
- ①i-1日目がa・b・cだった各場合の幸福度の最大値を計算する
- ②aに至るルート(b→a/c→a)・bに至るルート(a→b/c→b)・cに至るルート(a→c/b→c)を比較し、大きい方をi日目の幸福度の最大値とする
- ③上記をN日目まで繰り返す
*/

int main() {
  // 入力
  int N;
  cin >> N;
  vector<vector<int>> happiness(N, vector<int>(3));
  // 各日の幸福度を入力
  for (int i = 0; i < N; i++) {
    for (int j = 0; j < 3; j++) {
      cin >> happiness.at(i).at(j);
    }
  }
  
  // DPテーブル定義(3列N行)
  vector<vector<int>> dp(N, vector<int>(3, 0));
  
  // DPループ
  for (int i = 0; i < N; i++) {
    for (int j = 0; j < 3; j++) {
      if (i == 0) {
        // 1日目はベタで幸福度を入れる
        dp[i][j] = happiness[i][j];
      }
      else {
        if (j == 0) {
          // aの場合：前日bの場合・前日cの場合のうち大きい方に加算する
          dp[i][j] = max(dp[i-1][1], dp[i-1][2]) + happiness[i][j];
        }
        else if (j == 1) {
          // bの場合：前日aの場合・前日cの場合のうち大きい方に加算する
          dp[i][j] = max(dp[i-1][0], dp[i-1][2]) + happiness[i][j];
        }
        else if (j == 2) {
          // cの場合：前日aの場合・前日bの場合のうち大きい方に加算する
          dp[i][j] = max(dp[i-1][0], dp[i-1][1]) + happiness[i][j];
        }
      }
    }
  }
  
  // 最大値の算出
  cout << max(max(dp[N-1][0], dp[N-1][1]), dp[N-1][2]) << endl;
}

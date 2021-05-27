/*
input:
3
10 40 70
20 50 80
30 60 90

output:
210

AtCoder: https://atcoder.jp/contests/dp/tasks/dp_c
この章のテーマである動的計画法について最初にざっくりまとめておくと、プログラム的には、
複数次元の配列を用意して、左上から順に、それぞれの値が最新になるよう更新をかけていく
これによって同じ計算の繰り返しを避け、効率化を図ることができる

配列dp[i]によってi日目までに得られる幸福度を表すか、i日目が始まる時点での幸福度を表すか二つの立場が考えられる
解答は後者なので、それに沿って考える

解答ではchmaxが以下のようにinlineを使って定義されているが、inlineは使わなくても書ける
template<class T> inline bool chmax(T& a, T b) { if (a < b) { a = b; return 1; } return 0; }
*/

#include <iostream>
#include <vector>
#include <algorithm>
#define _GLIBCXX_DEBUG
using namespace std;

template <typename T>
bool chmax(T &a, const T b) {
  if (a < b) {
    a = b;
    return true;
  }
  return false;
}

int main() {
  int n;
  cin >> n;
  vector<vector<int>> vec(n, vector<int>(3));
  for (int i = 0; i < n; i++) {
    for (int j = 0; j < 3; j++) {
      cin >> vec[i][j];
    }
  };

  // 初日の行動として何を選ぼうとも、初日が始まる時点では満足度0なので初期値は0とする
  vector<vector<int>> dp(n + 1, vector<int>(3, 0));
  // dp[i + 1]に、a_i, b_i, c_iのそれぞれを選んだ場合について、満足度の合計が最大でいくつになるかを格納する
  for (int i = 0; i < n; i++) {
    // i - 1日目の行動
    for (int j = 0; j < 3; j++) {
      // i日目の行動
      for (int k = 0; k < 3; k++) {
        // 「i - 1日目までの（インデックスj以外の）満足度合計 + i日目のインデックスkの行動による満足度」の最大値を
        // i日目が終わる時点での幸福度dp[i + 1]に格納する
        // dp[i + 1]の初期値は0なので、ループの初回で必ず更新される
        if (k == j) continue;
        chmax(dp[i + 1][k], dp[i][j] + vec[i][k]);
      }
    }
  }

  // 最終日までの合計の最大値
  cout << *max_element(dp[n].begin(), dp[n].end()) << endl;
  // ライブラリalgorithmに頼らない場合
  // int res = 0;
  // for (int i = 0; i < 3; ++i) chmax(res, dp[n][i]);
  // cout << res << endl;
}
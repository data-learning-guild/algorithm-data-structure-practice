/*
input 1:
3
7
1
2 3 5

output 1:
NO
7を作るために必要な整数は2と5の2個で、1個より多いため

input 2:
3
7
2
2 3 5

output 2:
YES

解説: https://qiita.com/drken/items/a5e6fe22863b7992efdb#%E5%95%8F%E9%A1%8C-6k%E5%80%8B%E4%BB%A5%E5%86%85%E9%83%A8%E5%88%86%E5%92%8C%E5%95%8F%E9%A1%8C

5_2や5_3で使ったdpを改変して、dp[i][j]に、
最初のi個（インデックスiの手前まで）の整数を足し合わせてjを作るときに必要な整数の最小個数を格納していく
jを作れない場合は、1000を入れておく（これは、n <= 500より大きい）
*/

#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>
#include <numeric>
using namespace std;

template <typename T>
bool chmin(T &a, const T b) {
  if (a > b) {
    a = b;
    return true;
  }
  return false;
}

int main() {
  int n, w, k;
  cin >> n >> w >> k;
  vector<int> vec(n);
  for (int i = 0; i < n; i++) cin >> vec[i];
  int INF = 1000;

  vector<vector<int>> dp(n + 1, vector<int>(w + 1, INF));
  dp[0][0] = 0;

  for (int i = 0; i < n; i++) {
    for (int j = 0; j <= w; j++) {
      if (dp[i][j] < INF) {
        chmin(dp[i + 1][j], dp[i][j]);
      }

      if (j >= vec[i]) {
        chmin(dp[i + 1][j], dp[i][j - vec[i]] + 1);
        // chminを使わず以下のようにする場合、引数の型を合わせるために+ 0が必要（他の対策もあるかもしれない）
        // dp[i + 1][j] = min(dp[i + 1][j] + 0, dp[i][j - vec[i]] + 1);
      }
    }
  }

  if (dp[n][w] <= k) {
    cout << "YES" << endl;
  } else {
    cout << "NO" << endl;
  }
}

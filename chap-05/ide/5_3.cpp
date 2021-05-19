/*
Input:
3
2 3 5

Output:
7
0, 2, 3, 5, 7, 8, 10の7通り

AtCoder: https://atcoder.jp/contests/tdpc/tasks/tdpc_contest
解説: https://drken1215.hatenablog.com/entry/2020/12/21/153600
*/

#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>
#include <numeric>
using namespace std;

int main() {
  int n;
  cin >> n;
  vector<int> vec(n);
  for (int i = 0; i < n; i++) cin >> vec[i];
  int sum = accumulate(vec.begin(), vec.end(), 0);

  // dp[i][j]に、最初のi個（インデックスiの手前まで）の整数を使ってjを作れるかどうかを表す真偽値を格納していく
  vector<vector<bool>> dp(n + 1, vector<bool>(sum + 1, false));
  dp[0][0] = true;

  for (int i = 0; i < n; i++) {
    for (int j = 0; j <= sum; j++) {
      // インデックスiの手前までの整数でjを作れるなら、インデックスiまでの整数でも必ずjを作れる
      if (dp[i][j]) {
        dp[i + 1][j] = true;
        continue;
      }

      // インデックスiの整数vec[i]を選べる場合
      if (j >= vec[i]) {
        dp[i + 1][j] = dp[i + 1][j] || dp[i][j - vec[i]];
      }
    }
  }

  int res = 0;
  for (int j = 0; j <= sum; j++) {
    if (dp[n][j]) res++;
  }
  cout << res << endl;
}

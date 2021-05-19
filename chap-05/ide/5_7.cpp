/*
input:
axyb
abyxb

output:
axb or ayb

AtCoder: https://atcoder.jp/contests/dp/tasks/dp_f
解説: https://qiita.com/drken/items/03c7db44ccd27820ea0d#f-%E5%95%8F%E9%A1%8C---lcs

dp[i][j]に、Sの最初のi文字とTの最初のj文字の最長共通部分列の長さの最大値を格納しておき、
それを使って最長共通部分列を求める
*/

#include <iostream>
#include <string>
#include <vector>
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
  string s, t;
  // cin >> s >> t;
  s = "axyb";
  t = "abyxb";
  vector<vector<int>> dp(s.size() + 1, vector<int>(t.size() + 1, 0));

  for (int i = 0; i < s.size(); i++) {
    for (int j = 0; j < t.size(); j++) {
      chmax(dp[i + 1][j + 1], dp[i + 1][j]);
      chmax(dp[i + 1][j + 1], dp[i][j + 1]);
      if (s[i] == t[j]) chmax(dp[i + 1][j + 1], dp[i][j] + 1);
    }
  }

  // 長さを答える場合
  // cout << dp[s.size()][t.size()] << endl;

  // dpを元に最長の共通部分列を復元
  string res = "";
  int i = s.size(), j = t.size();
  // dpの右下を起点にする
  while (i > 0 && j > 0) {
    if (dp[i][j] == dp[i - 1][j]) {
      i--;
    }
    // このelse if節の内容は、上のif節と入れ替えても問題ない
    else if (dp[i][j] == dp[i][j - 1]) {
      j--;
    }
    // 上のどちらにも当てはまらない場合、左上の数値に1を足したものが降りてきたと分かるので、この経路を遡る
    else {
      res = s[i - 1] + res;
      i--;
      j--;
    }
  }
  cout << res << endl;
}

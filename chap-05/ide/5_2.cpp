/*
input 1:
3
10
7 5 3

output 1:
NO
7 + 3で10を作ることができる

input 2:
4
14
5 6 2 3

output 2:
YES
5 + 6 + 3で14を作ることができる

p.78の、動的計画法を用いたナップサック問題の解法を参考に解く
4.6では、同じ問題を再帰関数を使って解いている
この2つの解法の計算量はともにO(NW)となる（p.57）

模範解答のコードは理解できないので、
https://github.com/drken1215/book_algorithm_solution/blob/master/solutions/chap05.md
Qiita記事を参考にコーディングした
https://qiita.com/drken/items/a5e6fe22863b7992efdb#%E5%95%8F%E9%A1%8C-3%E9%83%A8%E5%88%86%E5%92%8C%E5%95%8F%E9%A1%8C
*/

#include <iostream>
// #include <cstring>
// #include <algorithm>
#include <vector>
#define _GLIBCXX_DEBUG
using namespace std;

int main() {
  int n, w;
  cin >> n >> w;
  vector<int> vec(n);
  for (int i = 0; i < n; i++) cin >> vec[i];

  // dp[i][j]に、最初のi個（インデックスiの手前まで）の整数を足し合わせてjを作れるかどうかを表す真偽値を格納していく
  vector<vector<bool>> dp(n + 1, vector<bool>(w + 1, false));
  dp[0][0] = true;

  // iについてループを回すときはインデックスi + 1の行を操作するので、i < nの不等号に等号はつけなくても最後の行まで更新できる
  for (int i = 0; i < n; i++) {
    for (int j = 0; j <= w; j++) {
      // インデックスiの手前までの整数でjを作れるなら、インデックスiまでの整数でも必ずjを作れる
      if (dp[i][j]) {
        dp[i + 1][j] = true;
        continue;
      }

      // インデックスiの整数vec[i]を選べる場合
      if (j >= vec[i]) {
        // ↓をif文の中に入れない場合、j - vec[i]が負の値となったときエラーが出てしまう
        // rubyなら||=と書けるところだが、効率的な書き方が分からない
        // 解答では|=が使われていたが、cstringなどincludeしてもエラーになる
        dp[i + 1][j] = dp[i + 1][j] || dp[i][j - vec[i]];
      }
    }
  }

  // dpの値を確認
  for (int i = 0; i <= n; i++) {
    for (int j = 0; j <= w; j++) {
      cout << dp[i][j] << " ";
    }
    cout << endl;
  }

  if (dp[n][w]) cout << "YES" << endl;
  else cout << "NO" << endl;
}

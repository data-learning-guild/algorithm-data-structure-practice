/*
AOJ: https://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=2877
（本書の問題文より、AOJの問題文のほうが具体的で状況を想像しやすい。また、AOJではM個の区間ではなくM個以下という指定になっているという違いがあるが、分け方を細かくした方が求める総和は大きくなるはずなので答えは変わらないと思われる）
解説: https://drken1215.hatenablog.com/entry/2020/12/21/202300

Input 1:
3 2
8 4 2
3個の整数を、2個の区間に分けることを考える。すなわち、N = 3, M = 2
8と4の間で分けた場合、6 + 2 = 8
4と2の間で分けた場合、8 + 3 = 11
よって、最大値は11となる

Output 1:
11
*/

#include <iostream>
#include <vector>
#include <iomanip>
#define _GLIBCXX_DEBUG
using namespace std;

template <typename T>
bool chmax(T &a, const T& b) {
  if (a < b) {
    a = b;
    return true;
  }
  return false;
}

const long long INF = 1LL << 60;

int main() {
  int N, M;
  cin >> N >> M;
  vector<long long> vec(N);
  for (int i = 0; i < N; i++) cin >> vec[i];

  // meansには、区間[j, i)の平均値を格納していく（これは下準備で、コードの全体像を把握するためにはdpの方が重要）
  // 解答では行数はN + 1だが、以下のようにNとしても問題ない
  vector<vector<double> > means(N, vector<double>(N + 1));
  for (int i = 1; i <= N; i++) {
    for (int j = 0; j < i; j++) {
      double sum = 0;
      for (int k = j; k < i; k++) sum += vec[k];
      means[j][i] = sum / (i - j);
    }
  }

  // dp[i][k]に、与えられた整数のうち最初のi個についてk個の区間に分割するときの、平均の総和の最大値を格納していく
  vector<vector<double> > dp(N + 1, vector<double>(M + 1, -INF));
  // 最初の0個を0個の区間に分けるという操作は定義できないように思うが、便宜上ここの値は0としないとうまくいかない
  dp[0][0] = 0;
  for (int i = 1; i <= N; i++) { // iの初期値は、解答では0となっているが、1とすべき
    for (int j = 0; j < i; j++) {
      // ここで、jは区間の始点、iは区間の終点の手前のインデックスを指す
      for (int k = 1; k <= M; k++) {
        // dp[j][k - 1]は最初のj個についてk - 1個に分割するときの平均の総和の最大値
        // これに一つの区間[j, i)の平均を足した値で、dp[i][k]を可能ならば更新する
        chmax(dp[i][k], dp[j][k - 1] + means[j][i]);
      }
    }
  }

  // 解答は以下のとおり
  // double res = -INF;
  // for (int m = 0; m <= M; m++) chmax(res, dp[N][m]);
  // cout << fixed << setprecision(6) << res << endl;

  // （M個以下ではなく）M個の区間に分ける場合だけ考えるなら、以下のように書けばOK
  cout << fixed << setprecision(6) << dp[N][M] << endl;
}

/*
解説: https://drken1215.hatenablog.com/entry/2021/02/25/223800
AtCoder: https://atcoder.jp/contests/abc077/tasks/arc084_a

Input:
2
1 5
2 4
3 6

Output:
3

Input:
2
1 5
2 4
4 6

Output:
3
1-2-4, 1-2-6, 1-4-6
*/

#include <iostream>
#include <vector>
#include <algorithm>
#define _GLIBCXX_DEBUG
using namespace std;

int main() {
  int N;
  cin >> N;
  vector<int> a(N), b(N), c(N); // 上部、中部、下部
  for (int i = 0; i < N; i++) cin >> a[i];
  for (int i = 0; i < N; i++) cin >> b[i];
  for (int i = 0; i < N; i++) cin >> c[i];

  sort(a.begin(), a.end());
  sort(b.begin(), b.end());
  sort(c.begin(), c.end());

  // b[j] を固定して考える
  long long res = 0;
  for (int j = 0; j < N; j++) {
    // aの要素のうち、b[j]未満のものの個数
    long long a_j = lower_bound(a.begin(), a.end(), b[j]) - a.begin();
    // cの要素のうち、b[j]より大きいものの個数
    // 解説では「b[j]以上の」個数を求めると書いていたが、これは問題の要件を満たさない気がする
    // 本当にb[j]以上の個数を求める場合は、b[j]未満のものの数をlower_boundで求めて、それをNから引くことになる
    long long c_j = N - (upper_bound(c.begin(), c.end(), b[j]) - c.begin());
    // cout << "j: " << j << ", a_j: " << a_j << ", c_j: " << c_j << endl;
    res += a_j * c_j;
  }
  cout << res << endl;
}

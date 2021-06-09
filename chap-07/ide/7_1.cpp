/*
Input 1:
3
1 2 3
0 2 4

Output 1:
2

正しさをどう検証する？
*/

#include <iostream>
#include <vector>
#include <algorithm>
#define _GLIBCXX_DEBUG
using namespace std;

int main() {
  int N;
  cin >> N;
  vector<int> A(N), B(N);
  // rep(i, N)は、Pythonでいうfor i in range(N)のような命令
  rep(i, N) cin >> A[i];
  rep(i, N) cin >> B[i];

  // 計算量はそれぞれO(N log N)
  sort(A.begin(), A.end());
  sort(B.begin(), B.end());

  // この値以上のaを使ってループを回す
  int INF = 1 << 30;
  int tmp = -INF;
  int cnt = 0;
  // Aの要素を（小さい順に）aに入れる
  for (auto a : A) {
    if (a < tmp) continue;

    // Bから、aより大きいうちで最小の要素を求める（計算量はlog N）
    // この要素がBにおいて消費されるとともに、次にAからこれ以上で最小の要素をaとして使う
    tmp = *upper_bound(B.begin(), B.end(), a);
    cnt += 1;
  }

  cout << cnt << endl;
}

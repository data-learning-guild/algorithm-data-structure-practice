#include <bits/stdc++.h>
#include <string>
using namespace std;

int main() {
  // 入力受け取り
  int N;
  string S;
  cin >> N >> S;
  
  long total = 0;
  // 各桁間に+を入れるか入れないか、を0・1で表現すると考える。
  // すると、2の(N-1)乗通りの数式ができうると考えられる。
  for (int bit = 0; bit < (1 << (N-1)); bit++) {
    // 途中数を格納しておく変数
    string tmp = {S.at(0)};
    for (int i = 0; i < (N-1);  i++) {
      // 1の場合：その桁までで区切り、できた値を加算
      if (bit & (1 << i)) {
        total += stol(tmp);
        // 足したのでそれまでの数はリセット
        tmp = "";
      }
      // 次の数字を格納しておく
      tmp += S.at(i+1);
    }
    // 最後に残った数を加算
    total += stol(tmp);
  }
  
  // 結果出力
  cout << total << endl;
}

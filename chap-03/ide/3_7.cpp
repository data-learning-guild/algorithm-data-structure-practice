/*
想定する入力
3
152
AAPG4bのch3_05の例題が似ている
*/

#include <iostream>
#include <bitset>
using namespace std;

int main() {
  int n;
  string s;
  cin >> n >> s;
  int s_size = s.size();
  int result = 0;

  // bitおよびbsには、+をどこに入れるかを表す数値が入る
  for (int bit = 0; bit < (1 << s_size - 1); bit++) {
    bitset<19> bs(bit); // nは最大で20と仮定（s_size - 1を<>の中に書くことはできないため）

    // 計算すべき数式を文字列で表してみる（evalのような命令がデフォルトでは用意されていないので、これを直接評価することはできない）
    // string f;
    // f = s[0];
    // for (int i = 0; i < s_size - 1; i++) {
    //   if (bs.test(i)) f += " + ";
    //   f += s[i + 1];
    // }
    // cout << f << endl;

    int tmp = 0;
    for (int i = 0; i < s_size - 1; i++) {
      tmp *= 10;
      tmp += s[i] - '0';
      if (bs.test(i)) {
        // やりたい処理は、bsを左から見ていって1であればそこに+を入れることだが、コードとして簡単になるように、今のtmpをresultに加えてからtmpをリセットする
        result += tmp;
        // cout << "bs: " << bs << ", tmp: " << tmp << endl;
        tmp = 0;
      }
    }

    // 最後の+より後ろの部分の加算処理
    tmp *= 10;
    tmp += s.back() - '0';
    result += tmp;
  }

  cout << result << endl;
}
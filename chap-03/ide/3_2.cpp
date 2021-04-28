/*
想定する入力
カウントする数値v、配列の要素数、配列の値の順に与える
5
4
3 5 1 9
*/

#include <iostream>
#include <vector>
using namespace std;

int main() {
  int v, n;
  cin >> v >> n;
  vector<int> vec(n);
  for (int i = 0; i < n; i++) cin >> vec[i];
  // for (auto &e : vec) cin >> e; // 次のように短く書くことも可能

  // 番兵を使えば、さらに定数倍高速化することは可能と思われる
  int v_count = 0;
  for (int i = 0; i < n; i++) {
    if (vec[i] == v) {
      v_count++;
    }
  }
  cout << v_count << endl;
}
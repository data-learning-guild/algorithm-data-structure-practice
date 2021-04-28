/*
想定する入力
配列の要素数、配列の値の順に与える
5
4 36 24 8 16
*/

#include <iostream>
#include <vector>
using namespace std;
#define _GLIBCXX_DEBUG

int how_many_twos(int num) {
  int two_count = 0;
  while (num % 2 == 0) {
    num /= 2;
    two_count++;
  }
  return two_count;
}

int main() {
  int n;
  cin >> n;
  vector<int> vec(n);

  for (int i = 0; i < n; i++) cin >> vec[i];

  int result = 10000; // 与えられる数値は10000未満であると仮定できる場合
  for (auto e : vec) {
    result = min(result, how_many_twos(e));
  }
  cout << result << endl;
}
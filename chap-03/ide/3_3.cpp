/*
想定する入力
7
3 5 1 9 2 10 4
*/

#include <iostream>
#include <vector>
using namespace std;

int main() {
  int n;
  cin >> n;
  vector<int> vec(n);
  for (int i = 0; i < n; i++) cin >> vec[i];

  int first = vec[0];
  int second = 10000; // 与えられる数値は10000未満であると仮定できる場合
  for (int i = 1; i < n; i++) {
    if (vec[i] < first) {
      second = first;
      first = vec[i];
    }
    else if (vec[i] < second) {
      second = vec[i];
    }
  }
  cout << second << endl;
}
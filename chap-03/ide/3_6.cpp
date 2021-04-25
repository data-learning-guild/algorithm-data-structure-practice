/*
想定する入力
K, Nの順に与える
3 3
*/

#include <iostream>
using namespace std;

int main() {
  int k, n;
  cin >> k >> n;

  int counter = 0;
  for (int x = 0; x <= min(k, n); x++) {
    // 模範解答では、以下のn - xの箇所がnとなっていた
    for (int y = 0; y <= min(k, n - x); y++) {
      int z = n - (x + y);
      if (z >= 0 && z <= k) counter++;
    }
  }

  cout << counter << endl;
}
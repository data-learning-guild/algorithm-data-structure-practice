#include <bits/stdc++.h>
using namespace std;
const int INF = 20000000;

int main() {
  // 入力受け取り
  int N;
  cin >> N;
  vector<int> a(N);
  for (int i = 0; i < N; i++) cin >> a[i];
  
  // 最小値と2番目に小さい値を線形探索
  int first_min = INF - 1;
  int second_min = INF;
  for (int i = 0; i < N; i++) {
    if (a[i] < first_min) {
      // 暫定最小値を2番目の最小値に移し、新しい最小値を代入
      second_min = first_min;
      first_min = a[i];
    }
    else if (a[i] < second_min) {
      // 暫定の2位最小値を新しい2位最小値に置き換える
      second_min = a[i];
    }
  }
  
  // 結果出力
  cout << second_min << endl;
}

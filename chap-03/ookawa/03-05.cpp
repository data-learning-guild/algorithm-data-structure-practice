#include <bits/stdc++.h>
using namespace std;
const int INF = 20000000;

int main() {
  // 入力受け取り
  int N;
  cin >> N;
  vector<int> a(N);
  for (int i = 0; i < N; i++) cin >> a[i];
  
  // 最小の回数を記録する変数を作成
  int min_div_cnt = INF;
  for (int i = 0; i < N; i++) {
    int div_cnt = 0;
    // 2で割れる＝2で割った余りが0
    for (int tmp = a[i]; tmp % 2 == 0; tmp /= 2) {
      // ２で割った回数を記録する
      div_cnt++;
    }
    if (div_cnt < min_div_cnt) {
      // 最小回数を更新
      min_div_cnt = div_cnt;
    }
  }
  
  // 結果出力
  cout << min_div_cnt << endl;
}

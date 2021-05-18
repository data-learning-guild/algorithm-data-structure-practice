#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;

/* 設計の指針
- ①i-1個目の整数を使って得られるW以下の総和のパターンを配列に記録
- ②全ての整数に対してi個目の整数を足し合わせ、W以下ものを配列に追加
- ③最後にi個目の整数そのものを配列に追加し、配列の重複を除去してユニークにする
- ④「②→③」をN-1個目の整数まで繰り返したとき、配列の要素数が求める解となる
*/

int main() {
  // 入力
  int N,W;
  cin >> N >> W;
  vector<int> nums(N);
  // 入力された整数をnumsに格納
  for (int i = 0; i < N; i++) {
    cin >> nums.at(i);
  }
  
  // 総和のパターンを入れる配列を作成
  vector<int> sums;
  
  // DPループ
  for (int x : nums) {
    for (int y : sums) {
      if (x + y <= W) {
        // ②和がW以下のものを配列に追加
        sums.push_back(x + y);
      }
    }
    if (x <= W) {
      // ③最後にナマの整数を追加
      sums.push_back(x);
    }
    
    // ③重複を取り除く 
    sort(sums.begin(), sums.end());
    sums.erase(unique(sums.begin(), sums.end()), sums.end());
  }
  
  // ④答えとなる個数の算出
  cout << sums.size() << endl;
}

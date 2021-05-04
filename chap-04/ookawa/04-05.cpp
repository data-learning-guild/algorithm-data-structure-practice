// わからなかったので公式解答を参照した
// https://github.com/drken1215/book_algorithm_solution/blob/master/solutions/chap04.md
#include <bits/stdc++.h>
using namespace std;

void nagomi(int N, int ans, int use_flg, int &cnt) {
  // ベースケース
  if (ans > N) return; // 規定値を超えたら終了
  if (use_flg == 0b111) cnt++; // 3・5・7を全て使っている場合はカウントアップ

  /* note
   0bは2進数を表すプレフィックス。「|」はビット和を表す演算子
   use_flgは3ビットで頭から（7の使用有無 / 5の使用有無 / 3の使用有無）を表している
  */
  
  // 7を末尾に付け加える
  nagomi(N, ans * 10 + 7, use_flg | 0b100, cnt);

  // 5を末尾に付け加える
  nagomi(N, ans * 10 + 5, use_flg | 0b010, cnt);
  
  // 3を末尾に付け加える
  nagomi(N, ans * 10 + 3, use_flg | 0b001, cnt);
}

int main() {
  int N; cin >> N;
  int cnt = 0;
  nagomi(N, 0, 0b000, cnt);
  cout << cnt << endl;
}

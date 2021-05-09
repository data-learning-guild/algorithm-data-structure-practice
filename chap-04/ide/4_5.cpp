#include <iostream>
#include <bitset>
using namespace std;

// bitsetを使う場合、int is_usedはbitset<3> is_usedとする
// counterは、一つの変数に加算していきたいことから参照渡しとなっていることに注意
void count_753(int k, int cur, int is_used, int &counter) {
  if (cur > k) return;
  if (is_used == 0b111) counter++; // bitsetを使うなら、条件式はis_used.all()

  count_753(k, cur * 10 + 7, (is_used | 0b001), counter); // curの後ろに7をつけた数で条件を満たすものがいくつあるかカウント（以下の行についても同様）
  count_753(k, cur * 10 + 5, (is_used | 0b010), counter);
  count_753(k, cur * 10 + 3, (is_used | 0b100), counter);
  // count_753(k, cur * 10 + 3, (is_used | (1 << 2)), counter); // bitsetを使う場合にも対応した書き方。整数bitと(1 << i)を&などでつなぐ書き方は、たとえばp.37にも登場する
}

int main() {
  int k;
  cin >> k;
  int counter = 0;
  count_753(k, 0, 0, counter);
  cout << counter << endl;
}
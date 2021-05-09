/*
Input:
4
14
5 6 2 3

Output:
Yes
*/

#include <iostream>
#include <vector>
#include <map>
#define _GLIBCXX_DEBUG
using namespace std;

vector<vector<int>> memo;

// vecの最初のi個からいくつか選んで、総和をwにできるかどうかをboolで返す
// vecは参照渡しにしているのは、値渡しに伴う処理を省いて効率をよくするためと思われる
bool func(int i, int w, vector<int> &vec) {
  // ベースケース
  if (i == 0) {
    if (w == 0) return true;
    else return false;
  }
  else if (memo[i][w] != -1) {
    return memo[i][w];
  }

  // vec[i - 1]を選ばない場合
  if (func(i - 1, w, vec)) {
    return memo[i][w] = true;
  }
  // vec[i - 1]を選ぶ場合
  if (func(i - 1, w - vec[i - 1], vec)) {
    return memo[i][w] = true;
  }
  // vec[i - 1]を選んでも選ばなくてもtrueが返ってこない場合、falseを返す
  return memo[i][w] = false;
}

int main() {
  int n, w;
  cin >> n >> w;
  vector<int> vec(n);
  for (auto &e : vec) cin >> e;

  memo.assign(n + 1, vector<int>(w + 1, -1));
  if (func(n, w, vec)) cout << "Yes" << endl;
  else cout << "No" << endl;
}
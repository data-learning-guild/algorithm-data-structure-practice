// 4_1の解答は、この解答からメモ化の手続きを省略したもの
#include <iostream>
#include <vector>
#define _GLIBCXX_DEBUG
using namespace std;

vector<int> memo;

int tribo(int n) {
  if (n == 0 || n == 1) return 0;
  else if (n == 2) return 1;
  else if (memo[n - 1] > -1) return memo[n - 1];
  return tribo(n - 1) + tribo(n - 2) + tribo(n - 3);
}

int main() {
  int n;
  cin >> n;

  // memoの長さはn + 1とした方が、上でmemo[n - 1]の代わりにmemo[n]と書けるので可読性が上がるかも
  memo.assign(n, -1);
  cout << tribo(n) << endl;
}
/*
問題: https://atcoder.jp/contests/abc131/tasks/abc131_d
解説: https://drken1215.hatenablog.com/entry/2019/06/22/224800

Input:
5
2 4
1 9
1 8
4 9
3 12

Output:
Yes

締め切りが最も早い仕事から実施していく
*/

#include <iostream>
#include <vector>
#include <algorithm>
#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
using namespace std;

bool cmp(pair<int, int> a, pair<int, int> b) {
  return a.second < b.second;
}

int main() {
  int N;
  cin >> N;
  vector<pair<int, int>> tasks(N);
  rep(i, N) cin >> tasks[i].first >> tasks[i].second;
  // 締め切り、すなわち2番目の要素でソートするため、別に関数を用意する
  sort(tasks.begin(), tasks.end(), cmp);

  bool is_in_time = true;
  long long sum = 0;
  rep(i, N) {
    sum += tasks[i].first;
    if (sum > tasks[i].second) is_in_time = false;
  }
  if (is_in_time) cout << "Yes" << endl;
  else cout << "No" << endl;
}

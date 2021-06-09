/*
問題: https://atcoder.jp/contests/abc091/tasks/arc092_a
解説:  https://youtu.be/DqqPuIZvJTk?t=1894

Input:
3
2 0
3 1
1 3
4 2
0 4
5 5

Output:
2

青い点を左から右に向かって見ていく場合、どの赤い点とペアにするかを決めるに当たっては、青い点より左下の範囲で、なるべく下にあるものを残すという方針で考えればよい
*/

#include <iostream>
#include <vector>
#include <tuple>
#include <algorithm>
#define _GLIBCXX_DEBUG
#define rep(i, n) for (int i = 0; i < (int)(n); i++)
using namespace std;

bool sort_by_second_desc(tuple<int, int, bool> &a, tuple<int, int, bool> &b) {
  return get<1>(a) > get<1>(b);
}

int main() {
  int N;
  cin >> N;
  // tupleの要素は順に、x座標、y座標、その点が既に使われたかどうか。boolの初期値はfalseになるもよう
  vector<tuple<int, int, bool>> reds(N);
  vector<tuple<int, int>> blues(N);
  // その赤い点が使われたかどうかを座標と一緒に管理してしまうためにtupleを採用したが、C++ではpairの方が以下のようにシンプルに書けるのでこれを採用し、usedの管理には別の配列is_usedを使うほうがよかったかもしれない
  // rep(i, N) cin >> reds[i].first >> reds[i].second;
  // rep(i, N) cin >> blues[i].first >> blues[i].second;
  rep(i, N) cin >> get<0>(reds[i]) >> get<1>(reds[i]);
  rep(i, N) cin >> get<0>(blues[i]) >> get<1>(blues[i]);
  // 赤い点はなるべく上にあるものから使っていきたいので、予めインデックス1の要素で降順ソートしておく
  sort(reds.begin(), reds.end(), sort_by_second_desc);
  sort(blues.begin(), blues.end());
  int cnt = 0;

  for (auto blue : blues) {
    for (auto &red : reds) {
      if (get<2>(red) || get<0>(red) >= get<0>(blue) || get<1>(red) >= get<1>(blue)) {
        continue;
      }

      get<2>(red) = true;
      cnt++;
      // cout << get<1>(red) << ", " << get<2>(red) << endl;
      break;
    }
  }

  cout << cnt << endl;
}
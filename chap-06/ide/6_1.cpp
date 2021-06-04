/*
何番目に小さい値か

Input:
5
12 43 7 15 9

Output:
2 4 0 3 1
*/

#define _GLIBCXX_DEBUG
#include <iostream>
#include <vector>
#include <algorithm> // sortを使うために必要
using namespace std;

int main() {
  int N;
  cin >> N;
  vector<int> a(N);
  for (int i = 0; i < N; i++) cin >> a[i];

  auto b = a;
  sort(b.begin(), b.end());

  vector<int> res(N);
  for (int i = 0; i < N; i++) {
    // lower_boundでa[i]以上で最小のもの（を指すイテレータ）を調べ、そこからbの先頭を指すイテレータを差し引くと、bの中で何番目に小さいか（ただし先頭の数は0とする）分かる
    // lower_boundの計算量はO(log N)で、ループをN回まわすので、プログラム全体の計算量はO(N log N)となる
    res[i] = lower_bound(b.begin(), b.end(), a[i]) - b.begin();
    // cout << "i: " << i << ", res[i]: " << res[i] << endl;
  }

  for (int i = 0; i < N; ++i) cout << res[i] << endl;
}
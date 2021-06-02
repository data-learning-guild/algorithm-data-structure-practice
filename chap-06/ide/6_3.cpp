/*
解説: https://drken1215.hatenablog.com/entry/2021/02/25/223800
問題（日本情報オリンピック）: https://www.ioi-jp.org/joi/2007/2008-ho-prob_and_sol/2008-ho.pdf#page=6

解説が分かりやすい

Input:
4 50
3
14
15
9

Output:
48
*/

#include <iostream>
#include <vector>
#include <algorithm>
#define _GLIBCXX_DEBUG
using namespace std;

int main() {
  int N, M;
  cin >> N >> M;
  vector<int> A(N);
  for (int i = 0; i < N; i++) cin >> A[i];
  A.push_back(0); // どの整数も選ばないことを表す0を追加

  // 2つの整数の和を格納するための配列Bを用意
  vector<int> B;
  for (int i = 0; i < A.size(); i++) {
    for (int j = 0; j < A.size(); j++) {
      B.push_back(A[i] + A[j]);
    }
  }
  // オーダーには影響しないかもしれないが、ここでBから重複を取り除く作業を入れると処理を効率化できそう
  sort(B.begin(), B.end());

  int res = 0;
  for (int b : B) {
    // Bの要素でM - bより大きいもののうち、最小のものを指すイテレータ
    auto it = upper_bound(B.begin(), B.end(), M - b);
    if (it == B.begin()) continue;
    // イテレータが指す要素を一つ前、つまりM - b以下で最大のものとする
    it--;
    res = max(res, b + *it);
  }
  cout << res << endl;
}
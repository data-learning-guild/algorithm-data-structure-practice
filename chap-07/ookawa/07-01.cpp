/*
--解答の指針--
1. aとbの配列をソートする{計算量：O(NlogN)}
2. aiでループをかけ、各ステップにて以下の処理を行う{計算量：O(N)}
   [1]bjにてループ処理を開始する{j=nextからスタートすることで計算量がO(N)に}
   [2]ai<bjの場合、N-jで条件を満たす組数をカウントしてループを抜ける{next=jで次のループ開始位置を記録}
   [3]ai>=bjの場合、次のループへ
3. ループが終わったら、カウントした組数を出力
*/
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    // 入力を受け取る
    int N;
    cin >> N;
    vector<int> a(N), b(N);
    for (int i = 0; i < N; ++i) cin >> a[i];
    for (int i = 0; i < N; ++i) cin >> b[i];

    // 条件を満たす整数組の数を格納する変数
    int cnt = 0;
    // bの次のループ開始位置を格納する変数
    int next = 0;

    // 1. aとbの配列をソートする
    sort(a.begin(), a.end());
    sort(b.begin(), b.end());
  
    // 2. aiでループをかける
    for (int i = 0; i < N; i++) {
        // [1]bjにてループ処理を開始
        for (int j = next; j < N; j++) {
          if (a[i] < b[j]) {
            // [2]ai<bjの場合、N-jで条件を満たす組数をカウントしてループを抜ける{next=jで次のループ開始位置を記録}
            cnt += N - j;
            next = j;
            break;
          }
          // [3]ai>=bjの場合、次のループへ
          continue;
        }
    }
    cout << cnt << endl;
}

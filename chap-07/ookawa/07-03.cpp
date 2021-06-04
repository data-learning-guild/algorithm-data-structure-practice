// 7.2のサンプルコード(116ページ）をベースにつくる
#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>
using namespace std;

// 所要時間と締切時刻の組を pair<int,int> で表す
typedef pair<int,int> Work;

// 仕事を締切時刻で大小比較する関数
bool cmp(const Work &a, const Work &b) {
    return a.second < b.second;
}

int main() {
    // 入力
    int N;
    cin >> N;
    vector<Work> works(N);
    for (int i = 0; i < N; ++i)
        cin >> works[i].first >> works[i].second;

    // 締切時刻が早い順にソートする
    sort(works.begin(), works.end(), cmp);

    // 貪欲に処理していく
    string flg = "Yes";
    int current_time = 0;
    for (int i = 0; i < N; ++i) {
        // 仕事の処理完了時間を計算
        current_time += works[i].first;
        // 締切時刻に終わっていなければ、対応不可として終了
        if (works[i].second < current_time) {
          flg = "No";
          break;
        }
        // 締切時刻に終わっていれば、次の仕事へ
        continue;
    }
    cout << flg << endl;
}

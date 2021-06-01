// わからなかったので解答を丸写し
// https://github.com/drken1215/book_algorithm_solution/blob/master/solutions/chap06.md

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

// 無限大を表す値←rightの値に使うだけなので、ソート後のa[n-1]とするのが最大効率かも？
const long long INF = 1LL<<50;

int main() {
    // 入力
    int N, M;
    cin >> N >> M;
    vector<long long> a(N);
    for (int i = 0; i < N; ++i) cin >> a[i];
    sort(a.begin(), a.end());

    // 「すべての間隔を x 以上にして M 個の小屋を選べるか？」という判定問題を解く
    // Yes であるような最大の x を求める
    long long left = 0, right = INF;
    while (right - left > 1) {
        long long x = (left + right) / 2;

        int cnt = 1; // 何個の小屋が取れたか (最初の小屋は選ぶ)
        int prev = 0; // 前回の小屋の index
        for (int i = 0; i < N; ++i) {
            // x 以上の間隔が空いたら採用する
            if (a[i] - a[prev] >= x) {
                ++cnt;
                prev = i;
            }
        }

        // 判定問題が Yes か No か（真ん中で切った左側を選ぶか、右側を選ぶかの判断）
        if (cnt >= M) left = x;
        else right = x;
    }
    cout << left << endl;
}

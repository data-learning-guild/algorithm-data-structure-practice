// わからなかったので解説を丸写しで作成し、コメント書き下しをした
// https://github.com/drken1215/book_algorithm_solution/blob/master/solutions/chap05.md

#include <iostream>
#include <vector>
using namespace std;

template<class T> inline bool chmin(T& a, T b) {
    if (a > b) {
        a = b;
        return 1;
    }
    return 0;
}

// 無限大を表す値
const long long INF = 1LL<<60;

int main() {
    // 入力
    int N;
    cin >> N;
    vector<long long> a(N);
    for (int i = 0; i < N; ++i) cin >> a[i];

    // 累積和をとる(累積和→1個目まで・2個目まで・…・n個目までの総和)
    vector<long long> S(N+1, 0);
    for (int i = 0; i < N; ++i) S[i+1] = S[i] + a[i];

    // DP（i個目の仕切り〜j個目の仕切りといったように区間を指定するので、二次元に）
    vector<vector<long long>> dp(N+1, vector<long long>(N+1, INF));

    // 初期条件（i個目の仕切り〜i+1個目の仕切り…スライム単体を指すので合体せずコストは0）
    for (int i = 0; i < N; ++i) dp[i][i+1] = 0;

    // 更新
    for (int bet = 2; bet <= N; ++bet) {
        for (int i = 0; i + bet <= N; ++i) {
            // 1個先の仕切りは初期条件で確認しているため、2個以上先の仕切りを指定
            int j = i + bet;

            // dp[i][j] を更新する（i〜jの間に1箇所ずつ仕切り"k"を入れ、全パターン検証）
            for (int k = i+1; k < j; ++k) {
                /* 
                  iからjまでのスライムをつくる最小コストについての計算式
                  (iからkまでのスライムをつくる最小コスト)+(kからjまでのスライムをつくる最小コスト)+(両方のスライムを合体させるコスト)
                  ※両方のスライムを合体させるコストは、iからjまでのスライムの大きさの総和と常に一致する
                */
                chmin(dp[i][j], dp[i][k] + dp[k][j] + S[j]-S[i]);
            }
        }
    }
    // 1個目からN+1個目の仕切りまでのスライムをつくる最小コストが正解となる
    cout << dp[0][N] << endl;
}

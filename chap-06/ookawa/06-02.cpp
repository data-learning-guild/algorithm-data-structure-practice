// 6.4のコードをベースに作成
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;
const long INF = 2000000000; // 十分大きな値に

int main() {
    // 入力を受け取る
    int N;
    cin >> N;
    vector<int> A(N), B(N), C(N);
    for (int i = 0; i < N; ++i) cin >> A[i];
    for (int i = 0; i < N; ++i) cin >> B[i];
    for (int i = 0; i < N; ++i) cin >> C[i];

    // 条件を満たす整数組の数を格納する変数
    long cnt = 0;

    // A,C をソート
    sort(A.begin(), A.end());
    sort(C.begin(), C.end());

    // A,C に無限大を表す値 (INF) を追加しておく
    // これを行うことで、iter_* = *.end() となる可能性を除外する
    A.push_back(INF);
    C.push_back(INF);

    // B を固定して解く
    for (int i = 0; i < N; i++) {
        // A の中で B[i] 以上の範囲での最小値を示すイテレータ
        auto iter_a = lower_bound(A.begin(), A.end(), B[i]);
        // iter_aのインデックスを取り出す
        int idx_a = distance(A.begin(), iter_a);
        // A[0]からA[idx_a - 1]までは条件を満たすので、その数を導出
        long cnt_a = idx_a;
      
        // C の中で B[i] より大きい範囲での最小値を示すイテレータ
        auto iter_c = upper_bound(C.begin(), C.end(), B[i]);
        // iter_cのインデックスを取り出す
        int idx_c = distance(C.begin(), iter_c);
        // C[idx_c]からC[N-1]までは条件を満たすので、その数を導出
        long cnt_c = (N - 1) - (idx_c - 1);
      
        // あるBに対して、条件を満たすA・Cの組の数は cnt_a * cnt_c で求められる
        cnt += cnt_a * cnt_c;
    }
    cout << cnt << endl;
}

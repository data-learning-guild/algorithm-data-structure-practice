#include <bits/stdc++.h>
using namespace std;

// 47頁のcode 4.5をもとに3項へ展開
int tribo(int N) {
    // ベースケース
    if (N == 0) return 0;
    else if (N == 1) return 0;
    else if (N == 2) return 1;

    // 再帰呼び出し
    return tribo(N - 1) + tribo(N - 2) + tribo(N - 3);
}

int main() {
    int N; cin >> N;
    cout << tribo(N) << endl;
}

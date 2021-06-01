// 難しかったので以下のURLを参考にしつつ…
// http://mmxsrup.hatenablog.com/entry/2016/09/21/162734
#include <iostream>
#include <iomanip>
#include <cmath>
using namespace std;

const double pi = 3.14159265358979;
double A,B,C;

// f(t)の条件式を満たすかどうか
double f(double t) {
    return (A * t) + (B * sin(C * t * pi));
}

int main() {
    // 入力値を受け取る
    cin >> A >> B >> C;
  
    // f(t)の関数の構造的に、t=0とt=201の間で確実に解は見つかる
    // [理由]A=1,B=100のとき、t=201でA*t=201となり、Bsin(CtΠ)の最小は-100であるため
    // →連続関数の中間値の定理より、f(0)=0とf(201)=101以上で100を挟み込める
    double l = 0.0, r = 402.0;
  
    // 二分探索を行い、精度をブラッシュアップしていく
    for(int i = 0; i < 10000; i++) {
        double mid = (l + r) / 2.0;
        // 真ん中で100より小さいなら、真ん中より右側に条件を満たすtが存在する
        if(f(mid) < 99.999999) l = mid;
        // 真ん中で100より大きいなら、真ん中より左側に条件を満たすtが存在する
        else if (f(mid) > 100.000001) r = mid;
      
        // 必要な精度を満たした場合、tを出力して終了
        else {
          // 注: coutのみだと第6位までしか出力されず、精度不足になる
          cout << setprecision(15) << mid << endl;
          break;
        }
    }
}

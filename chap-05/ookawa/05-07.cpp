// わからなかったので解説をベースにコードを作成
// https://qiita.com/drken/items/a5e6fe22863b7992efdb#%E5%95%8F%E9%A1%8C-8%E6%9C%80%E9%95%B7%E5%85%B1%E9%80%9A%E9%83%A8%E5%88%86%E5%88%97-lcs-%E5%95%8F%E9%A1%8C
// https://itliberta.tech/ichi-ichi-edpc-f/
 
#include <iostream>
#include <cstring>
#include <algorithm>
using namespace std;
 
int main() {
  // 入力
  string S, T;
  cin >> S >> T;
  // 文字列の長さを数値として持っておく
  int s_length = (int)S.size();
  int t_length = (int)T.size();
 
  // DP テーブル
  vector<vector <int>> dp(s_length + 1, vector<int>(t_length + 1, 0));
  
  for (int i = 0; i < s_length; ++i) {
    for (int j = 0; j < t_length; ++j) {
      // ※注： Sのi文字目までとTのj文字目までの間の部分文字列の最長はdp[i+1][j+1]
      if (S[i] == T[j]) {
        // Sのi文字目とTのj文字目が同じの場合、部分文字列の長さが1増える
        dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j] + 1);
        
      }
      // 上記とdp[i+1][j]とdp[i][j+1]のうち、最大のものがdp[i+1][j+1]で確定する
      dp[i+1][j+1] = max(dp[i+1][j+1], dp[i+1][j]);
      dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j+1]);
    }
  }
  
  // 右下(DPの最終地点)から左上(DPの開始地点)へ戻り、文字列を特定する
  string ans = "";
  int i = s_length, j = t_length;
  while(i > 0 && j > 0){
    // 左と値が同じであれば、左に戻る
    if(dp[i][j] == dp[i-1][j]) {
      i--;
    }
    // 上と値が同じであれば、上に戻る
    else if(dp[i][j] == dp[i][j-1]) {
      j--;
    }
    // どちらとも値が違う場合は、その地点の文字列を先頭に追加し左上に戻る
    else{
      ans = S[i-1] + ans;
      i--; j--;
    }
  }
 
  cout << ans << endl;
}

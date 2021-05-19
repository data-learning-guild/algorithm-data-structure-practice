"""
けんちょん本p.78に解説がある
また、AOJに問題がある（ただし螺旋本に解説はない）
https://onlinejudge.u-aizu.ac.jp/problems/DPL_1_E

Wikipediaの解説はこちら
https://ja.wikipedia.org/wiki/%E3%83%AC%E3%83%BC%E3%83%99%E3%83%B3%E3%82%B7%E3%83%A5%E3%82%BF%E3%82%A4%E3%83%B3%E8%B7%9D%E9%9B%A2
"""
from pprint import pprint


def func():
    S1 = input()
    S2 = input()
    # 編集距離は最大で1000 + 1000 = 2000なので
    INF = 2000
    # 編集前のS1の各文字を縦に、S2を横に並べるイメージ
    # dp[i][j]には、S1の最初のi文字分と、S2の最初のj文字分の間の編集距離を格納していく
    dp = [[INF] * (len(S2) + 1) for _ in range(len(S1) + 1)]
    dp[0][0] = 0

    for i in range(len(S1) + 1):
        for j in range(len(S2) + 1):
            # 次の3つのアプローチを試し、編集距離dp[i][j]が最も小さくなるように更新していく
            if i > 0 and j > 0:
                # 対応づけ操作（二次元配列dpにおける右下への矢印、何も加算しない）
                if S1[i - 1] == S2[j - 1]:
                    dp[i][j] = min(dp[i][j], dp[i - 1][j - 1])
                # 変更操作（二次元配列dpにおける右下への矢印、1を加算する）
                else:
                    dp[i][j] = min(dp[i][j], dp[i - 1][j - 1] + 1)
            # 削除操作（二次元配列dpにおける下への矢印、1を加算する）
            if i > 0:
                dp[i][j] = min(dp[i][j], dp[i - 1][j] + 1)
            # 挿入操作（二次元配列dpにおける右への矢印、1を加算する）
            if j > 0:
                dp[i][j] = min(dp[i][j], dp[i][j - 1] + 1)

    # pprint(dp)
    print(dp[-1][-1])


if __name__ == '__main__':
    func()

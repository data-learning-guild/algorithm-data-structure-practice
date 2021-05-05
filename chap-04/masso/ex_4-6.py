# 部分和問題を再帰関数で解いたコード4.9をメモ化せよ

# まずはメモ化していないコード4.9をそのままPythonで書き下す
# > https://github.com/drken1215/book_algorithm_solution/blob/master/codes/chap04/code_4_9.cpp

import sys
from typing import List


calc_cnt = 0
memo = None


def partial_sum(i: int, W: int, numbers: List[int]) -> bool:
    # show process
    global calc_cnt
    print('call i={0}, W={1}, calc:{2}'.format(i, W, calc_cnt))
    calc_cnt += 1
    
    # base case
    if i == 0:
        if W == 0:
            return True
        else:
            return False
    
    # a[i-1] を選ばない場合
    if partial_sum(i - 1, W, numbers):
        return True
    # a[i-1] を選ぶ場合
    if partial_sum(i - 1, W - numbers[i - 1], numbers):
        return True
    return False

# 解答（メモ化したもの）
# > O(NW)なので、N個の要素とW個の要素の組み合わせの結果を配列に入れるのかな？
# >> いやその場合だと計算量の削減につながらない。W=wについてTrue/Falseかわかればいいのだから、Wについて結果を保持しておけばいい
def partial_sum_memorized(i: int, W: int, numbers: List[int]) -> bool:
    # すでに計算しているならば計算しない
    global memo
    if memo[W] is not None:
        return memo[W]
    
    # show process
    global calc_cnt
    print('call i={0}, W={1}, calc:{2}'.format(i, W, calc_cnt))
    calc_cnt += 1
    
    # base case
    if i == 0:
        if W == 0:
            memo[W] = True
            return memo[W]
        else:
            memo[W] = False
            return memo[W]
    
    # a[i-1] を選ばない場合
    if partial_sum_memorized(i - 1, W, numbers):
        memo[W] = True
        return memo[W]
    # a[i-1] を選ぶ場合
    if partial_sum_memorized(i - 1, W - numbers[i - 1], numbers):
        memo[W] = True
        return memo[W]
    memo[W] = False
    return memo[W]


# 正解を参考につくったもの
# > メモ化は、memo[i][w] ← func(i, w, a) の答え とする
# > つまり当初の想定でよかった。
# > 自分の回答と異なる点は、(1)ベースケースのあとにメモチェックがきていること
# > (2)ベースケース内でメモに値を入れていないこと
# > 付け加えて、メモは「空/True/False」が表現できるようにintにしておけばよかった
# ---
# ふと思ったんだけど、「部分和問題なのに、どこで和をとってるの？」ってなったけど、
# 毎回「W-cur/W」を選択する操作が、和を取る計算だった。
# ---
# 確認してみたが、numbersが1000個ぐらいになると、答えでもRecursionErrorになる
# numbers=[1-1000],W=100にすると、メモ化ありなしの差が顕著になる
def partial_sum_memorized_solution(i: int, W: int, numbers: List[int]) -> int:
    # show process
    global calc_cnt
    print('call i={0}, W={1}, calc:{2}'.format(i, W, calc_cnt))
    calc_cnt += 1
    # base case
    if i == 0:
        if W == 0:
            return True
        else:
            return False
    global memo
    # すでに計算している場合
    if memo[i][W] != -1: return memo[i][W]
    # a[i-1] を選ばない場合
    if partial_sum_memorized_solution(i - 1, W, numbers):
        memo[i][W] = 1
        return memo[i][W]
    # a[i-1] を選ぶ場合
    if partial_sum_memorized_solution(i - 1, W - numbers[i - 1], numbers):
        memo[i][W] = 1
        return memo[i][W]
    memo[i][W] = 0
    return memo[i][W]


if __name__ == "__main__":
    numbers = list(map(int, input("numbers(sep space) >> ").split()))
    #numbers = [i for i in range(500)] # auto test
    N = len(numbers)
    W = int(input('W? >>'))
    
    # init memo 2d array
    memo = [[-1 for _ in range(W+1)] for _ in range(N+1)] # row=N+1, col=W+1
    #memo = [None for _ in range(W+1)]
    
    if partial_sum_memorized_solution(N, W, numbers) == 1:
        print('Yes')
    else:
        print('No')
    print('N=', N, ' W=', W)

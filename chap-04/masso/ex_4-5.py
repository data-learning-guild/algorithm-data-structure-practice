import sys
from typing import Tuple

# K以下の753数を数える問題。許容計算量は、O(3^d)。ただし、dはKの桁数。

# 動的計画法で解く
def count_consists_of_three_nums_dp(
    n1: int, n2: int, n3: int, K: int) -> Tuple[int, int]:
    """考えたこと
        - 制約条件は、K以下。dが3以上。これがベースケース作りに繋がると思う
        - 漸化式の元になる考え方は
            - 7/5/3からそれぞれ1つずつ選び、
            - d-3>0 ならば、d-3回、重複ありで7/5/3から数値を選ぶ
            - d桁取り終えたら、並べ替えのパターンを昇順にforループを回し
            - Kより大になったらおわり
        - わからないところ
            - 関数一回でなんの処理をすればいいか、どういう単位で再帰すればいいか
    """
    return 1, 2


# 動的計画法で解く（回答）
class ThreeNumbersCounter():
    def __init__(self) -> None:
        self.count = 0
        self.cal_cnt = 0
    
    def RecursiveCount(self, K, cur_number, used_number):
        # ベースケース
        if cur_number > K:
            return
        # 7 を含む場合はカウンターを増やす
        if used_number == 0b111:
            self.count += 1
        
        # ...
        # I don't know !

# ====================
# C++ 実装だと以下だが、
# - if (use == 0b111) ++counter
# - func(... user | 0b001 , ...)
# の部分が何をやってるのかがわからない。故にPython実装も書けない
#include <iostream>
# using namespace std;

# // N: 入力
# // cur: 現在の値
# // use: 7, 5, 3 のうちどれを使ったか
# // counter: 答え
# void func(long long N, long long cur, int use, long long &counter){
#     if (cur > N) return; // ベースケース
#     if (use == 0b111) ++counter; // 答えを増やす

#     // 7 を付け加える
#     func(N, cur * 10 + 7, use | 0b001, counter);

#     // 5 を付け加える
#     func(N, cur * 10 + 5, use | 0b010, counter);

#     // 3 を付け加える
#     func(N, cur * 10 + 3, use | 0b100, counter);
# }

# int main() {
#     long long N;
#     cin >> N;
#     long long counter = 0;
#     func(N, 0, 0, counter);
#     cout << counter << endl;
# }
# ====================


# 線形探索
def count_consists_of_three_nums(n1: int, n2: int, n3: int, K: int) -> Tuple[int, int]:
    # K is more than the least number
    sorted_numbers = sorted([n1, n2, n3])
    the_least_number = int(''.join(map(str, sorted_numbers)))
    if K < the_least_number:
        return 0, 1
    
    count = 0
    cal_cnt = 0
    for k in range(the_least_number, K+1):
        digits = [int(d) for d in str(k)]
        found_n1 = 0
        found_n2 = 0
        found_n3 = 0
        found_other = False
        for d in digits:
            cal_cnt += 1
            if n1 == d:
                found_n1 += 1
            elif n2 == d:
                found_n2 += 1
            elif n3 == d:
                found_n3 += 1
            else:
                found_other = True
                break
        if (found_other == False) and (found_n1 > 0) and (found_n2 > 0) and (found_n3 > 0):
            count += 1
    return count, cal_cnt


K = int(input('K? >>'))
n1 = int(input('n1? >>'))
n2 = int(input('n2? >>'))
n3 = int(input('n3? >>'))
result, cal_cnt = count_consists_of_three_nums(n1, n2, n3, K)
print('==============')
print('ans : ', result)
print('cal cnt : ', cal_cnt)

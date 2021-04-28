import sys

K = int(input('K? >>'))
N = int(input('N? >>'))
cal_cnt = 0
cnt = 0

# 正解を参考に
# - z は x, y が決まれば確定することを利用
# - x, y の取得は、min(K, N)にできる
cand_li = list(range(0, min(N+1, K+1)))
ans_li = []

for x in cand_li:
    for y in cand_li:
        cal_cnt += 1
        z = N - x - y
        if (0 <= z) and (z <= N):
            cnt += 1
            ans_li.append('{0}, {1}, {2}'.format(x, y ,z))

"""
■ 全探索解法しか思いつかなかった
cand_li = list(range(0, K+1))
ans_li = []

for X in cand_li:
    for Y in cand_li:
        for Z in cand_li:
            cal_cnt += 1
            if X + Y + Z == N:
                cnt += 1
                ans_li.append('{0}, {1}, {2}'.format(X, Y , Z))
"""

print('ans: ', cnt)
print('ans list: ', ans_li)
print('N : ', N, ' K : ', K, ' M(min(N,K)): ', min(N, K))
print('calc count : ', cal_cnt)
import sys

a = list(map(int,input().split()))
v = int(input())
cal_cnt = 0

cnt = 0
for av in a:
    cal_cnt += 1
    if v == av:
        cnt += 1

print('ans: ', cnt)
print('input list len: ', len(a))
print('calc count : ', cal_cnt)

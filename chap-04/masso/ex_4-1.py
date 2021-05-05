import sys

def tribo(N: int):
    print('call func : ', N)
    
    # base case
    if N == 0 or N == 1:
        return 0
    if N == 2:
        return 1
    
    # recursive proc
    result = tribo(N-1) + tribo(N-2) + tribo(N-3)
    print('calcurated : ', result)
    return result

cal_cnt = 0
N = int(input('N? >>'))
result = tribo(N)
print('==============')
print('ans : ', result)

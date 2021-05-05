import sys


class Calcurator:
    def __init__(self, N: int) -> None:
        self.initialize(N)
        
    def initialize(self, N: int) -> None:
        self.memo = [-1] * N
        self.cal_cnt = 0
        
    def tribo(self, N: int) -> int:
        print('call func : ', N)
        self.cal_cnt += 1
        
        # base case
        if N == 0 or N == 1:
            return 0
        if N == 2:
            return 1
        
        # recursive proc with memo
        if self.memo[N-1] == -1:
            self.memo[N-1] = self.tribo(N-1) + self.tribo(N-2) + self.tribo(N-3)
            print('calcurated : ', self.memo[N-1])
        result = self.memo[N-1]
        return result


N = int(input('N? >>'))
c = Calcurator(N)
result = c.tribo(N)
print('==============')
print('ans : ', result)
print('cal cnt : ', c.cal_cnt)

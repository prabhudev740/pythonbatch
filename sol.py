
def solution(n: int) -> int :
    n += 1

    while True:            
        s = str(n) # 99 [4 4 4 3 2]
        flag = True
        for i in range(len(s)-1):
            if s[i] == s[i+1]:
                flag = False
                break

        if flag:
            return n
        n += 1
    return 0

print(solution(int(input())))
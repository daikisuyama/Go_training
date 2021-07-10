n,k=map(int,input().split())
a=[[j,i] for i,j in enumerate(list(map(int,input().split())))]
a.sort()
ans=[k//n]*n
for i in range(k%n):
    ans[a[i][1]]+=1
for i in range(n):
    print(ans[i])
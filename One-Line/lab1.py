print ["VALID","INVALID"][(lambda s:reduce(lambda x,i:x+[s[i],[2*s[i],2*s[i]-9][s[i]>4]][i%2],range(len(s)),0))(map(int,raw_input()[::-1]))%10!=0]

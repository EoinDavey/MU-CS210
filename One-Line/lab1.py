print ["VALID","INVALID"][(lambda s: sum(s+[x-(x/5)*9 for x in s[1::2]]))(map(int,raw_input()[::-1]))%10!=0]

print ["VALID","INVALID"][reduce(lambda (x,b),c:(x+[c,[2*c,2*c-9][c>4]][b],not b),map(int,raw_input()[::-1]),(0,False))[0]%10!=0]

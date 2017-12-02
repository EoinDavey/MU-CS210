print (lambda xs: max(reduce(lambda x, n:x[:-1]+[(x[-1][0]+1,n)] if x[-1][1]==n else x+[(1,n)],xs[1:],[(1,xs[0])]))[1])(sorted([input() for _ in range(input())]))

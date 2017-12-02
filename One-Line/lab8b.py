print (lambda n:n[0] if n!=[] else "empty")(reduce(lambda x,y:[sorted(x+[y[7:]],key=lambda j:(len(j),j)),[[],x[1:]][x!=[]]][y=="REMOVE"],[raw_input() for _ in range(input())],[]))

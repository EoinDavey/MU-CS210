print (lambda xs:min(xs, key=lambda x: abs(x-(sum(xs)/(1.0*len(xs))))))([input() for _ in range(input())])

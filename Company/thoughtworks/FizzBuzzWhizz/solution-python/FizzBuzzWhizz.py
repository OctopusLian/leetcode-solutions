def FizzBuzzWhizz(args):
    """args[0] = Fizz, Buzz, Whizz
        args[1]= 3, 5, 7"""
    def FBW(Number):
        return Number%args[1] and Number or args[0]
    return FBW

def sayWhat(l_sayWhat,Number):
    return l_sayWhat.count(Number)<3 and "".join([s for s in l_sayWhat if type(s) is str]) or Number

def zmap(func,seq):  
    mapped_seq = []  
    for eachItem in func:  
        mapped_seq.append(eachItem(seq))  
    return mapped_seq  

def even_filter(nums, rule):
    for num in range(1,nums):
        yield sayWhat(zmap(map(FizzBuzzWhizz, rule), num),num)


rule = [("Fizz",3),("Buzz", 5),("Whizz",7)]
count = 101

for even in even_filter(count,rule):
    print even


fiz = lambda a,b,c,d:['Fizz'*(x%a==0)+'Buzz'*(x%b==0)+'Whizz'*(x%c==0) or x for x in range(1,d)]
print fiz(3,5,7,101)
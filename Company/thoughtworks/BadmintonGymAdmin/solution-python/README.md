# Solution  

## Python  

```py
import numpy as np
import datetime
def orderInput():
    order = "0"
    orderLists = []
    while True:
        order = input()
        if order == '':
            break
 
        if order[4] != ' ':
            print("Error: the booking is invalid!")
            continue
 
        orderList = order.split(" ")
        if len(orderList)==4:
            orderList.append("D")
 
        ## 设置为默认不打折
        orderList.append('10')
 
        ## 判断时间是否合法
        orderTime = orderList[2].split("~")
        orderList[2] = orderTime
        ST1 = orderTime[0]
        ET1 = orderTime[1]
        st = ST1.split(":")
        et = ET1.split(":")
        if (st[1] != "00" or et[1] != "00" or int(st[0])22 or int(et[0])>22 or int(st[0]) >= int(et[0])):
            print("Error: the booking is invalid!")
            continue
 
        # 判断场馆是否合法
        orderPlace = orderList[3]
        if (orderPlace not in ("A", "B", "C", "D")):
            print("Error: the booking is invalid!")
            continue
 
        cancelFlag = orderList[4]
        timeConFlag = False
        MatchFlag = False
        ## 判断时间是否冲突
        if (cancelFlag != "C"):
            for items in orderLists:
                if items[4] == "C":
                    continue
                T2 = items[2]
                ST2 = T2[0]
                ET2 = T2[1]
                st2 = ST2.split(":")
                et2 = ET2.split(":")
                timeConFlag = ( isTimeConflict(st, et, st2, et2) and isEqual(orderList[1], items[1]) and isEqual(orderList[3], items[3]) )
                if timeConFlag:
                    print("Error: the booking conflicts with existing bookings!")
                    break
        ## 判断是否取消订单
        else:
            for items in orderLists:
                if ( isEqual(items[0],orderList[0]) and isEqual(items[1],orderList[1]) and isEqual(items[2],orderList[2]) and isEqual(items[3],orderList[3]) and items[4]!="C"):
                    items[4] = "C"
                    print("Success: the booking is accepted!")
                    MatchFlag = True
                    break
            if not MatchFlag:
                print("Error: the booking being cancelled does not exist!")
        ## 判断是否打折
        lines = readDiscount()
        for line in lines:
            if (line[0:10] = orderList[1]):
                orderList[5] = line[22]
 
        if (not timeConFlag) and (cancelFlag != "C"):   
            orderLists.append(orderList) 
            print("Success: the booking is accepted!")
    return orderLists
def readDiscount():
    f = open("./discount.txt","r")
    lines = f.readlines()
    return lines
def isEqual(v1, v2):
    if v1 == v2:
        return True
    else:
        return False
 
def isTimeConflict(st1, et1, st2, et2):
    if not (et1<=st2 or et2<=st1) :
        return True
    else:
        return False
 
def dayOfWeek(date):
    date = date.split("-")
    year = int(date[0])
    month = int(date[1])
    day = int(date[2])
    anyday=datetime.datetime(year,month,day).strftime("%w")
    return anyday
def calFare(date,time,cancelFlag,discount):
    ST = time[0].split(":")
    ET = time[1].split(":")
    st = int(ST[0]) - 9
    et = int(ET[0]) - 9
    fareListofWD = np.array([30,30,30,50,50,50,50,50,50,80,80,60,60])
    fareListofWE = np.array([40,40,40,50,50,50,50,50,50,60,60,60,60])
    day = dayOfWeek(date)
    if day in ('6', '7'):
        fare = sum(fareListofWE[st:et]) * int(discount)/10
        cutoff = sum(fareListofWE[st:et]) - fare
        if cancelFlag == 'C':
            fare = fare/4
    else:
        fare = sum(fareListofWD[st:et]) * int(discount)/10
        cutoff = sum(fareListofWD[st:et]) - fare
        if cancelFlag == 'C':
            fare = fare/2
    return int(fare), int(cutoff)
def printFareP(orderLists, place):
    groundFare = 0
    orderListsP, countP = findPlaceOrder(orderLists, place)
    idxP = sortOrderLists(orderListsP)
    for index in idxP:
        items = orderListsP[index]
        timeT = items[2]
        if items[4] == 'C':
            fare, cutoff = calFare(items[1],items[2],items[4],items[5])
            #print(items[1],timeT[0],'~',timeT[1],"违约金",fare,"元")
            print('{} {}{}{} {} {}{}'.format(items[1],timeT[0],'~',timeT[1],"违约金",fare,"元"))
        else:
            fare, cutoff = calFare(items[1],items[2],items[4],items[5])
            if items[5] == '10':
                print('{} {}{}{} {}{}'.format(items[1],timeT[0],'~',timeT[1],fare,"元"))
                #print(items[1],timeT[0],'~',timeT[1],fare,"元")
            else:
                print('{} {}{}{} {}{} {}{}{}'.format(items[1],timeT[0],'~',timeT[1],fare,"元","已优惠:",cutoff,"元"))
                #print(items[1],timeT[0],'~',timeT[1],fare,"元","已优惠:",cutoff,"元")
        groundFare += fare
    return groundFare
 
def printFare(orderLists):
    totalFare = 0   
    print('收入汇总')
    print('---')   
    print('场地:A')
    groundFare = printFareP(orderLists, 'A')
    totalFare += groundFare
    print('{} {}{}'.format('小计:', groundFare, '元'))
    #print('小计：', groundFare, '元')
    print('\n')
    print('场地:B')
    groundFare = printFareP(orderLists, 'B')
    totalFare += groundFare
    print('{} {}{}'.format('小计:', groundFare, '元'))
    print('\n')
    print('场地:C')
    groundFare = printFareP(orderLists, 'C')
    totalFare += groundFare
    print('{} {}{}'.format('小计:', groundFare, '元'))
    print('\n')
    print('场地:D')
    groundFare = printFareP(orderLists, 'D')
    totalFare += groundFare
    print('{} {}{}'.format('小计:', groundFare, '元'))
    print('---')
    #print('总计：', totalFare, '元')
    print('{} {}{}'.format('总计:', totalFare, '元'))
def findPlaceOrder(orderLists, place):
    orderListsP = []
    countP = 0
    for items in orderLists:
        if items[3] == place:
            orderListsP.append(items)
            countP += 1
    return orderListsP, countP
 
def sortOrderLists(orderListsP):
    arrayP = []
    index = 0
    for items in orderListsP:
        tempD = items[1]
        tempT = items[2]
        tempST = tempT[0]
        strP = tempD[0:4] + tempD[5:7] + tempD[8:10] + tempST[0:2]
        arrayP.append(int(strP))
        index =  index+1
    #npList = np.array(orderListsP)
    idx = np.argsort(arrayP[:], axis = 0)
    return idx
 
if __name__ == '__main__':
 
    # 测试
    # date = "2018-07-25"      
    # print(dayOfWeek(date))
 
    orderLists = orderInput()
    printFare(orderLists)
```
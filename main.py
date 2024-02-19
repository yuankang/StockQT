# -*- coding: utf-8 -*-

import requests
from time import sleep
from datetime import datetime, time
from dateutil import parser

# fiducial value/ base value/ reference value
# 开盘价 收盘价 最高价 最低价, 基准值: 20天内收盘价的平均值
# 最好是用本地数据来计算基准值, 基准值也要动态变化, 至少每天重算一次
def getBaseValue():
    return 40

# sh600036 招商, sh600519 茅台
def getTick():
    page = requests.get("http://qt.gtimg.cn/q=sh600036")
    stock = page.text
    sList = stock.split('~')
    
    name = sList[1]
    code = sList[2]
    bPrice = float(sList[5])
    ePrice = float(sList[4])
    dt = sList[30]
    # dt0 = parser.parse(info[4]).time()
    info = (name, code, bPrice, ePrice, dt)
    return info

def buy():
    print("buy")

def sell():
    print("sell")

# 确定基准值, 每5分钟更新一次, 低于基准值5%买入, 高于基准值5%卖出
def strategy(ma20):
    low = (1-0.05)*ma20
    high = (1+0.05)*ma20
    print(ma20, low, high)
    
    while True:
        info = getTick()
        print(info)
        if info[2] < low:
            sell()
        elif info[2] > high:
            buy()
        else:
            print("---")
        sleep(5)

bv = getBaseValue()
strategy(bv)
    
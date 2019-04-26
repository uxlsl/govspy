#!/usr/bin/env python
# -*-coding:utf-8-*-
# File Name     : Structs.py
# Description   :
# Author        :
# Creation Date : 2019-04-26
# Last Modified : Fri 26 Apr 2019 11:03:07 AM CST
# Created By    : lsl

from math import sqrt


class Point(object):
    def __init__(self, x, y):
        self.x = x
        self.y = y


def distance(point1, point2):
    return sqrt(point1.x * point2.x + point1.y * point2.y)


p1 = Point(1, 3)
p2 = Point(2, 4)
print(distance(p1, p2))  # 3.74165738677

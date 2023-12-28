import pandas as pd
import sys
import random
import itertools

with open("inputs/3.txt", "rb") as f:
    two_by_two = [bytearray(b.rstrip()) for b in f.readlines()]
    single_arr = []
    single_arr.extend(b for b in two_by_two)

dot = 46
numbers = list(ord(str(b)) for b in range(0, 10))

accounted = []

def account(arr):
    for i in range(0, len(arr)):
        prev = i - 1
        next = i + 1

        try:
            if arr[prev] == 46 and arr[next] == 46:
                continue
            else:
                accounted.append(i)
        except IndexError:
            continue

def account_diag(arr):
    # look at all diags
    # if it checks add the index of curr
    diags = [(-1,-1), (-1,1), (1, -1), (1, 1)]
    get_all_diags = lambda x,y : [(x+x1, y+y1) for x1,y1 in diags]

    for i in range(0, len(arr)):
        for j in range(0, len(arr)):
            try:
                if all([arr[x][y] != 46 for x,y in get_all_diags(i,j)]):
                   accounted.append(i * 140 + j)
            except IndexError:
                continue

single_arr_t = two_by_two
for i in range(0, len(two_by_two)):
    for k in range(0, len(two_by_two)):
        single_arr_t[k][i] = two_by_two[i][k]

print(f"DEBUGPRINT[8]: 2_1.py:50: single_arr_t={single_arr_t[0]}", file=sys.stderr)
print(f"DEBUGPRINT[9]: 2_1.py:49: two_by_two[0]={two_by_two[0]}", file=sys.stderr)

# account(single_arr)
# account_diag(two_by_two)
# print(accounted)

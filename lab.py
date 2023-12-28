import sys

a = list(range(0,100))

def two_i_plus_j():
    printed = []
    vals = []

    for i in range(0,4):
        for j in range(0,25):
            idx = i * 25 + j
            curr = str(idx)
            printed.append(curr)
            vals.append(a[idx])

            if len(curr) == 1:
                printed.append("   ")
            else:
                printed.append("  ")
        if i != 3:
            printed.append("\n")
    printed = "".join(printed)
    print(f"{printed}")
    print(vals)

two_i_plus_j()

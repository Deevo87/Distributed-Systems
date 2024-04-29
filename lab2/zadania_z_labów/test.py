def binary_search(A, low, high, k):
    if high >= low:
        mid = (low + high) // 2
        if A[mid] is None:
            return binary_search(A, low, mid-1, k)
        if k == A[mid]:
            return mid
        elif k < A[mid]:
            return binary_search(A, low, mid-1, k)
        else:
            return binary_search(A, mid+1, high, k)
    else:
        return high, low

if __name__ == '__main__':
    arr = [1,3,5,6]
    tab = [None] * 5
    x = 2
    n = len(arr)
    print(binary_search(arr, 0, n-1, x))
    # print(binary_search(tab, 0, len(tab)-1, 2))

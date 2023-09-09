def quicksort(arr,l,h):
	if l<h:
		p=partition(arr,l,h)
		quicksort(arr,l,p-1)
		quicksort(arr,p+1,h)

def partition(arr,l,h):
    print("l,h",l,h,arr)
    pivot=arr[h]
    indx=l-1
    for i in range(l,h):
        print(indx)
        if arr[i]<pivot:
            indx+=1
            arr[indx],arr[i]=arr[i],arr[indx]
    print(indx)

    arr[indx+1],arr[h] = arr[h],arr[indx+1]
    return indx+1

if __name__ == '__main__':
    arr=[10,7,8,9,1,5]
    n=len(arr)
    quicksort(arr,0,n-1)
    print("Sorted array is:")
    for i in range(n):
        print("%d" %arr[i])
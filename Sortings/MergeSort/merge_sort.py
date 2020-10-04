def merge_sort(array):
	n = len(array)

	if n == 1:
		return array

	a = array[: int(n/2)] 
	b = array[int(n/2): ]	

	a = merge_sort(a)
	b = merge_sort(b)

	return merge(a, b)

def merge(a, b):
	list_3 = []

	while a and b:
		list_3.append(b.pop(0)) if a[0] > b[0] else list_3.append(a.pop(0))
	
	for i in a:
		list_3.append(i)	
	for j in b:
		list_3.append(j)

	return list_3
	

array = [100, 2, 24, 23, 1, 0, 2]	
print(merge_sort(array))
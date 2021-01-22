from segment_tree import SegmentTree 

arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11] 

t = SegmentTree(arr) 

a = t.query(2, 9, "max") 
print("The maximum value of this range is : ", a) 


a = t.query(2, 9, "min") 
print("The minimum value of this range is : ", a) 

a = t.query(2, 7, "sum") 
print("The sum of this range is : ", a) 

t.update(2, 25) 

print("The updated array is : ", arr) 

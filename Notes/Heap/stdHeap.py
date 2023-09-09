import heapq 

nums = [5,3,17,10,84,19,6]
heapq.heapify(nums)
print(nums)
heapq.heappush(nums, 2)
print(nums)
heapq.heappop(nums)
print(nums)


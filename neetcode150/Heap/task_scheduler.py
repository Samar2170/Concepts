from typing import List
import heapq
from collections import Counter, deque
class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        task_count = Counter(tasks)
        maxHeap = [-cnt for cnt in task_count.values()]
        heapq.heapify(maxHeap)

        time=0
        q=deque()
        while maxHeap or q:
            time+=1
            if maxHeap:
                cnt = 1+ heapq.heappop(maxHeap)
                if cnt:
                    q.append((cnt,time+n))

            if q and q[0][1]==time:
                ##why ? if its time again pop from queue and append to maxHeap again
                heapq.heappush(maxHeap, q.popleft()[0])
        return time

            
### iterate while maxHeap or q:  -> get count variable , add 1 -> if count is non zero then append to q as tuple( count, time+n) 
###                             -> if q[0][1] ==time -> popleft from queue and heappush to maxHeap 
### time is calculated in this loop.

if __name__ == '__main__':
    tasks = ["A","A","A","B","B","B","C"]
    n = 2
    print(Solution().leastInterval(tasks,n))
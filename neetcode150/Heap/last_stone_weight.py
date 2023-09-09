from typing import List
import heapq

class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        stones=[-1*s for s in stones]
        heapq.heapify(stones)
        
        while len(stones)>1:
            s1 = heapq.heappop(stones)
            s2 = heapq.heappop(stones)
            rem=s1-s2
            heapq.heappush(stones,rem)

        return -1*stones[0]

        
if __name__ == '__main__':
    stones = [2,7,4,1,8,1]
    print(Solution().lastStoneWeight(stones))

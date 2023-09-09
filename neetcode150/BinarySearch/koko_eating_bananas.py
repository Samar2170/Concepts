from typing import List
class Solution:
    def minEatingSpeed(self, piles: List[int], H: int) -> int:
        def canEatAll(piles, H, K):
            return sum((p-1)//K+1 for p in piles) <= H

        l, r = 1, max(piles)
        while l < r:
            m = (l+r)//2
            if canEatAll(piles, H, m):
                r = m
            else:
                l = m+1
        return l


if __name__ == "__main__":
    piles = [3,6,7,11]
    H = 8
    print(Solution().minEatingSpeed(piles, H))
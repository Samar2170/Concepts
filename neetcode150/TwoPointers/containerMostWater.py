class Solution:
    def maxArea(self,height) -> int:
        n=len(height)
        i,k=0,n-1
        max_area = 0
        while i<k:
            areaHere = min(height[i], height[k]) * (k-i)
            if areaHere>max_area:
                max_area=areaHere
            ## premise is to move the pointer with smaller height
            if height[i]>height[k]:
                k-=1
            else:
                i+=1

        return max_area

if __name__ == "__main__":
    s = Solution()
    print(s.maxArea([1,8,6,2,5,4,8,3,7]))
    print(s.maxArea([2,3,4,5,18,17,6]))
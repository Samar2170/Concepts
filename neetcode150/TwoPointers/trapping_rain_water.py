class Solution:
    def trap(self,height)-> int:
        n=len(height)
        if n<3:
            return 0
        i,k=0,n-1
        max_area = 0
        maxleft = [0] * n
        maxright = [0]  * n

        maxleft[0]=height[0]
        maxright[k]=height[k]
        for i in range(1,n):
            maxleft[i] = max(height[i], maxleft[i-1])
        
        for j in range(k-1,-1,-1):
            maxright[j] = max(height[j], maxright[j+1])

        totalWater=0
        print(maxleft)
        print(maxright)
        for i in range(1,k):
            water= min(maxright[i+1], maxleft[i-1]) - height[i]
            print(i,water)

            if water>0:
                totalWater+=water 

        return totalWater

if __name__ == "__main__":
    s = Solution()
    # print(s.trap([0,1,0,2,1,0,1,3,2,1,2,1]))
    print(s.trap([4,2,0,3,2,5]))

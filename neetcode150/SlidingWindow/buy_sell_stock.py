class Solution:
    def maxProfit(self, prices):
        minPrice=prices[0]
        maxProfit=0

        for i in range(1,len(prices)):
            minPrice = min(minPrice,prices[i])
            profit = prices[i]-minPrice
            maxProfit = max(maxProfit,profit)
        return maxProfit



if __name__=='__main__':
    s = Solution()
    print(s.maxProfit([7,1,5,3,6,4]))
    print(s.maxProfit([7,6,4,3,1]))
    print(s.maxProfit([2,4,1]))
    print(s.maxProfit([1,2]))
    print(s.maxProfit([2,1,2,0,1]))
from typing import List
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        ans = [0] * len(temperatures)

        for i in range(len(temperatures)):
            while stack and temperatures[i] > temperatures[stack[-1]]:
                index = stack.pop()
                ans[index] = i - index
            stack.append(i)
        return ans



if __name__ == "__main__":
    s = Solution()
    print(s.dailyTemperatures([73,74,75,71,69,72,76,73]))
    print(s.dailyTemperatures([30,40,50,60]))
    print(s.dailyTemperatures([30,60,90]))



### iterate over temps  -> if stack exists compare current temp to stack[-1] temp, if current temp is greater than stack[-1] temp, then pop stack[-1] and calculate the difference between current index and stack[-1] index, 
# else append current index to stack
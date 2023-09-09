# Backtracking
## Brute Force 
### Not for Optimization problem




### 1. Permutations
#### 1.1. Permutations

```python
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        self.dfs(nums, [], res)
        return res
    
    def dfs(self, nums, path, res):
        if not nums:
            res.append(path)
            return
        for i in range(len(nums)):
            self.dfs(nums[:i] + nums[i+1:], path + [nums[i]], res)
```

#### 1.2. Permutations II
```python
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()
        self.dfs(nums, [], res)
        return res
    
    def dfs(self, nums, path, res):
        if not nums:
            res.append(path)
            return
        for i in range(len(nums)):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            self.dfs(nums[:i] + nums[i+1:], path + [nums[i]], res)
```

#### 1.3. Combinations
```python
class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []
        self.dfs(1, n, k, [], res)
        return res
    
    def dfs(self, start, n, k, path, res):
        if k == 0:
            res.append(path)
            return
        for i in range(start, n+1):
            self.dfs(i+1, n, k-1, path + [i], res)
```


#### 1.4. Subsets
```python
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        self.dfs(nums, 0, [], res)
        return res
    
    def dfs(self, nums, index, path, res):
        res.append(path)
        for i in range(index, len(nums)):
            self.dfs(nums, i+1, path + [nums[i]], res)
```

#### 1.5. Subsets II
```python
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()
        self.dfs(nums, 0, [], res)
        return res
    
    def dfs(self, nums, index, path, res):
        res.append(path)
        for i in range(index, len(nums)):
            if i > index and nums[i] == nums[i-1]:
                continue
            self.dfs(nums, i+1, path + [nums[i]], res)
```

#### 1.6. Palindrome Partitioning
```python
class Solution:
    def partition(self, s: str) -> List[List[str]]:
        res = []
        self.dfs(s, [], res)
        return res
    
    def dfs(self, s, path, res):
        if not s:
            res.append(path)
            return
        for i in range(1, len(s)+1):
            if self.isPalindrome(s[:i]):
                self.dfs(s[i:], path + [s[:i]], res)
    
    def isPalindrome(self, s):
        return s == s[::-1]
```

#### 1.7. Palindrome Partitioning II
```python
class Solution:
    def minCut(self, s: str) -> int:
        n = len(s)
        dp = [i for i in range(n)]
        for i in range(n):
            self.helper(s, i, i, dp)
            self.helper(s, i, i+1, dp)
        return dp[-1]
    
    def helper(self, s, left, right, dp):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            if left == 0:
                dp[right] = 0
            else:
                dp[right] = min(dp[right], dp[left-1] + 1)
            left -= 1
            right += 1
```


#### 1.8. Letter Combinations of a Phone Number
```python
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        res = []
        self.dfs(digits, 0, "", res)
        return res
    
    def dfs(self, digits, index, path, res):
        if len(path) == len(digits):
            res.append(path)
            return
        for i in range(index, len(digits)):
            for c in self.getLetters(digits[i]):
                self.dfs(digits, i+1, path + c, res)
    
    def getLetters(self, digit):
        if digit == "2":
            return "abc"
        elif digit == "3":
            return "def"
        elif digit == "4":
            return "ghi"
        elif digit == "5":
            return "jkl"
        elif digit == "6":
            return "mno"
        elif digit == "7":
            return "pqrs"
        elif digit == "8":
            return "tuv"
        elif digit == "9":
            return "wxyz"
```

#### 1.9. Generate Parentheses
```python
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        self.dfs(n, n, "", res)
        return res
    
    def dfs(self, left, right, path, res):
        if left > right:
            return
        if left == 0 and right == 0:
            res.append(path)
            return
        if left > 0:
            self.dfs(left-1, right, path + "(", res)
        if right > 0:
            self.dfs(left, right-1, path + ")", res)
```

#### 1.10. Remove Invalid Parentheses
```python

class Solution:
    def removeInvalidParentheses(self, s: str) -> List[str]:
        res = []
        self.dfs(s, 0, 0, "(", ")", res)
        return res
    
    def dfs(self, s, last_i, last_j, open_p, close_p, res):
        stack = 0
        for i in range(last_i, len(s)):
            if s[i] == open_p:
                stack += 1
            if s[i] == close_p:
                stack -= 1
            if stack >= 0:
                continue
            for j in range(last_j, i+1):
                if s[j] == close_p and (j == last_j or s[j-1] != close_p):
                    self.dfs(s[:j] + s[j+1:], i, j, open_p, close_p, res)
            return
        reversed_s = s[::-1]
        if open_p == "(":
            self.dfs(reversed_s, 0, 0, ")", "(", res)
        else:
            res.append(reversed_s)
```


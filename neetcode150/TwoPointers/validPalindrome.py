import re
class Solution:
    def isPalindrome(self,s):
        """
        :type s: str
        :rtype: bool
        """
        # Remove all non-alphanumeric characters
        s = re.sub('[^a-zA-Z0-9]','',s)
        # Convert to lowercase
        s = s.lower()
        # Check if it is a palindrome
        return s==s[::-1]


if __name__ == "__main__":
    s = Solution()
    print(s.isPalindrome("A man, a plan, a canal: Panama"))
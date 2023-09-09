class Solution:
    def encode(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        return ''.join(s.replace('|', '||') + '|' for s in strs)

    def decode(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        return [t.replace('||', '|') for t in s.split('|')[:-1]]


if __name__ == "__main__":
    s = Solution()
    print(s.encode(["Hello", "World"]))
    print(s.decode("Hello|World|"))
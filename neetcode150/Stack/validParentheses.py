class Solution:
    def isValid(self, s: str) -> bool:
        stack=[]

        for i in range(len(s)):
            if s[i]=='(' or s[i]=='{' or s[i]=='[':
                stack.append(s[i])
            elif s[i]==')':
                if len(stack)==0 or stack[-1]!='(':
                    return False
                else:
                    stack.pop()
            elif s[i]=='}':
                if len(stack)==0 or stack[-1]!='{':
                    return False
                else:
                    stack.pop()
            elif s[i]==']':
                if len(stack)==0 or stack[-1]!='[':
                    return False
                else:
                    stack.pop()
        return True




if __name__=='__main__':
    sol = Solution()
    print(sol.isValid("()"))
    print(sol.isValid("()[]{}"))
    print(sol.isValid("(]"))
    print(sol.isValid("([)]"))
    print(sol.isValid("{[]}"))       

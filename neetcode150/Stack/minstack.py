class MinStack:
    def __init__(self):
        self.stack = []
        self.min = []
    def push(self,x):
        if len(self.stack)==0:
            self.min.append(x)
        else:
            self.min.append(min(self.min[-1],x))
        self.stack.append(x)

    def pop(self):
        self.stack.pop()
        self.min.pop()
    def top(self):
        return self.stack[-1]
    def getMin(self):
        return self.min[-1]


if __name__=='__main__':
    minstack = MinStack()
    minstack.push(-2)
    minstack.push(0)
    minstack.push(-3)
    print(minstack.getMin())
    minstack.pop()
    print(minstack.top())
    print(minstack.getMin())
class MinHeap:
	def __init__(self,maxsize=15):
		self.size=maxsize
		self.front=1
		self.size=0 
		self.heap = [0] * (maxsize+1)
	def leftChild(self,pos):
		return 2*pos 

	def rightChild(self,pos):
		return 2*pos + 1

	def parent(self,pos):
		return pos // 2

	def isLeaf(self,pos):
		if not 2*pos > self.size and self.heap[2*pos]!=0:
			return True
		return False

	def swap(self,fpos,spos):
		self.heap[fpos],self.heap[spos]=self.heap[spos], self.heap[fpos]

	def insert(self,element):
		self.size+=1
		self.heap[self.size]=element 
		current=self.size
		while self.heap[current] < self.heap[self.parent(current)]:
			self.swap(current,self.parent(current))
			current=self.parent(current)

	def minHeapify(self,pos):
		if not self.isLeaf(pos):
			if self.heap[pos]>self.heap[self.leftChild(pos)] or self.heap[pos]>self.heap[self.rightChild(pos)]:
				if self.heap[pos]>self.heap[self.leftChild(pos)]:
					self.swap(pos,self.leftChild(pos))
					self.minHeapify(self.leftChild(pos))
				else:
					self.swap(pos,self.rightChild(pos))
					self.minHeapify(self.rightChild(pos))

	def extractMin(self):
		popped = self.heap[self.front]
		self.heap[self.front] = self.heap[self.size]
		self.size-=1
		self.minHeapify(self.front)
		return popped



if __name__ == "__main__":
    minHeap=MinHeap()
    minHeap.insert(5)
    minHeap.insert(3)
    minHeap.insert(17)
    minHeap.insert(10)
    minHeap.insert(84)
    minHeap.insert(19)
    minHeap.insert(6)
    print(minHeap.heap)



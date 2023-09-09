from typing import List
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        l,h=0,len(matrix)-1
        rmatrix = []
        while l<=h:
            mid = (l+h)//2
            if matrix[mid][0] <= target <= matrix[mid][-1]:
                rmatrix = matrix[mid]
                break
            elif matrix[mid][0] > target:
                h = mid-1
            else:
                l = mid+1
        l,h=0,len(rmatrix)-1
        while l<=h:
            mid = (l+h)//2
            if rmatrix[mid] == target:
                return True
            elif rmatrix[mid] < target:
                l = mid+1
            else:
                h = mid-1
        return False



if __name__ == '__main__':
    s = Solution()
    print(s.searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]],3))
    print(s.searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]],11))
    print(s.searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]],60))
    print(s.searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]],0))
    print(s.searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]],61))
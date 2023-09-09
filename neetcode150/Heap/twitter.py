from collections import defaultdict, deque
import itertools
from typing import List
import heapq


class Twitter:

    def __init__(self):
        self.follows = defaultdict(set)
        self.tweets = defaultdict(deque)
        self.timer = itertools.count(step=-1)

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.tweets[userId].appendleft((next(self.timer), tweetId))
        

    def getNewsFeed(self, userId: int) -> List[int]:
        tweets = heapq.merge(*(self.tweets[u] for u in self.follows[userId] | {userId}))
        return [t for _, t in itertools.islice(tweets, 10)]

    def follow(self, followerId: int, followeeId: int) -> None:
        self.follows[followerId].add(followeeId)
        

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.follows[followerId].discard(followeeId)
        

if __name__ == '__main__':
    # import ipdb; ipdb.set_trace()
    twitter = Twitter()
    twitter.postTweet(1, 5)
    twitter.getNewsFeed(1)
    twitter.follow(1, 2)
    twitter.postTweet(2, 6)
    twitter.getNewsFeed(1)
    twitter.unfollow(1, 2)
    twitter.getNewsFeed(1)
    twitter.postTweet(1, 7)
    

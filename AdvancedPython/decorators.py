import time

def timer(func):
    def inner_func(*args, **kwargs):
        t0=time.time()
        func(*args,**kwargs)
        print("Time taken: ", time.time()-t0)
    return inner_func


@timer
def test_func():
    for i in range(1000):
        pass

test_func()
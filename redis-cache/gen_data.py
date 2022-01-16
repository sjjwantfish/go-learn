import redis

r = redis.StrictRedis(host='localhost', port=6379, db=0)
for i in range(10000):
    r.set(str(i), "RANDOMPASS" * 500)
    # r.delete(str(i))


sample = {'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5}

# Sort by key
print(sorted(sample.items(), key=lambda x:x[0]))


# Sort by value
print(sorted(sample.items(), key=lambda x:x[1]))

# Sort by value, descending
print(sorted(sample.items(), key=lambda x:x[1], reverse=True))

# Sort by key, descending
print(sorted(sample.items(), key=lambda x:x[0], reverse=True))

# Sort by length of key string
print(sorted(sample.items(), key=lambda x:len(x[0])))

# Sort by length of value string
print(sorted(sample.items(), key=lambda x:len(str(x[1]))))


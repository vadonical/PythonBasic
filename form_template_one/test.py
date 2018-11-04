def map_by_first_letter(text):
    mapped = dict()
    for line in text.split('\r\n'):
        print(line)
        for word in [x for x in line.split(' ') if len(x) > 0]:
            if word[0] not in mapped:
                mapped[word[0]] = []
            mapped[word[0]].append(word)
    return mapped


if __name__ == '__main__':
    a = 'hello,world\nhhhshhjha hsjjf\nsjkkkskkaf skfak' \
        'skafkaf k' \
        'sfkskafk'
    ret = map_by_first_letter(a)


"""
qqqqqq 
wwww wwww 
eeee eee 
rrrrrrr 
"""
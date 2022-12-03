def price(item):
    return ord(item) - 38 if item.isupper() else ord(item) - 96


def part1():
    total = 0
    for line in open('input.txt'):
        rucksack = line.strip()
        size = len(rucksack) // 2
        total += sum([price(item) for item in set(rucksack[:size]) & set(rucksack[size:])])
    print(total)


def part2():
    total = 0
    badges = []
    for line in open('input.txt'):
        badges.append(line.strip())
        if len(badges) == 6:
            badge1 = list(set(badges[0]) & set(badges[1]) & set(badges[2]))[0]
            badge2 = list(set(badges[3]) & set(badges[4]) & set(badges[5]))[0]
            total += price(badge1) + price(badge2)
            badges.clear()
    print(total)


part1()
part2()

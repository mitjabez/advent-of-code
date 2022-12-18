col_count = 7
minus = [[True] * 4]
plus = [[False, True, False], [True] * 3, [False, True, False]]
arrow = [[True] * 3, [False, False, True], [False, False, True]]
dash = [[True]] * 4
block = [[True] * 2] * 2
rocks = [minus, plus, arrow, dash, block]

chamber = [[False] * col_count for _ in range(4)]
with open('input.txt', 'r') as file:
    jet = file.read().rstrip()

max_hit_count = 1000000000000
current_rock_id = 0
rock = rocks[current_rock_id]
y = len(chamber) - 1
x = 2
i = 0
tower_size = 0
hit_count = 0
history_height_per_column = []
history_block = []
history_jet = []
history_tower_size = []
solved_parts = {1: False, 2: False}


def draw():
    for y_pos in range(len(chamber) - 1, -1 if len(chamber) < 20 else len(chamber) - 20, -1):
        print(f'{y_pos + 1:2}', end='')
        for x_pos in range(len(chamber[y])):
            if y_pos >= y and y_pos <= (y + len(rock) - 1) and x_pos >= x and x_pos <= (x + len(rock[0]) - 1):
                if rock[-(y - y_pos)][x_pos - x]:
                    print('@', end='')
                    continue
            print('#' if chamber[y_pos][x_pos] else '.', end='')
        print()
    print('--------------------')


def can_move(new_x: int, new_y: int) -> bool:
    if new_x < 0 or new_x >= col_count or new_y < 0:
        return False
    for y_pos in range(len(rock)):
        for x_pos in range(len(rock[y_pos])):
            if rock[y_pos][x_pos] and chamber[new_y + y_pos][new_x + x_pos]:
                return False
    return True


def relative_height_per_column():
    height_per_col = {i: 0 for i in range(col_count)}
    # Relative height of last 20 rows
    for y_pos in range(tower_size, tower_size - 20 if tower_size > 20 else 0, -1):
        for x_pos in range(col_count):
            if chamber[y_pos][x_pos] and height_per_col[x_pos] == 0:
                height_per_col[x_pos] = tower_size - y_pos
    return height_per_col


while False in solved_parts.values():
    new_x = x + (1 if jet[i % len(jet)] == '>' else -1)
    new_x = new_x if new_x >= 0 and new_x + len(rock[0]) - 1 < len(chamber[0]) else x
    if can_move(new_x, y):
        x = new_x
    i += 1

    new_y = y - 1
    if can_move(x, new_y):
        y = new_y
    else:
        if hit_count == 2022:
            solved_parts[1] = True
            print(f'part1: {tower_size}')

        hit_count += 1
        tower_size = max(tower_size, y + len(rock))

        for y_pos in range(len(rock)):
            for x_pos in range(len(rock[y_pos])):
                chamber[y + y_pos][x + x_pos] |= rock[y_pos][x_pos]

        current_rock_id = (current_rock_id + 1) % len(rocks)
        rock = rocks[current_rock_id]
        y = tower_size + 3
        x = 2

        lines_to_add = (y + len(rock)) - len(chamber)
        chamber.extend([[False] * col_count for _ in range(lines_to_add)])

        history_block.append(current_rock_id)
        history_jet.append(i % len(jet))
        history_height_per_column.append(relative_height_per_column())
        history_tower_size.append(tower_size)
        for p in range(hit_count - 1):
            if solved_parts[2]:
                break
            if history_block[-1] == history_block[p] and \
                    history_jet[-1] == history_jet[p] and \
                    history_height_per_column[-1] == history_height_per_column[p]:
                diff_count = len(history_block) - 1 - p
                delta_tower_size = history_tower_size[-1] - history_tower_size[p]
                trillion_tower_size = (max_hit_count - hit_count) // diff_count * delta_tower_size
                steps_till_max = (max_hit_count - hit_count) % diff_count
                remaining_tower_size = history_tower_size[p + steps_till_max] - history_tower_size[p]
                solved_parts[2] = True
                print(f'part2: {trillion_tower_size + tower_size + remaining_tower_size}')

        # draw()
        # input()

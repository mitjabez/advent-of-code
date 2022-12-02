ROCK = 'rock'
PAPER = 'paper'
SCISSORS = 'scissors'

MAPPER = {
    'A': ROCK, 'X': ROCK,
    'B': PAPER, 'Y': PAPER,
    'C': SCISSORS, 'Z': SCISSORS
}
OUTCOMES = {
    'X': 'defeats',
    'Y': 'draw',
    'Z': 'loses',
}

CONFIG = {
    ROCK: {'draw': ROCK, 'defeats': SCISSORS, 'loses': PAPER, 'score': 1},
    PAPER: {'draw': PAPER, 'defeats': ROCK, 'loses': SCISSORS, 'score': 2},
    SCISSORS: {'draw': SCISSORS, 'defeats': PAPER, 'loses': ROCK, 'score': 3}
}


def score_round(player, opponent):
    draw_score = 1 if player == opponent else 0
    win_score = 1 if CONFIG[player]['defeats'] == opponent else 0
    shape_score = CONFIG[player]['score']
    return draw_score * 3 + win_score * 6 + shape_score


def part1():
    total = 0
    for line in open('input.txt'):
        opponent, player = [MAPPER[p] for p in line.strip().split(' ')]
        total = total + score_round(player, opponent)
    print(total)


def part2():
    total = 0
    for line in open('input.txt'):
        opponent_id, outcome = line.strip().split(' ')
        opponent = MAPPER[opponent_id]
        player = CONFIG[opponent][OUTCOMES[outcome]]
        total = total + score_round(player, opponent)
    print(total)


part1()
part2()

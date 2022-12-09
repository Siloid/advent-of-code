#!/usr/bin/env python3

"""
 Values
 Rock = 1
 Paper = 2
 Scissors = 3
 Loss = 0
 Tie = 3
 Win = 6

 Inputs
 A/X = Rock
 B/Y = Paper
 C/Z = Scissors
"""


class RPS(object):
    loss = 0
    tie = 3
    win = 6
    values = {"X": 1, "Y": 2, "Z": 3}
    outcomes = {
        "A": {"X": tie, "Y": win, "Z": loss},
        "B": {"X": loss, "Y": tie, "Z": win},
        "C": {"X": win, "Y": loss, "Z": tie},
    }
    outcome_map = {"X": loss, "Y": tie, "Z": win}

    @staticmethod
    def round(p1, p2):
        outcome = RPS.outcomes[p1][p2]
        value = RPS.values[p2]
        score = outcome + value
        return score

    @staticmethod
    def round_with_predicted_outcome(p1, result):
        for throw, outcome in RPS.outcomes[p1].items():
            if RPS.outcome_map[result] == outcome:
                need_to_throw = throw
                break
        return RPS.round(p1, need_to_throw)


if __name__ == "__main__":
    game = RPS()

    # Part1
    total_score = 0
    with open("./input.txt", "r") as fp:
        for line in fp.readlines():
            p1, p2 = line.split(" ")
            total_score += game.round(p1, p2.strip())
    print(f"Part1 Answer: Total score: {total_score}")

    # Part2
    total_score = 0
    with open("./input.txt", "r") as fp:
        for line in fp.readlines():
            p1, result = line.split(" ")
            total_score += game.round_with_predicted_outcome(p1, result.strip())
    print(f"Part2 Answer: Total score: {total_score}")

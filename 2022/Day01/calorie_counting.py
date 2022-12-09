#!/usr/bin/env python3


class elf(object):
    def __init__(self):
        self.food = []

    def add_food(self, calories):
        self.food.append(int(calories))

    def get_foods(self):
        return self.food

    def get_total_calories(self):
        return sum(self.food)


def generate_elves(path):
    elves = []
    with open(path, "r") as fp:
        current_elf = elf()
        for line in fp.readlines():
            if line == "\n":
                elves.append(current_elf)
                current_elf = elf()
            else:
                current_elf.add_food(line)
        elves.append(current_elf)
    return elves


if __name__ == "__main__":
    elves = generate_elves("./input.txt")
    # print elf carrying the most food
    elf_sums = [elf.get_total_calories() for elf in elves]
    elf_sums.sort()
    print(f"Part1 Answer: {elf_sums[-1]}")

    # print sum of top 3 elves
    top_three_elves = elf_sums[-3:]
    print(f"Part2 Answer: {sum(top_three_elves)}")

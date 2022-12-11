#!/usr/bin/env python3

import re

MONKEYS = []


class Monkey(object):
    def __init__(self):
        self.inspections = 0
        self.items = None
        self.op = None
        self.test = None
        self.worry_divisor = None
        self.common_divisor = None
        self.m1 = None  # monkey1, throw to in pass case
        self.m2 = None  # monkey2, throw to in fail case

    def reprogram(self, items, op, test, m1, m2, wd, cd):
        self.inspections = 0
        self.items = items
        self.op = op
        self.test = test
        self.m1 = m1
        self.m2 = m2
        self.worry_divisor = wd
        self.common_divisor = cd

    def catch(self, item):
        self.items.append(item)

    def take_turn(self):
        while self.items:
            item = self.items.pop(0)
            self._throw(item)

    def _inspect(self, item):
        self.inspections += 1
        value = self.op(item)
        if self.common_divisor:
            value = value % self.common_divisor
        value = value // self.worry_divisor
        return value

    def _throw(self, item):
        item = self._inspect(item)
        if item % self.test == 0:
            self.m1.catch(item)
        else:
            self.m2.catch(item)


def precreate_monkeys(data):
    for line in data:
        if re.match("Monkey \d+:", line):
            MONKEYS.append(Monkey())


def program_monkeys(data, worry_divisor, common_divisor):
    for i in range(len(MONKEYS)):
        program_monkey(data[i * 7 : (i + 1) * 7 - 1], worry_divisor, common_divisor)


def program_monkey(data, wd, cd):
    monkey_id = int(re.match("Monkey (\d+):", data[0]).group(1))

    starting_items_str = re.match("  Starting items: (.*)", data[1]).group(1)
    starting_items = starting_items_str.split(", ")
    starting_items = [int(item) for item in starting_items]

    operation_str = re.match("  Operation: new = (.*)\n", data[2]).group(1)
    exec(f"def m{monkey_id}_op(old): return {operation_str}")
    operation = locals()[f"m{monkey_id}_op"]

    test = int(re.match("  Test: divisible by (\d+)", data[3]).group(1))

    pass_str = re.match("    If true: throw to monkey (\d+)", data[4]).group(1)
    pass_monkey = MONKEYS[int(pass_str)]
    fail_str = re.match("    If false: throw to monkey (\d+)", data[5]).group(1)
    fail_monkey = MONKEYS[int(fail_str)]

    MONKEYS[monkey_id].reprogram(starting_items, operation, test, pass_monkey, fail_monkey, wd, cd)


def run_rounds(count):
    for _ in range(count):
        for monkey in MONKEYS:
            monkey.take_turn()


def get_common_divisor():
    divisors = [monkey.test for monkey in MONKEYS]
    common_div = 1
    for val in divisors:
        common_div = common_div * val
    return common_div


if __name__ == "__main__":
    with open("./input.txt", "r") as fp:
        data = fp.readlines()
    precreate_monkeys(data)

    program_monkeys(data, 3, None)
    run_rounds(20)
    inspections = [monkey.inspections for monkey in MONKEYS]
    inspections.sort(reverse=True)
    print(f"Part1 Answer: {inspections[0]*inspections[1]}")

    program_monkeys(data, 1, get_common_divisor())
    run_rounds(10000)
    inspections = [monkey.inspections for monkey in MONKEYS]
    inspections.sort(reverse=True)
    print(f"Part2 Answer: {inspections[0]*inspections[1]}")

#!/usr/bin/env python3
import re
import sys


if len(sys.argv) != 2:
    SOLVE_FOR = "part1"
else:
    SOLVE_FOR = sys.argv[1]

STACKS = {}


class Stack(object):
    def __init__(self):
        self.stack = ""

    def preload(self, item):
        self.stack = item + self.stack

    def remove(self, count):
        self.stack, ret = self.stack[:-count], self.stack[-count:]
        if SOLVE_FOR == "part1":
            return ret[::-1]
        return ret

    def add(self, items):
        self.stack = self.stack + items


def create_stacks(file_path):
    global STACKS
    with open(file_path, "r") as fp:
        for line in fp.readlines():
            if re.match("^ 1.*", line):
                count = line.strip().split(" ")[-1]
                for i in range(1, int(count) + 1):
                    STACKS[i] = Stack()
                return
    raise Exception("never found the stack count line")


def init_stacks(file_path):
    global STACKS
    with open(file_path, "r") as fp:
        for line in fp.readlines():
            if re.match("^ 1.*", line):
                return
            for i in range(0, len(STACKS)):
                index = (i * 4) + 1
                if line[index] != " ":
                    STACKS[i + 1].preload(line[index])


def process_moves(file_path):
    global STACKS
    with open(file_path, "r") as fp:
        for line in fp.readlines():
            if re.match("move \d+ from \d+ to \d+", line):
                count, src, dst = re.findall("\d+", line)
                substack = STACKS[int(src)].remove(int(count))
                STACKS[int(dst)].add(substack)


def print_stacks():
    global STACKS
    for _, stack in STACKS.items():
        print(stack.stack)


if __name__ == "__main__":
    input_file = "./input.txt"
    create_stacks(input_file)
    init_stacks(input_file)
    process_moves(input_file)
    answer_str = ""
    for stack in STACKS.values():
        answer_str = answer_str + stack.stack[-1]
    print(f"{SOLVE_FOR} Answer: {answer_str}")

#!/usr/bin/env python3
CUR_X = 1
CYCLES = [1]  # X value during each cycle

CYCLES_TO_CHECK = [20, 60, 100, 140, 180, 220]


def process_noop():
    CYCLES.append(CUR_X)


def process_addx(value):
    global CUR_X
    CYCLES.append(CUR_X)
    CYCLES.append(CUR_X)
    CUR_X += value


def compute_signal_str():
    sum = 0
    for cycle in CYCLES_TO_CHECK:
        sum += cycle * CYCLES[cycle]
    return sum


def print_display():
    line = ""
    for i in range(1, 241):
        position = (i % 40) - 1
        if position < 0:
            position = 39
        if abs(position - CYCLES[i]) > 1:
            line = line + "."
        else:
            line = line + "#"
        if i % 40 == 0:
            print(line)
            line = ""


if __name__ == "__main__":
    with open("input.txt", "r") as fp:
        for line in fp.readlines():
            details = line.strip().split(" ")
            if details[0] == "noop":
                process_noop()
            else:
                process_addx(int(details[1]))

    sig_str = compute_signal_str()
    print(f"Part1 Answer: {sig_str}")
    print(f"Part2 Answer: See characters below")
    print_display()

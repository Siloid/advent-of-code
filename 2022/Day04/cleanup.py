#!/usr/bin/env python3

elf_duplicates = 0
overlap_total = 0
with open("./input.txt", "r") as fp:
    for line in fp.readlines():
        elf1, elf2 = line.strip().split(",")
        elf1_start, elf1_end = elf1.split("-")
        elf2_start, elf2_end = elf2.split("-")

        elf1_set = set(range(int(elf1_start), int(elf1_end) + 1))
        elf2_set = set(range(int(elf2_start), int(elf2_end) + 1))

        if elf1_set.issubset(elf2_set) or elf2_set.issubset(elf1_set):
            elf_duplicates += 1

        this_overlap = elf1_set & elf2_set
        if len(this_overlap):
            overlap_total += 1

    print(f"Part1 Answer: Duplicates: {elf_duplicates}")
    print(f"Part2 Answer: Overlaps: {overlap_total}")

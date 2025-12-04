#!/usr/bin/env python3
import re


def parse_input(filepath):
    product_ids = []
    with open(filepath, "r") as inputfile:
        input = inputfile.read()
        ranges = input.split(",")
        for id_range in ranges:
            range_start, range_end = id_range.split("-")
            product_ids += range(int(range_start), int(range_end) + 1)
    return product_ids


def get_twice_repeated_ids(ids):
    matching_ids = []
    for id in ids:
        id_str = str(id)
        id_length = len(id_str)
        if id_length % 2 == 0:
            if id_str[: id_length // 2] == id_str[id_length // 2 :]:
                matching_ids.append(int(id_str))
    return matching_ids


def get_multiple_repeating_ids(ids):
    matching_ids = []
    for id in ids:
        id_str = str(id)
        id_length = len(id_str)
        for i in range(1, (id_length // 2) + 1):
            id_slice = id_str[:i]
            pattern = rf"^({id_slice})+$"
            match = re.match(pattern, id_str)
            if match != None:
                matching_ids.append(int(id_str))
                break
    return matching_ids


def main():
    product_ids = parse_input("input.txt")

    part1_invalid_ids = get_twice_repeated_ids(product_ids)
    print(f"Part1 - Twice Repated ID Sum: {sum(part1_invalid_ids)}")

    part2_invalid_ids = get_multiple_repeating_ids(product_ids)
    print(f"Part2 - N Repated ID Sum: {sum(part2_invalid_ids)}")


if __name__ == "__main__":
    main()

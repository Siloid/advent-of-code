#!/usr/bin/env python3


def parse_input(filepath):
    fresh_ranges = []
    with open(filepath, "r") as inputfile:
        input = inputfile.read()
        sections = input.split("\n\n")
        for line in sections[0].split("\n"):
            begin, end = line.split("-")
            fresh_ranges.append((int(begin), int(end)))
        ingredients = [int(line) for line in sections[1].split("\n")]
    return fresh_ranges, ingredients


def count_fresh_ingredients(fresh_ranges, ingredients):
    fresh_counter = 0
    for ingredient in ingredients:
        for begin, end in fresh_ranges:
            if ingredient >= begin and ingredient <= end:
                fresh_counter += 1
                break
    return fresh_counter


def collapse_ranges(fresh_ranges):
    sorted_ranges = sorted(fresh_ranges)
    collapsed_ranges = []
    begin, end = sorted_ranges[0]
    for i in range(1, len(sorted_ranges)):
        new_begin, new_end = sorted_ranges[i]
        if new_begin <= end + 1:
            if new_end > end:
                end = new_end
        else:
            collapsed_ranges.append((begin, end))
            begin = new_begin
            end = new_end
    if (begin, end) not in collapsed_ranges:
        collapsed_ranges.append((begin, end))
    return collapsed_ranges


def sum_ranges(fresh_ranges):
    total = 0
    for begin, end in fresh_ranges:
        total += (end - begin) + 1
    return total


def main():
    fresh_ranges, ingredients = parse_input("input.txt")
    fresh_ingredient_count = count_fresh_ingredients(fresh_ranges, ingredients)
    print(f"Part1 - Fresh Ingredient Count: {fresh_ingredient_count}")

    collapsed_ranged = collapse_ranges(fresh_ranges)
    fresh_ids_count = sum_ranges(collapsed_ranged)
    print(f"Part2 - Total Fresh Ids: {fresh_ids_count}")


if __name__ == "__main__":
    main()

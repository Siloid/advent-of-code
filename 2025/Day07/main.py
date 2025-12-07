#!/usr/bin/env python3


INPUT_FILE = "input.txt"


def parse_input(filepath):
    with open(filepath, "r") as inputfile:
        return inputfile.readlines()


def read_diagram(diagram):
    starting_position = None
    splitters = []
    for line in diagram:
        new_splitters = []
        for index, character in enumerate(line):
            if character == "S":
                starting_position = index
            elif character == "^":
                new_splitters.append(index)
        splitters.append(new_splitters)
    return starting_position, splitters


def compute_splits(starting_position, splitters):
    beams = set([starting_position])
    split_counter = 0
    for splitter_list in splitters:
        new_beams = set()
        for beam in beams:
            if beam in splitter_list:
                new_beams.add(beam - 1)
                new_beams.add(beam + 1)
                split_counter += 1
            else:
                new_beams.add(beam)
        beams = new_beams
    return split_counter


def compute_timelines(starting_position, splitters, time_line_cache={}):
    splitters_lookup = tuple(tuple(row) for row in splitters)
    if (starting_position, splitters_lookup) in time_line_cache.keys():
        return time_line_cache[(starting_position, splitters_lookup)]
    for index, splitter_list in enumerate(splitters):
        if starting_position in splitter_list:
            timelines = compute_timelines(
                starting_position - 1, splitters[index:], time_line_cache
            ) + compute_timelines(starting_position + 1, splitters[index:], time_line_cache)
            time_line_cache[(starting_position, splitters_lookup)] = timelines
            return timelines
    time_line_cache[(starting_position, splitters_lookup)] = 1
    return 1


def main():
    diagram = parse_input(INPUT_FILE)
    starting_position, splitters = read_diagram(diagram)
    splits = compute_splits(starting_position, splitters)
    print(f"Part1 - Total Splits: {splits}")

    timelines = compute_timelines(starting_position, splitters)
    print(f"Part2 - Total Timelines: {timelines}")


if __name__ == "__main__":
    main()

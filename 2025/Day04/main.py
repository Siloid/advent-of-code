#!/usr/bin/env python3


def parse_input(filepath):
    with open(filepath, "r") as inputfile:
        layout = inputfile.read()
        rows = layout.split("\n")
        return [list(row) for row in rows]


directions = [
    (0, -1),  # North
    (1, -1),  # Northeast
    (1, 0),  # East
    (1, 1),  # Southeast
    (0, 1),  # South
    (-1, 1),  # Southwest
    (-1, 0),  # West
    (-1, -1),  # Northwest
]


def is_accessible(layout, x, y):
    blockers = 0
    for x_offset, y_offset in directions:
        x_to_check = x + x_offset
        y_to_check = y + y_offset
        if (
            x_to_check < 0
            or x_to_check >= len(layout[x])
            or y_to_check < 0
            or y_to_check >= len(layout)
        ):
            continue
        if layout[y_to_check][x_to_check] == "@":
            blockers += 1
            if blockers >= 4:
                return False
    return True


def find_accessible_rolls(layout):
    accessible_rolls = []
    for x in range(len(layout[0])):
        for y in range(len(layout)):
            if layout[y][x] == "@":
                if is_accessible(layout, x, y):
                    accessible_rolls.append((x, y))
    return accessible_rolls


def remove_rolls(layout, to_remove):
    for x, y in to_remove:
        layout[y][x] = "."


def main():
    paper_map = parse_input("input.txt")
    accessible_rolls = find_accessible_rolls(paper_map)
    print(f"Part1 - Initial accessible rolls: {len(accessible_rolls)}")

    total_accessible_rolls = 0
    while accessible_rolls:
        total_accessible_rolls += len(accessible_rolls)
        remove_rolls(paper_map, accessible_rolls)
        accessible_rolls = find_accessible_rolls(paper_map)
    print(f"Part2 - Total accessible rolls: {total_accessible_rolls}")


if __name__ == "__main__":
    main()

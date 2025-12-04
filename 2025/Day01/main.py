#!/usr/bin/env python3


DIAL_DIGITS = 100
STARTING_POSITION = 50


def parse_input(filepath):
    with open(filepath, "r") as inputfile:
        return [rotation for rotation in inputfile]


def turn_dial(position, rotation):
    direction = rotation[0]
    digits = int(rotation[1:])

    if direction == "L":
        position -= digits
        while position < 0:
            position += DIAL_DIGITS
    else:  # right
        position += digits
        while position > DIAL_DIGITS - 1:
            position -= DIAL_DIGITS

    return position


def turn_dial_counting_zero_passes(position, rotation):
    direction = rotation[0]
    digits = int(rotation[1:])
    zero_passes = 0
    previous_position = position
    if direction == "L":
        current_position = previous_position - digits
        while current_position < 0:
            if previous_position != 0:
                zero_passes += 1
            previous_position = current_position
            current_position += DIAL_DIGITS
        if current_position == 0:
            zero_passes += 1
    else:  # right
        current_position = previous_position + digits
        while current_position > DIAL_DIGITS - 1:
            zero_passes += 1
            previous_position = current_position
            current_position -= DIAL_DIGITS
    return current_position, zero_passes


def main():
    rotations = parse_input("input.txt")

    zero_counter = 0
    current_position = STARTING_POSITION
    for rotation in rotations:
        current_position = turn_dial(current_position, rotation)
        if current_position == 0:
            zero_counter += 1
    print(f"Part1 - Count Zero Landings: {zero_counter}")

    zero_counter = 0
    current_position = STARTING_POSITION
    for rotation in rotations:
        current_position, zero_passes = turn_dial_counting_zero_passes(
            current_position, rotation
        )
        zero_counter += zero_passes
    print(f"Part2 - Count Zero Passes: {zero_counter}")


if __name__ == "__main__":
    main()

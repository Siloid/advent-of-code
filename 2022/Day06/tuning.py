#!/usr/bin/env python3


def find_unique_str(file_path, count):
    with open(file_path, "r") as fp:
        data = fp.read()
        for i in range(count, len(data) + 1):
            if len(set(data[i - count : i])) == count:
                return i


if __name__ == "__main__":
    start_of_packet = find_unique_str("./input.txt", 4)
    print(f"Part1 Answer: {start_of_packet}")
    start_of_message = find_unique_str("./input.txt", 14)
    print(f"Part2 Answer: {start_of_message}")

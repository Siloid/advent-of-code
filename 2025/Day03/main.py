#!/usr/bin/env python3


def parse_input(filepath):
    battery_pack = []
    with open(filepath, "r") as inputfile:
        for line in inputfile:
            line = line.strip()
            bank = [int(line[x]) for x in range(len(line))]
            battery_pack.append(bank)
    return battery_pack


def compute_max_joltage(bank, battery_count):
    battery_list = []
    for i in range(battery_count):
        if battery_count - (i + 1) == 0:
            next_battery = get_highest_battery_in_bank(bank)
        else:
            next_battery = get_highest_battery_in_bank(
                bank[: -(battery_count - (i + 1))]
            )
        battery_list.append(next_battery)
        next_battery_position = bank.index(next_battery)
        bank = bank[next_battery_position + 1 :]
    joltage_str = "".join([str(battery) for battery in battery_list])
    return int(joltage_str)


def get_highest_battery_in_bank(bank):
    sorted_bank = sorted(bank, reverse=True)
    return sorted_bank[0]


def main():
    battery_pack = parse_input("./input.txt")
    # Part1
    max_joltages = []
    for bank in battery_pack:
        max_joltages.append(compute_max_joltage(bank, 2))
    print(f"Part1 - Max Joltage (2 Batteries): {sum(max_joltages)}")

    # Part2
    max_joltages = []
    for bank in battery_pack:
        max_joltages.append(compute_max_joltage(bank, 12))
    print(f"Part2 - Max Joltage (12 Batteries): {sum(max_joltages)}")


if __name__ == "__main__":
    main()

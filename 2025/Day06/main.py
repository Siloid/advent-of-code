#!/usr/bin/env python3


INPUT_FILE = "input.txt"


def parse_input_human_math(filepath):
    data = None
    with open(filepath, "r") as inputfile:
        for line in inputfile:
            elements = line.split()
            if data is None:
                data = [[] for _ in range(len(elements))]
            for index, value in enumerate(elements):
                data[index].append(value)
    return data


def parse_input_cephalopod_math(filepath):
    raw_data = None
    data = []
    with open(filepath, "r") as inputfile:
        lines = inputfile.readlines()
        for line in lines:
            if raw_data is None:
                raw_data = [[] for _ in range(len(line))]
            for index, character in enumerate(line):
                raw_data[index].append(character)

    def append_numbers(data, numbers):
        values = []
        for number in numbers:
            # Clean out spaces in numbers and find operators
            if "+" in number:
                number = number.replace("+", "")
                operator = "+"
            if "*" in number:
                number = number.replace("*", "")
                operator = "*"
            if number.strip() != "":
                values.append(number.strip())
        values.append(operator)
        data.append(values)

    numbers = []
    for index, column in enumerate(reversed(raw_data)):
        if set(column) == set(["\n"]):
            pass
        if set(column) == set([" "]):
            append_numbers(data, numbers)
            numbers = []
        numbers.append("".join(column))
    append_numbers(data, numbers)
    return data


def compute_column_values(data):
    column_values = []
    for column in data:
        value = 0
        if column[-1] == "+":
            for num in column[:-1]:
                value += int(num)
        else:  # *
            value = 1
            for num in column[:-1]:
                value *= int(num)
        column_values.append(value)
    return column_values


def main():
    data = parse_input_human_math(INPUT_FILE)
    values = compute_column_values(data)
    print(f"Part1 - Sum of column values (human): {sum(values)}")

    cephalopod_data = parse_input_cephalopod_math(INPUT_FILE)
    values = compute_column_values(cephalopod_data)
    print(f"Part2 - Sum of column values (cephalopod): {sum(values)}")


if __name__ == "__main__":
    main()

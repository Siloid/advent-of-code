#!/usr/bin/env python3


def get_item_priority(item):
    sub_value = 96
    if item.isupper():
        sub_value = 38
    return ord(item) - sub_value


class RuckSack(object):
    def __init__(self, contents):
        if len(contents) % 2 != 0:
            raise Exception(
                f"RuckSack contents are not even, not sure how to handle:{contents}:{len(contents)}"
            )

        self.contents = contents
        self.compartment1 = contents[: int(len(contents) / 2)]
        self.compartment2 = contents[int(len(contents) / 2) :]

    def get_common_elements(self):
        return set(self.compartment1) & set(self.compartment2)


if __name__ == "__main__":
    rucksacks = []
    # Part1
    priority_sum = 0
    with open("./input.txt", "r") as fp:
        for line in fp.readlines():
            rucksack = RuckSack(line.strip())
            rucksacks.append(rucksack)
            priority_sum += get_item_priority(rucksack.get_common_elements().pop())
    print(f"Part1 Answer: Total priority: {priority_sum}")

    # Part2
    badge_sum = 0
    for x in range(1, int(len(rucksacks) / 3) + 1):
        current_rucksacks = rucksacks[(x - 1) * 3 : (x * 3)]
        badge = (
            set(current_rucksacks[0].contents)
            & set(current_rucksacks[1].contents)
            & set(current_rucksacks[2].contents)
        )
        badge_sum += get_item_priority(badge.pop())
    print(f"Part2 Answer: Total badge priority: {badge_sum}")

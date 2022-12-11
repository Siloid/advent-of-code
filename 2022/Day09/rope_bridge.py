#!/usr/bin/env python3


class RopeKnot(object):
    def __init__(self):
        self.path = [(0, 0)]
        self.nknot = None

    def assign_next_knot(self, knot):
        self.nknot = knot

    def location(self):
        return self.path[-1]

    def move(self, x, y):
        location = self.location()
        new_x = location[0] + x
        new_y = location[1] + y
        self.path.append((new_x, new_y))

        if self.nknot:
            nknot_loc = self.nknot.location()
            nknot_x = nknot_loc[0]
            nknot_y = nknot_loc[1]
            if abs(new_x - nknot_x) >= 2 or abs(new_y - nknot_y) >= 2:
                nknot_move_x = 0
                nknot_move_y = 0
                if new_x != nknot_x and new_y != nknot_y:
                    # move diagaonally
                    if new_x > nknot_x:
                        nknot_move_x = 1
                    else:
                        nknot_move_x = -1
                    if new_y > nknot_y:
                        nknot_move_y = 1
                    else:
                        nknot_move_y = -1
                elif new_x != nknot_x:
                    # move horizontally
                    if new_x > nknot_x:
                        nknot_move_x = 1
                    else:
                        nknot_move_x = -1
                else:
                    # move vertically
                    if new_y > nknot_y:
                        nknot_move_y = 1
                    else:
                        nknot_move_y = -1
                self.nknot.move(nknot_move_x, nknot_move_y)

    def get_unique_location_count(self):
        return len(set(self.path))


def move_knot(knot, direction, distance):
    for _ in range(int(distance)):
        if direction == "U":
            knot.move(0, 1)
        elif direction == "D":
            knot.move(0, -1)
        elif direction == "L":
            knot.move(-1, 0)
        elif direction == "R":
            knot.move(1, 0)


def build_knot_rope(length):
    rope = []
    for i in range(length):
        rope.append(RopeKnot())
        if i > 0:
            rope[i - 1].assign_next_knot(rope[i])
    return rope


if __name__ == "__main__":
    rope = build_knot_rope(10)
    with open("input.txt", "r") as fp:
        for line in fp.readlines():
            move_knot(rope[0], *line.strip().split(" "))

    knot2_count = rope[1].get_unique_location_count()
    print(f"Part1 Answer: {knot2_count}")
    knot10_count = rope[9].get_unique_location_count()
    print(f"Part2 Answer: {knot10_count}")

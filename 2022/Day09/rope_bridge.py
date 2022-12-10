#!/usr/bin/env python3

HEAD_PATH = [(0, 0)]
TAIL_PATH = [(0, 0)]


def move(direction, distance):
    current_head = HEAD_PATH[-1]
    current_tail = TAIL_PATH[-1]
    for _ in range(int(distance)):
        move_tail = False
        if direction == "U":
            next_head = (current_head[0] + 1, current_head[1])
            if next_head[0] - 2 >= current_tail[0]:
                move_tail = True
        elif direction == "D":
            next_head = (current_head[0] - 1, current_head[1])
            if next_head[0] + 2 <= current_tail[0]:
                move_tail = True
        elif direction == "L":
            next_head = (current_head[0], current_head[1] - 1)
            if next_head[1] + 2 <= current_tail[1]:
                move_tail = True
        elif direction == "R":
            next_head = (current_head[0], current_head[1] + 1)
            if next_head[1] - 2 >= current_tail[1]:
                move_tail = True
        current_head = next_head
        HEAD_PATH.append(current_head)
        if move_tail:
            next_tail_x = current_tail[0]
            next_tail_y = current_tail[1]
            if current_head[0] != current_tail[0] and current_head[1] != current_tail[1]:
                # move diagaonally
                if current_head[0] > current_tail[0]:
                    next_tail_x += 1
                else:
                    next_tail_x -= 1
                if current_head[1] > current_tail[1]:
                    next_tail_y += 1
                else:
                    next_tail_y -= 1
            elif current_head[0] != current_tail[0]:
                # move horizontally
                if current_head[0] > current_tail[0]:
                    next_tail_x += 1
                else:
                    next_tail_x -= 1
            else:
                # move vertically
                if current_head[1] > current_tail[1]:
                    next_tail_y += 1
                else:
                    next_tail_y -= 1
            current_tail = (next_tail_x, next_tail_y)
            TAIL_PATH.append(current_tail)


if __name__ == "__main__":
    with open("input.txt", "r") as fp:
        for line in fp.readlines():
            move(*line.strip().split(" "))

    unique_tail_locations = set(TAIL_PATH)
    print(f"Part1 Answer: {len(unique_tail_locations)}")

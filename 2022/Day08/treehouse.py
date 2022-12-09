#!/usr/bin/env python3


def get_tree_stats(forest, x, y):
    # returns is_visible(bool), scenic_score(int)
    height = forest[x][y]
    hidden_north = False
    hidden_south = False
    hidden_east = False
    hidden_west = False
    trees_north = 0
    trees_south = 0
    trees_east = 0
    trees_west = 0
    # check north
    for n in range(x - 1, -1, -1):
        trees_north += 1
        if forest[n][y] >= height:
            hidden_north = True
            break
    # check south
    for s in range(x + 1, len(forest)):
        trees_south += 1
        if forest[s][y] >= height:
            hidden_south = True
            break
    for e in range(y + 1, len(forest[y])):
        trees_east += 1
        if forest[x][e] >= height:
            hidden_east = True
            break
    for w in range(y - 1, -1, -1):
        trees_west += 1
        if forest[x][w] >= height:
            hidden_west = True
            break

    if hidden_north and hidden_south and hidden_east and hidden_west:
        return False, trees_north * trees_south * trees_east * trees_west
    return True, trees_north * trees_south * trees_east * trees_west


if __name__ == "__main__":
    with open("input.txt", "r") as fp:
        forest = fp.readlines()
        forest = [row.strip() for row in forest]

    outer_visible = 2 * len(forest) + (2 * (len(forest[0]) - 2))
    max_scenic_score = 0
    inner_visible = 0
    for x in range(1, len(forest[0]) - 1):
        for y in range(1, len(forest) - 1):
            is_visible, scenic_score = get_tree_stats(forest, x, y)
            if is_visible:
                inner_visible += 1
            if scenic_score > max_scenic_score:
                max_scenic_score = scenic_score

    print(f"Part1 Answer: {outer_visible + inner_visible}")
    print(f"Part2 Answer: {max_scenic_score}")

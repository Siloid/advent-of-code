#!/usr/bin/env python3


INPUT_FILE = "input.txt"


def parse_input(filepath):
    tiles = []
    with open(filepath, "r") as inputfile:
        for line in inputfile:
            x, y = line.split(",")
            tiles.append((int(x), int(y)))
    return tiles


def find_perimeter(tiles):
    green_tiles = []
    previous_tile = tiles[0]
    for tile in tiles[1:]:
        green_tiles += get_tiles_between(previous_tile, tile)
        previous_tile = tile
    # connect first and last
    green_tiles += get_tiles_between(tiles[0], tiles[-1])
    return tiles + green_tiles


def get_tiles_between(tileA, tileB):
    ref_tiles = [tileA, tileB]
    ref_tiles.sort()
    tiles = []
    for x in range(ref_tiles[0][0], ref_tiles[1][0]+1):
        for y in range(ref_tiles[0][1], ref_tiles[1][1]+1):
            if (x,y) not in ref_tiles:
                tiles.append((x,y))
    return tiles


def is_within_perimeter(boundary, perimeter):
    min_x, max_x, min_y, max_y = boundary
    for p_tile in perimeter:
        if p_tile[0] > min_x and p_tile[0] < max_x and p_tile[1] > min_y and p_tile[1] < max_y:
            return False
    return True


def find_areas(tiles):
    areas = {}
    for index, tileA in enumerate(tiles):
        for tileB in tiles[index+1:]:
            x_dist = abs(tileA[0] - tileB[0]) + 1 
            y_dist = abs(tileA[1] - tileB[1]) + 1
            area = x_dist * y_dist
            areas[(tileA, tileB)] = area
    return {tiles: area for tiles, area in sorted(areas.items(), key=lambda item: item[1], reverse=True)}


def find_largest_area_green_red_tiles(areas, perimeter):
    for tiles, area in areas.items():
        is_all_red_green = True
        xs = [tiles[0][0], tiles[1][0]]
        ys = [tiles[0][1], tiles[1][1]]
        xs.sort()
        ys.sort()
        if is_within_perimeter(xs + ys, perimeter):
            return area


def main():
    red_tiles = parse_input(INPUT_FILE)
    perimeter = find_perimeter(red_tiles)
    areas = find_areas(red_tiles)

    max_red_corner_area = next(iter(areas.values()))
    print(f"Part1 - Largest rectangle area: {max_red_corner_area}")

    rg_max_area = find_largest_area_green_red_tiles(areas, perimeter)
    print(f"Part2 - Largest red and green only area: {rg_max_area}")
    

if __name__ == "__main__":
    main()

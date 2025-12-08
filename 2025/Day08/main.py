#!/usr/bin/env python3
import math


INPUT_FILE = "input.txt"
PART1_CONNECTIONS = 1000


class JunctionBox:
    def __init__(self, id, x, y, z):
        self.id = id
        self.x = x
        self.y = y
        self.z = z

    def distance_to(self, box):
        x_dist = self.x - box.x
        y_dist = self.y - box.y
        z_dist = self.z - box.z
        return math.sqrt((x_dist**2) + (y_dist**2) + (z_dist**2))


def parse_input(filepath):
    boxes = []
    with open(filepath, "r") as inputfile:
        box_id = 0
        for line in inputfile:
            box_id += 1
            x, y, z = line.split(",")
            boxes.append(JunctionBox(box_id, int(x), int(y), int(z)))
    return boxes


def compute_all_distances(junction_boxes):
    distances = {}
    for index, boxA in enumerate(junction_boxes):
        for boxB in junction_boxes[index + 1 :]:
            ids = [boxA.id, boxB.id]
            ids.sort()
            distances[tuple(ids)] = boxA.distance_to(boxB)
    return {ids: distance for ids, distance in sorted(distances.items(), key=lambda item: item[1])}


def connect_circuits(boxes, connections=None):
    distances = compute_all_distances(boxes)
    distance_keys = list(distances.keys())
    circuits = [[box.id] for box in boxes]
    if connections is None:
        connections = len(distances)
    idA = 0
    idB = 0
    for i in range(connections):
        idA, idB = distance_keys[i]
        circuitA = None
        circuitB = None
        for circuit in circuits:
            if idA in circuit:
                circuitA = circuit
            if idB in circuit:
                circuitB = circuit
        if circuitA != circuitB:
            circuits.remove(circuitA)
            circuits.remove(circuitB)
            circuits.append(circuitA + circuitB)
        if len(circuits) == 1:
            return circuits, idA, idB
    return circuits, idA, idB


def compute_circuit_value(circuits):
    circuits_by_length = sorted(circuits, key=len, reverse=True)
    total = 1
    for circuit in circuits_by_length[0:3]:
        total *= len(circuit)
    return total


def find_box_by_id(boxes, id):
    for box in boxes:
        if box.id == id:
            return box


def main():
    junction_boxes = parse_input(INPUT_FILE)
    connections, _, _ = connect_circuits(junction_boxes, PART1_CONNECTIONS)
    circuit_value = compute_circuit_value(connections)
    print(f"Part1 - Three largest circuits value: {circuit_value}")

    _, idA, idB = connect_circuits(junction_boxes)
    boxA = find_box_by_id(junction_boxes, idA)
    boxB = find_box_by_id(junction_boxes, idB)
    print(f"Part2 - Final connection A.x * B.x: {boxA.x * boxB.x}")


if __name__ == "__main__":
    main()

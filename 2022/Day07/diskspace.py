#!/usr/bin/env python3

import re


class File(object):
    def __init__(self, name, size):
        self.name = name
        self.size = int(size)


class Directory(object):
    def __init__(self, path):
        self.path = path
        self.subdirs = []
        self.files = []

    def add_file(self, name, size):
        self.files.append(File(name, size))

    def add_dir(self, directory):
        self.subdirs.append(directory)

    def get_size(self):
        size = 0
        for file in self.files:
            size += file.size
        for subdir in self.subdirs:
            size += subdir.get_size()
        return size


def process_ls(current_dir, history):
    directory = DISK[current_dir]
    while history:
        line = history.pop(0).strip()
        if re.match("^\$", line):
            history.insert(0, line)
            break
        line = line.split(" ")
        if line[0] == "dir":
            if directory.path == "/":
                subdir_path = "/" + line[1]
            else:
                subdir_path = directory.path + "/" + line[1]
            if subdir_path in DISK.keys():
                subdir = DISK[subdir_path]
            else:
                subdir = Directory(subdir_path)
                DISK[subdir_path] = subdir
            directory.add_dir(subdir)
        else:
            directory.add_file(line[1], line[0])


def process_cd(current_dir, cmd):
    details = cmd.strip().split(" ")
    if details[2] == "/":
        current_dir = "/"
    elif details[2] == "..":
        dir_list = current_dir.split("/")[:-1]
        if len(dir_list) == 1:
            current_dir = "/"
        else:
            current_dir = ("/").join(dir_list)
    else:
        if current_dir == "/":
            current_dir = current_dir + details[2]
        else:
            current_dir = current_dir + "/" + details[2]
    return current_dir


def build_disk(file_path):
    current_dir = "/"
    with open(file_path, "r") as fp:
        history = fp.readlines()
        while history:
            line = history.pop(0)
            if re.match("^\$ cd", line):
                current_dir = process_cd(current_dir, line.strip())
            elif re.match("^\$ ls", line):
                process_ls(current_dir, history)
            else:
                raise Exception(f"we're not processing something correctly, current line: {line}")


DISK = {"/": Directory("/")}
TOTAL_DISK_SIZE = 70000000
SIZE_FOR_UPDATE = 30000000

if __name__ == "__main__":
    build_disk("./input.txt")
    total_dir_size_smaller_than_100000 = 0
    for directory in DISK.values():
        if directory.get_size() < 100000:
            total_dir_size_smaller_than_100000 += directory.get_size()
    print(f"Part1 Answer: {total_dir_size_smaller_than_100000}")

    total_used = DISK["/"].get_size()
    free_space = TOTAL_DISK_SIZE - total_used
    min_space_to_free = SIZE_FOR_UPDATE - free_space

    dir_to_delete = min(
        [
            directory.get_size()
            for directory in DISK.values()
            if directory.get_size() >= min_space_to_free
        ]
    )
    print(f"Part2 Answer: {dir_to_delete}")


import os
import time

days_completed = 14

def run(program: str, input: str):
    start = time.time()
    os.system(f'go run {program} < {input}')
    print(f'Runtime {round(time.time() - start, 3)}s')


def run_day(dir: str):
    print(f'======== {dir} ========\n')

    print(f'Part 1: ')
    run(f'{dir}/part1/main.go', f'{dir}/input.txt')
    print()

    print(f'Part 2:')
    run(f'{dir}/part2/main.go', f'{dir}/input.txt')
    print()


def main():
    for i in range(1, days_completed + 1):
        run_day(f'Day{i:02}')


if __name__ == '__main__':
    main()

import argparse


def parse():
    parser = argparse.ArgumentParser(
        description="Simple CLI Application Launcher")
    parser.add_argument('-l', '--app-list-file', type=str,
                        required=True, help="Path to the application list YAML file")
    return parser.parse_args()

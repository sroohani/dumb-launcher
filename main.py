#!/usr/bin/env python3

import constants as cte
import cmd_ln as cl
import app_list as al
import sys

def main():
    args = cl.parse()
    try:
        app_dict = al.load_app_list(args.app_list_file)
    except Exception as e:
        print(f"Error loading/validating YAML data: {e}")
        sys.exit(-1)

if __name__ == "__main__":
    main()

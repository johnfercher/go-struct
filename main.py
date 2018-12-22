from src.infrastructure.file_finder import FileFinder

import sys

def main():
    if(len(sys.argv) != 3):
        print("Invalid arguments, should be: project_folder extension")
        return

    print(str(sys.argv))

    fileFinder = FileFinder()
    fileFinder.find_all_files(str(sys.argv[1]), "*." + str(sys.argv[2]))

if __name__ == "__main__":
    # execute only if run as a script
    main()

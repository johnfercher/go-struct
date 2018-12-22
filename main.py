from src.infrastructure.file_loader import FileLoader
from src.csharp.module_interpreter import ModuleInterpreter

import sys

def main():
    file_loader = FileLoader()
    module_interpreter = ModuleInterpreter()

    if(len(sys.argv) != 3):
        print("Invalid arguments, should be: project_folder extension")
        return

    print(str(sys.argv))

    files = file_loader.load_all_files(str(sys.argv[1]), "*." + str(sys.argv[2]))
    modules = module_interpreter.get_modules(files)


if __name__ == "__main__":
    # execute only if run as a script
    main()

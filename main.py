from src.infrastructure.services.file_loader import FileLoader
from src.csharp.cs_module_interpreter import CsModuleInterpreter

import sys

def main():
    file_loader = FileLoader()
    cs_module_interpreter = CsModuleInterpreter()

    if(len(sys.argv) != 3):
        print("Invalid arguments, should be: project_folder extension")
        return

    print(str(sys.argv))

    files = file_loader.load_all_files(str(sys.argv[1]), "*." + str(sys.argv[2]))

    for file in files:
        module = cs_module_interpreter.get_module(file)


if __name__ == "__main__":
    # execute only if run as a script
    main()

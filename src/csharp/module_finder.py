import os
from glob import glob

from src.domain.interfaces.Imodule_finder import IModuleFinder

class ModuleFinder(IModuleFinder):
    def find_all_modules(self, path, extension):
        result = [y for x in os.walk(path) for y in glob(os.path.join(x[0], extension))]
        print(result)
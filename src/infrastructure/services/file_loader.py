import os
from glob import glob

from src.domain.interfaces.Ifile_loader import IFileLoader
from src.domain.models.file import File

class FileLoader(IFileLoader):

    def load_all_files(self, path, extension):
        print("Searching for " + extension + " files in " + path + " folder.")

        filenames = self.__get_filenames(path, extension)

        if (len(filenames) == 0):
            print("Could not find any " + extension + " file in " + path + " folder.")
            return
        else:
            print("Found " + str(len(filenames)) + " " + extension + " files in " + path + " folder.")

        print("Loading files content.")

        files = self.__build_files(filenames)

        print("Files loaded.")

        return files

    def __get_filenames(self, path, extension):
        return [y for x in os.walk(path) for y in glob(os.path.join(x[0], extension))]

    def __build_files(self, filenames):
        files = list()

        for filename in filenames:
            with open(filename, 'r') as opened_file:
                file = File(filename, opened_file.read())
                files.append(file)

        return files

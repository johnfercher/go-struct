from src.domain.interfaces.Isyntax_interpreter import ISyntaxInterpreter

import re

class CsSyntaxInterpreter(ISyntaxInterpreter):

    def get_classes(self, string_content):
        # Interpret class ClassName
        match = re.search('class ([^\s]+)', string_content)

        if (match):
            return match.group(0).replace('class ', '')

        return None

    def has_entry_points(self, string_content):
        # Interpret class ClassName : Controller
        match = re.search('class ([^\s]+) : Controller', string_content)

        if (match):
            return True

        return False

    def get_methods(self, string_content):
        methods = list()

        methods.append(self.__get_sync_methods(string_content))
        methods.append(self.__get_async_methods(string_content))

        return methods

    def get_properties(self, string_content):
        # regex para propriedades public readonly ITipo Tipo;
        # /(public readonly ([^\s]+) ([^\s]+))[\s\S]/g

        # regex para propriedades public ITipo Tipo {
        # /(public ([^\s]+) ([^\s]+))[\s\S]{/g
        return 0

    def __get_sync_methods(self, string_content):
        # Interpret public ReturnType MethodName(
        pattern = '(public ([^\s]+) ([^\s]+)\()[\s\S]'

        return self.__get_all_method_names(pattern, string_content)

    def __get_async_methods(self, string_content):
        # Interpret public async Task<ReturnType> MethodName(
        pattern = '(public ([^\s]+) ([^\s]+) ([^\s]+)\()[\s\S]'

        return self.__get_all_method_names(pattern, string_content)

    def __get_all_method_names(self, pattern, string_content):
        occurrences = list()

        match_occurrences = re.findall(pattern, string_content)

        if (match_occurrences):
            for match_occurrence in match_occurrences:
                occurrences.append(match_occurrence[2])

        return occurrences

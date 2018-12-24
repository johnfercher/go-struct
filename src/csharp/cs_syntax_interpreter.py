from src.domain.interfaces.Isyntax_interpreter import ISyntaxInterpreter
from src.domain.models.class_definition import ClassDefinition
from src.domain.models.method_definition import MethodDefinition

import re

class CsSyntaxInterpreter(ISyntaxInterpreter):

    # TODO: Refactor
    def get_class_definition(self, string_content):
        class_declaration = None

        class_declaration = self.__get_class_with_inheritance(string_content)

        if(class_declaration != None):
            class_without_class_keyword = self.__remove_class(class_declaration)
            class_name_and_interface = class_without_class_keyword.split(" : ")
            return ClassDefinition(class_name_and_interface[0], class_name_and_interface[1])

        if(class_declaration == None):
            class_declaration = self.__get_class_without_inheritance(string_content)

        if(class_declaration == None):
            return None

        class_without_class_keyword = self.__remove_class(class_declaration)

        return ClassDefinition(class_without_class_keyword, None)

    def has_entry_points(self, string_content):
        # Interpret class ClassName : Controller
        match = re.search('class ([^\s]+) : Controller', string_content)

        if (match):
            return True

        return False

    def get_methods_definitions(self, string_content):
        raw_declarations = list()

        raw_declarations.extend(self.__get_sync_methods(string_content))
        raw_declarations.extend(self.__get_async_methods(string_content))

        methods_without_access_level = list(map(self.__remove_access_level, raw_declarations))
        methods_without_async = list(map(self.__remove_async, methods_without_access_level))
        methods_without_return = list(map(self.__remove_return, methods_without_async))
        methods = list(map(self.__remove_parenthesis, methods_without_return))
        methods_definitions =  [MethodDefinition(method, None) for method in methods]

        return methods_definitions

    def get_module_calls(self, string_content):
        return None

    def get_properties(self, string_content):
        raw_declarations = list()

        raw_declarations.extend(self.__get_readonly_properties(string_content))
        raw_declarations.extend(self.__get_non_readonly_properties(string_content))

        if(len(raw_declarations) == 0):
            return None

        properties_without_access_level = list(map(self.__remove_access_level, raw_declarations))
        properties_without_readonly = list(map(self.__remove_readonly, properties_without_access_level))
        properties = list(map(self.__get_only_property_type, properties_without_readonly))

        return properties_without_access_level

    def __get_sync_methods(self, string_content):
        # Interpret public ReturnType MethodName(
        pattern = '(public ([^\s]+) ([^\s]+)\()[\s\S]'
        return self.__get_entire_declarations(pattern, string_content)

    def __get_async_methods(self, string_content):
        # Interpret public async Task<ReturnType> MethodName(
        pattern = '(public ([^\s]+) ([^\s]+) ([^\s]+)\()[\s\S]'
        return self.__get_entire_declarations(pattern, string_content)

    def __get_readonly_properties(self, string_content):
        # Interpret public readonly IType Type;
        pattern = '(public readonly ([^\s]+) ([^\s]+))[\s\S]'
        return self.__get_entire_declarations(pattern, string_content)

    def __get_non_readonly_properties(self, string_content):
        # Interpret public IType Type {
        pattern = '(public ([^\s]+) ([^\s]+))[\s\S]{'
        return self.__get_entire_declarations(pattern, string_content)

    def __get_class_without_inheritance(self, string_content):
        # Interpret class ClassName
        pattern = '(class ([^\s]+))'
        declarations = self.__get_entire_declarations(pattern, string_content)

        if(declarations == None or len(declarations) == 0):
            return None

        return declarations[0]

    def __get_class_with_inheritance(self, string_content):
        # Interpret class ClassName : IClassName
        pattern = '(class ([^\s]+) : ([^\s]+))'
        declarations = self.__get_entire_declarations(pattern, string_content)

        if(declarations == None or len(declarations) == 0):
            return None

        if("Controller" in declarations[0]):
            return None

        return declarations[0]

    def __get_entire_declarations(self, pattern, string_content):
        occurrences = list()

        match_occurrences = re.findall(pattern, string_content)

        if (match_occurrences):
            for match_occurrence in match_occurrences:
                occurrences.append(match_occurrence[0])

        return occurrences

    def __remove_access_level(self, value):
        value = value.replace('public ', '')
        value = value.replace('private ', '')
        value = value.replace('protected ', '')
        value = value.replace('internal ', '')

        return value

    def __remove_async(self, value):
        value = value.replace('async ', '')
        return value

    def __remove_readonly(self, value):
        value = value.replace('readonly ', '')
        return value

    def __remove_class(self, value):
        value = value.replace('class ', '')
        return value

    def __remove_parenthesis(self, value):
        value = value.replace('(', '')
        return value

    def __remove_return(self, value):
        value = value.split(' ')
        return value[1]

    def __get_only_property_type(self, value):
        value = value.split(' ')
        return value[0]

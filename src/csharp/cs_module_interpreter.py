from src.domain.interfaces.Imodule_interpreter import IModuleInterpreter
from src.csharp.cs_syntax_interpreter import CsSyntaxInterpreter
from src.domain.models.module import Module

class CsModuleInterpreter(IModuleInterpreter):
    def __init__(self):
        self.syntax_interpreter = CsSyntaxInterpreter()

    def get_module(self, file):
        print("Interpreting file " + file.name + " to a module.")

        class_name = self.syntax_interpreter.get_classes(file.content)

        if (class_name == None):
            print("File " + file.name + " is not a module.")
            return None

        methods = self.syntax_interpreter.get_methods(file.content)
        properties = self.syntax_interpreter.get_properties(file.content)
        has_entry_points = self.syntax_interpreter.has_entry_points(file.content)

        print("Builded module " + str(class_name) + ".")

        return Module(file.name, class_name, methods, properties, has_entry_points)

from src.domain.interfaces.Imodule_interpreter import IModuleInterpreter
from src.csharp.cs_syntax_interpreter import CsSyntaxInterpreter
from src.domain.models.module import Module

class CsModuleInterpreter(IModuleInterpreter):
    def __init__(self):
        self.syntax_interpreter = CsSyntaxInterpreter()

    def get_module(self, file):
        print("Interpreting file " + file.name + " to a module.")

        class_definition = self.syntax_interpreter.get_class_definition(file.content)

        if (class_definition == None):
            print("File " + file.name + " is not a module.")
            return None

        methods = self.syntax_interpreter.get_methods(file.content)
        properties = self.syntax_interpreter.get_properties(file.content)
        has_entry_points = self.syntax_interpreter.has_entry_points(file.content)

        print("Builded module " + str(class_definition.name) + ".")

        module = Module(file.name, class_definition, properties, methods, has_entry_points)
        print("Module File: " + module.file_name)
        print("Module ClassName: " + str(module.class_definition.name))
        print("Module Interface: " + str(module.class_definition.interface))
        print("Module Has Entry Points: " + str(module.has_entry_points))
        print("Module Properties")
        print(module.properties)
        print("Module Methods")
        print(module.methods)

        return module

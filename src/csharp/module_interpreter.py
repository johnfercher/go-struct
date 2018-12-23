from src.domain.interfaces.Imodule_interpreter import IModuleInterpreter

class ModuleInterpreter(IModuleInterpreter):

    def get_modules(self, files):
        print(files[0].content)
        return 0

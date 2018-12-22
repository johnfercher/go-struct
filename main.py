from src.csharp.module_finder import ModuleFinder

def main():
    moduleFinder = ModuleFinder()
    moduleFinder.find_all_modules("mock/csharp", "*.cs")

if __name__ == "__main__":
    # execute only if run as a script
    main()
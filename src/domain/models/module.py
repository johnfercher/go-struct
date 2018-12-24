class Module(object):

    def __init__(self, file_name, class_definition, properties, methods_definitions, has_entry_points):
        self.file_name = file_name
        self.class_definition = class_definition
        self.properties = properties
        self.methods_definitions = methods_definitions
        self.has_entry_points = has_entry_points

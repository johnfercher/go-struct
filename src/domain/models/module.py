class Module(object):

    def __init__(self, file_name, class_name, properties, methods, has_entry_points):
        self.file_name = file_name
        self.class_name = class_name
        self.properties = properties
        self.methods = methods
        self.has_entry_points = has_entry_points

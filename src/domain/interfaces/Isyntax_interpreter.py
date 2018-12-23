class ISyntaxInterpreter(object):

    def get_classes(self, string_content):
        raise NotImplementedError()

    def has_entry_points(self, string_content):
        raise NotImplementedError()

    def get_methods(self, string_content):
        raise NotImplementedError()

    def get_properties(self, string_content):
        raise NotImplementedError()

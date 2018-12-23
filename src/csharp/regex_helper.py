class RegexHelper(object):

    @classmethod
    def get_classes(self, string_content):
        # regex para classes:
        # /class ([^\s]+)/g
        return 0

    @classmethod
    def get_methods(self, string_content):
        # regex para métodos public retorno nome(:
        # /(public ([^\s]+) ([^\s]+)\()[\s\S]/g

        # regex para métodos public async Task<retorno> nome(:
        # /(public ([^\s]+) ([^\s]+) ([^\s]+)\()[\s\S]/g
        return 0

    @classmethod
    def get_properties(self, string_content):
        # regex para propriedades public readonly ITipo Tipo;
        # /(public readonly ([^\s]+) ([^\s]+))[\s\S]/g

        # regex para propriedades public ITipo Tipo {
        # /(public ([^\s]+) ([^\s]+))[\s\S]{/g
        return 0

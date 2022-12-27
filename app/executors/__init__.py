""" this will be the system result class """


class SystemResult(dict):
    """This will contain the results of the system command"""

    def __init__(self, ret: bool, output: str, err: str, act_return_code: int) -> None:
        dict.__init__(
            self,
            return_code=ret,
            std_output=output.strip(),
            std_err=err.strip(),
            actual_return_code=act_return_code,
        )


class SingletonMeta(type):
    """
    The Singleton class can be implemented in different ways in Python. Some
    possible methods include: base class, decorator, metaclass. We will use the
    metaclass because it is best suited for this purpose.
    """

    _instances = {}

    def __call__(cls, *args, **kwargs):
        """
        Possible changes to the value of the `__init__` argument do not affect
        the returned instance.
        """
        if cls not in cls._instances:
            instance = super().__call__(*args, **kwargs)
            cls._instances[cls] = instance
        return cls._instances[cls]

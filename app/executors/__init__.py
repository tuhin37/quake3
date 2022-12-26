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

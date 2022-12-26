""" this will run the system commands """
import subprocess
from typing import List

from app.executors import SystemResult


class SystemCaller(object):
    """this class will handle the actual system call"""

    def __init__(self) -> None:
        pass

    def execute(self, command: List, check_ret: bool = True) -> SystemResult:
        """this will execute the command and return whether that command was successful or not"""
        ret = True
        output = ""
        output_err = ""
        actual_code = 0
        try:
            val = subprocess.run(command, capture_output=True, check=check_ret)
            ret = val.returncode == 0
            output = val.stdout
            output_err = val.stderr
            actual_code = val.returncode
        except subprocess.CalledProcessError as err:
            ret = False
            output_err = err.stderr
            output = err.stdout
            actual_code = err.returncode
        res = SystemResult(
            ret, output.decode("ascii"), output_err.decode("ascii"), actual_code
        )
        return res

""" this will run the system commands """
import subprocess
import traceback
from typing import List, Any

from app.executors import SystemResult, SingletonMeta


class SystemCaller(metaclass=SingletonMeta):
    """this class will handle the actual system call"""

    __running_instance = None

    def __init__(self) -> None:
        self.val = None

    def execute(
        self,
        command: Any,
    ) -> SystemResult:
        """this will execute the command and return whether that command was successful or not"""
        if self.val:
            try:
                self.val.terminate()
            except Exception as exp:
                print(
                    "Got {} error trying to terminate:: \n{}".format(
                        str(exp), traceback.format_exception()
                    )
                )
        ret = True
        output = ""
        output_err = ""
        actual_code = 0
        if isinstance(command, List):
            comm = " ".join(command)
        else:
            comm = str(command)
        try:
            print('Running command:: {}'.format(comm))
            self.val = subprocess.check_output(
                comm,
                stderr=subprocess.STDOUT,
                shell=True,
            )
            ret = True
            output = self.val
            ran_cmd = comm
            actual_code = 0
            print('Got output:: {}'.format(output))
        except subprocess.CalledProcessError as err:
            ret = False
            ran_cmd = comm
            output = str(err.output)
            actual_code = int(err.returncode)
        res = SystemResult(ret, output, ran_cmd, actual_code)
        return res

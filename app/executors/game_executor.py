from app.executors.system_exec import SystemCaller, SingletonMeta
import traceback
import os
import signal
import subprocess

class GameExecutor(metaclass=SingletonMeta):
    def __init__(self) -> None:
        self.exec_path = "/usr/local/games/quake3/ioq3ded.x86_64"
        self.process = None

    def run_game(self, settings_file):
        command = self.exec_path + " +exec " + settings_file
        try:
            if self.process is not None:
                print('Terminating existing process')
                os.killpg(os.getpgid(self.process.pid), signal.SIGTERM)
            print('Starting new process')
            self.process = subprocess.Popen(command, stdout=subprocess.PIPE, 
                       shell=True, preexec_fn=os.setsid, )
        except Exception as exp:
            print('Got exception :: {}. Details:: \n{}'.format(exp, traceback.format_exception()))
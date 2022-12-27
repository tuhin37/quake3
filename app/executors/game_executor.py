from app.executors.system_exec import SystemCaller


class GameExecutor:
    def __init__(self, settings_file) -> None:
        self.settings_file = settings_file
        pass

    def run_game(self):
        print("This will run the game")
        #SystemCaller().execute("ps -ef | grep python > ./target/result.txt")
        exec_path = '/usr/local/games/quake3/ioq3ded.x86_64'
        command = exec_path + ' +exec ' + self.settings_file
        SystemCaller().execute(command)

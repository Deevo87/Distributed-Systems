import sys
import Ice
from handlers.CoffeeMakerHandler import CoffeeMakerHandler
from handlers.QuitHandler import QuitHandler
from handlers.SmartBlindHandler import SmartBlindHandler
from handlers.AlarmTriggeredBlindHandler import AlarmTriggeredBlindHandler
from handlers.DailyScheduleBlindHandler import DailyScheduleBlindHandler
from handlers.ThermostatHandler import ThermostatHandler


class Command:
    def __init__(self, command_raw):
        parts = command_raw.split(':')
        if len(parts) == 1:
            self.command = parts[0]
            self.deviceName = None
            self.arguments = None
        elif len(parts) == 2:
            temp = parts[1].split('--args=')
            if len(temp) == 1:
                self.arguments = ""
            else:
                self.arguments = temp[1].strip()
            self.deviceName = parts[0].strip()
            self.command = temp[0].strip()

    @staticmethod
    def read_from_input(prompt="> "):
        return Command(input(prompt))


class Client:
    def __init__(self, args):
        self.quit = False
        self.args = args
        self.communicator = None
        self.commands_handlers = None

    def add_config_to_args(self):
        return self.args + ["--Ice.Config=config.client"]

    def main_loop(self):
        while not self.quit:
            handled = False
            command = Command.read_from_input()
            try:
                for commands_handler in self.commands_handlers:
                    handled = commands_handler.handle_command(command)
                    if handled:
                        break
                if not handled:
                    print("Couldn't find a handler to handle the command")
            except Ice.ConnectionRefusedException:
                print("Couldn't connect to the server")
            except Exception as e:
                print("Some strange problems occured")
                print(e)

    def initialize(self):
        try:
            args_with_config = self.add_config_to_args()
            self.communicator = Ice.initialize(args_with_config)
            self.commands_handlers = [
                QuitHandler(self),
                SmartBlindHandler(self.communicator),
                AlarmTriggeredBlindHandler(self.communicator),
                DailyScheduleBlindHandler(self.communicator),
                ThermostatHandler(self.communicator),
                CoffeeMakerHandler(self.communicator)
            ]
        except Exception as e:
            print(e)

    def run(self):
        try:
            if self.communicator:
                self.main_loop()
            else:
                print("Communication not initialized.")
        finally:
            if self.communicator:
                self.communicator.destroy()


if __name__ == "__main__":
    client = Client(sys.argv)
    client.initialize()
    client.run()

from handlers.CommandHandler import CommandHandler


class DeviceHandler(CommandHandler):
    def __init__(self, communicator):
        self.communicator = communicator

    def handle_command(self, command):
        if command == "info":
            print("Handling device information...")
            return True
        elif command == "failure":
            print("Handling device failure...")
            return True
        else:
            return False


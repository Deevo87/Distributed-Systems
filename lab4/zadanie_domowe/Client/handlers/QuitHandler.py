from handlers.CommandHandler import CommandHandler


class QuitHandler(CommandHandler):
    def __init__(self, client):
        self.client = client

    def handle_command(self, command):
        if command.command == "quit":
            print("Quitting...")
            self.client.quit = True
            return True
        return False

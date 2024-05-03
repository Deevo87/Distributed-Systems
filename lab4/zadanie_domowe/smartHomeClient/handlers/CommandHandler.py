from abc import ABC, abstractmethod


class CommandHandler(ABC):
    @abstractmethod
    def handle_command(self, command):
        pass

from generate_ice import _M_SmartHome
from handlers.DeviceHandler import DeviceHandler


class SmartBlindHandler(DeviceHandler):
    def __init__(self, communicator):
        super().__init__(communicator)
        self.proxy = None

    def handle_command(self, command):
        if command.deviceName == "smart blinds":
            proxy_str = "SmartBlinds/SmartBlinds:tcp -h 127.0.0.11 -p 10011 : udp -h 127.0.0.11 -p 10011"
            proxy = self.communicator.stringToProxy(proxy_str)
            self.proxy = _M_SmartHome.SmartBlindsPrx.checkedCast(proxy)
            if command.command == "open blinds":
                print("Opening blinds...")
                self.proxy.openBlinds()
                return True
            elif command.command == "close blinds":
                print("Closing blinds...")
                self.proxy.closeBlinds()
                return True
            elif command.command == "set angle":
                print("Setting closing angle...")
                print(command.arguments)
                self.proxy.setAngle(float(command.arguments))
                return True
            elif command.command == "set window coverage":
                print("Setting new window coverage...")
                self.proxy.setCustomWindowCoverage(int(command.arguments))
                return True
        return False

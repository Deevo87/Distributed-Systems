from handlers.SmartBlindHandler import SmartBlindHandler
from generate_ice import _M_SmartHome


class AlarmTriggeredBlindHandler(SmartBlindHandler):
    def __init__(self, communicator):
        super().__init__(communicator)

    def handle_command(self, command):
        if command.deviceName == "alarm triggered blinds":
            proxy_str = ("AlarmTriggeredBlinds/AlarmTriggeredBlinds:tcp -h 127.0.0.11 -p 10011 : udp -h 127.0.0.11 -p 10011")
            proxy = self.communicator.stringToProxy(proxy_str)
            self.proxy = _M_SmartHome.AlarmTriggeredBlindsPrx.checkedCast(proxy)
            if command.command == "activate alarm":
                print("Activating alarm...")
                self.proxy.activateAlarmMode()
                return True
            elif command.command == "deactivate alarm":
                print("Deactivating alarm...")
                self.proxy.deactivateAlarmMode()
                return True
        return False

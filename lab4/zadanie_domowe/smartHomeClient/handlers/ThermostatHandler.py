from handlers.DeviceHandler import DeviceHandler
from generate_ice import _M_SmartHome


class ThermostatHandler(DeviceHandler):
    def __init__(self, communicator):
        super().__init__(communicator)
        self.proxy = None

    def handle_command(self, command):
        if command.deviceName == "thermostat":
            proxy_str = "Thermostat/Thermostat:tcp -h 127.0.0.11 -p 10011 : udp -h 127.0.0.11 -p 10011"
            proxy = self.communicator.stringToProxy(proxy_str)
            self.proxy = _M_SmartHome.ThermostatPrx.checkedCast(proxy)
            if command.command == "set temperature":
                print("setting thermostat temperature...")
                self.proxy.setTemperature(float(command.arguments))
                return True
            elif command.command == "check temperature":
                temperature = self.proxy.getTemperature()
                print("Temperature is " + str(temperature) + " Celsius.")
                return True
        return False

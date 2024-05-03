from handlers.DeviceHandler import DeviceHandler
from generate_ice import _M_SmartHome


class CoffeeMakerHandler(DeviceHandler):
    def __init__(self, communicator):
        super().__init__(communicator)
        self.proxy = None

    def handle_command(self, command):
        if command.deviceName == "coffee maker":
            proxy_str = ("CoffeeMaker/CoffeeMaker:tcp -h 127.0.0.11 -p 10011 : udp -h 127.0.0.11 -p 10011")
            proxy = self.communicator.stringToProxy(proxy_str)
            self.proxy = _M_SmartHome.CoffeeMakerPrx.checkedCast(proxy)
            if command.command == "brew coffee":
                print("Brewing coffee...")
                self.proxy.brewCoffee()
                return True
            elif command.command == "set schedule":
                print("Setting schedule...")
                self.proxy.setSchedule(self.command_args_handler(command.arguments))
                return True
            elif command.command == "add coffee time":
                print("Changing day schedule...")
                coffee_data = command.arguments.split(', ')
                day = _M_SmartHome.Day._enumerators[int(coffee_data[0]) - 1]
                newCoffeeTime = self.create_coffee_time(coffee_data[1:])
                self.proxy.addCoffeeTime(newCoffeeTime, day)
                return True
            elif command.command == "clear schedule":
                print("Clearing schedule...")
                self.proxy.clearSchedule()
                return True
            elif command.command == "clear day schedule":
                day = _M_SmartHome.Day._enumerators[int(command.arguments[0]) - 1]
                print("Clearing schedule on " + str(day))
                self.proxy.clearDaySchedule(day)
                return True
        return False

    def command_args_handler(self, data):
        days = data.split(' | ')
        schedule_list = []
        for day in days:
            day_data = day.split(' || ')
            schedule_list.append(self.create_day_schedule(day_data))
        return schedule_list

    def create_day_schedule(self, day_data):
        day_schedule = _M_SmartHome.CoffeeDay()
        day_schedule.dayOfWeek = _M_SmartHome.Day._enumerators[int(day_data[0]) - 1]
        coffee_times = []
        for i in range(1, len(day_data)):
            coffee_data = day_data[i].split(', ')
            coffee_times.append(self.create_coffee_time(coffee_data))
        day_schedule.coffeeTimes = coffee_times
        return day_schedule

    def create_coffee_time(self, coffee_data):
        coffee_time = _M_SmartHome.CoffeeTime()
        coffee_time.startHour = int(coffee_data[0])
        coffee_time.startMinute = int(coffee_data[1])
        print(coffee_time)
        return coffee_time

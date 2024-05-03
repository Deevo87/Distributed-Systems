from handlers.SmartBlindHandler import SmartBlindHandler
from generate_ice import _M_SmartHome


class DailyScheduleBlindHandler(SmartBlindHandler):
    def __init__(self, communicator):
        super().__init__(communicator)
        self.proxy = None

    def handle_command(self, command):
        if command.deviceName == "daily schedule blinds":
            proxy_str = ("DailyScheduleBlinds/DailyScheduleBlinds:tcp -h 127.0.0.11 -p 10011 : udp -h 127.0.0.11 -p 10011")
            proxy = self.communicator.stringToProxy(proxy_str)
            self.proxy = _M_SmartHome.DailyScheduleBlindsPrx.checkedCast(proxy)
            if command.command == "change daily schedule":
                print("Changing daily blinds schedule...")
                self.proxy.changeDaySchedule(self.create_day_schedule(command.arguments.split(', ')))
                return True
            elif command.command == "set blinds schedule":
                print("Setting blinds schedule...")
                schedule = self.command_args_handler(command.arguments)
                self.proxy.setSchedule(schedule)
                return True
        return False

    def command_args_handler(self, data):
        days = data.split(' | ')
        schedule_list = []
        for day in days:
            day_data = day.split(', ')
            schedule_list.append(self.create_day_schedule(day_data))
        return schedule_list

    def create_day_schedule(self, day_data):
        day_schedule = _M_SmartHome.DaySchedule()
        day_schedule.dayOfWeek = _M_SmartHome.Day._enumerators[int(day_data[0]) - 1]
        day_schedule.openHour = int(day_data[1])
        day_schedule.openMinutes = int(day_data[2])
        day_schedule.closeHour = int(day_data[3])
        day_schedule.closeMinutes = int(day_data[4])
        return day_schedule

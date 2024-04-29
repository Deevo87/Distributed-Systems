module SmartHome
{
    enum Info
    {
        Failure,
        Working
    };

    exception AlreadyWorkingError {};
    exception DeviceFailureError {};
    interface Device
    {
        idempotent Info setInfo(Info info) throws AlreadyWorkingError;
        idempotent Info getInfo();
        idempotent void isFailure() throws DeviceFailureError;
    };

    exception TemperatureOutOfRangeError {};
    interface Thermostat extends Device
    {
        void setTemperature(double temperature) throws TemperatureOutOfRangeError;
        double getTemperature();
    };

    struct DaySchedule
    {
        string dayOfWeek;
        int openHour;
        int openMinutes;
        int closeHour;
        int closeMinutes;
    };
    sequence<DaySchedule> Schedule;

    struct CoffeeTime
    {
        int startHour;
        int startMinute;
    };
    sequence<CoffeeTime> CoffeeTimes;

    exception TooManyShedulesError {};
    exception HoursOutOfRangeError {};
    exception MinutesOutOffRangeError {};
    interface CoffeeMaker extends Device
    {
        void brewCoffee();
        void setSchedule(CoffeeTimes coffeeTimes) throws TooManyShedulesError, HoursOutOfRangeError, MinutesOutOffRangeError;
        void changeDaySchedule(CoffeeTime coffeeTime) throws HoursOutOfRangeError, MinutesOutOffRangeError;
    };

    exception AngleOutOfRangeError {};
    exception CoverageOutOfRangeError {};
    interface SmartBlind extends Device
    {
        void openBlind();
        void closeBlind();
        void setAngle(double angle) throws AngleOutOfRangeError;
        void setFullWindowCoverage(bool fullCoverage) throws CoverageOutOfRangeError;
        void setCustomWindowCoverage(double coverage);
    };

    interface DailyScheduleBlind extends SmartBlind
    {
        void changeDaySchedule(DaySchedule daySchedule) throws HoursOutOfRangeError, MinutesOutOffRangeError;
        void setSchedule(Schedule schedule) throws TooManyShedulesError, HoursOutOfRangeError, MinutesOutOffRangeError;
    };

    exception AlarmAlreadyOnError {};
    exception AlarmAlreadyOffError {};
    interface AlarmTriggeredBlind extends SmartBlind
    {
        void activateAlarmMode() throws AlarmAlreadyOnError;
        void deactivateAlarmMode() throws AlarmAlreadyOffError;
    };
};

module SmartHome
{
    enum Info
    {
        Failure,
        Working
    };

    enum Day {
        Monday,
        Tuesday,
        Wednesday,
        Thursday,
        Friday,
        Saturday,
        Sunday
    };

    exception AlreadyWorkingError {};
    interface Device
    {
        idempotent void setInfo(Info info) throws AlreadyWorkingError;
        idempotent bool isWorking();
        idempotent bool isFailure();
    };

    exception TemperatureOutOfRangeError {};
    interface Thermostat extends Device
    {
        void setTemperature(double temperature) throws TemperatureOutOfRangeError;
        double getTemperature();
    };

    struct DaySchedule
    {
        Day dayOfWeek;
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

    struct CoffeeDay {
        Day dayOfWeek;
        CoffeeTimes coffeeTimes;
    };
    sequence<CoffeeDay> CoffeeWeek;

    exception TooManyShedulesError {};
    exception HoursOutOfRangeError {};
    exception MinutesOutOffRangeError {};
    interface CoffeeMaker extends Device
    {
        void brewCoffee();
        void setSchedule(CoffeeWeek coffeeWeek) throws HoursOutOfRangeError, MinutesOutOffRangeError, TooManyShedulesError;
        void addCoffeeTime(CoffeeTime coffeeTime, Day day) throws HoursOutOfRangeError, MinutesOutOffRangeError;
        void clearDaySchedule(Day day);
        void clearSchedule();
    };

    exception AngleOutOfRangeError {};
    exception CoverageOutOfRangeError {};
    exception AlreadyOpenedError {};
    exception AlreadyClosedError {};
    interface SmartBlinds extends Device
    {
        void openBlinds() throws AlreadyOpenedError;
        void closeBlinds() throws AlreadyClosedError;
        void setAngle(double angle) throws AngleOutOfRangeError;
        void setCustomWindowCoverage(int coverage) throws CoverageOutOfRangeError;
    };

    interface DailyScheduleBlinds extends SmartBlinds
    {
        void changeDaySchedule(DaySchedule daySchedule) throws HoursOutOfRangeError, MinutesOutOffRangeError;
        void setSchedule(Schedule schedule) throws TooManyShedulesError, HoursOutOfRangeError, MinutesOutOffRangeError;
    };

    exception AlarmAlreadyOnError {};
    exception AlarmAlreadyOffError {};
    interface AlarmTriggeredBlinds extends SmartBlinds
    {
        void activateAlarmMode() throws AlarmAlreadyOnError;
        void deactivateAlarmMode() throws AlarmAlreadyOffError;
    };
};
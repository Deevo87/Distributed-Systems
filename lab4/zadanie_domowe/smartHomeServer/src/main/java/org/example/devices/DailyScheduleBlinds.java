package org.example.devices;

import SmartHome.*;
import com.zeroc.Ice.Current;

import java.util.HashMap;
import java.util.Map;

public class DailyScheduleBlinds extends SmartBlinds implements SmartHome.DailyScheduleBlinds {
    private Map<Day, DaySchedule> schedule = new HashMap<>();

    public DailyScheduleBlinds() {}

    @Override
    public void changeDaySchedule(DaySchedule daySchedule, Current current) throws HoursOutOfRangeError, MinutesOutOffRangeError {
        if (daySchedule.openHour > 24 || daySchedule.closeHour > 24 || daySchedule.openHour < 1 || daySchedule.closeHour < 1) {
            throw new HoursOutOfRangeError();
        }
        if (daySchedule.openMinutes > 59 || daySchedule.closeMinutes > 59 || daySchedule.openMinutes < 0 || daySchedule.closeMinutes < 0) {
            throw new MinutesOutOffRangeError();
        }
        this.schedule.put(daySchedule.dayOfWeek, daySchedule);
        System.out.println("Changed schedule on " + daySchedule.dayOfWeek);
    }

    @Override
    public void setSchedule(DaySchedule[] schedule, Current current) throws HoursOutOfRangeError, MinutesOutOffRangeError, TooManyShedulesError {
        if (schedule.length > 7) {
            throw new TooManyShedulesError();
        }
        for (DaySchedule daySchedule : schedule) {
            if (daySchedule.openHour > 24 || daySchedule.closeHour > 24 || daySchedule.openHour < 1 || daySchedule.closeHour < 1) {
                throw new HoursOutOfRangeError();
            }
            if (daySchedule.openMinutes > 59 || daySchedule.closeMinutes > 59 || daySchedule.openMinutes < 0 || daySchedule.closeMinutes < 0) {
                throw new MinutesOutOffRangeError();
            }
        }
        for (DaySchedule daySchedule : schedule) {
            this.schedule.put(daySchedule.dayOfWeek, daySchedule);
        }
        System.out.println("Day Schedule Set");
        daySchedulePrint();
    }

    private void daySchedulePrint() {
        for (DaySchedule daySchedule : this.schedule.values()) {
            System.out.println(daySchedule.dayOfWeek);
            System.out.println("opening at " + daySchedule.openHour + ":" + daySchedule.openMinutes);
            System.out.println("closing at " + daySchedule.closeHour + ":" + daySchedule.closeMinutes);
        }
    }
}

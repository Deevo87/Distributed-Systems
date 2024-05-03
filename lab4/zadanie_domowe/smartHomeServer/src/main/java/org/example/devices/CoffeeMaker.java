package org.example.devices;

import SmartHome.*;
import com.zeroc.Ice.Current;

import java.util.HashMap;
import java.util.Map;


public class CoffeeMaker extends Device implements SmartHome.CoffeeMaker {
    private boolean brewingCoffee = false;
    private Map<Day, CoffeeDay> coffeeWeek = new HashMap<>();

    public CoffeeMaker() {}

    @Override
    public void brewCoffee(Current current) {
        this.brewingCoffee = true;
    }

    @Override
    public void setSchedule(CoffeeDay[] coffeeWeek, Current current) throws HoursOutOfRangeError, MinutesOutOffRangeError, TooManyShedulesError {
        if (coffeeWeek.length > 7) {
            throw new TooManyShedulesError();
        }
        for (CoffeeDay coffeeDay : coffeeWeek) {
            for (CoffeeTime coffeeTime : coffeeDay.coffeeTimes) {
                if (coffeeTime.startHour > 24 || coffeeTime.startHour < 1) {
                    throw new HoursOutOfRangeError();
                }
                if (coffeeTime.startMinute > 59 || coffeeTime.startMinute < 0) {
                    throw new MinutesOutOffRangeError();
                }
            }
        }
        for (CoffeeDay coffeeDay : coffeeWeek) {
            this.coffeeWeek.put(coffeeDay.dayOfWeek, coffeeDay);
        }
        System.out.println("Schedule set.");
        this.printCoffeeSchedule();
    }

    @Override
    public void addCoffeeTime(CoffeeTime coffeeTime, Day day, Current current) throws HoursOutOfRangeError, MinutesOutOffRangeError {
        if (coffeeTime.startHour > 24 || coffeeTime.startHour < 1) {
            throw new HoursOutOfRangeError();
        }
        if (coffeeTime.startMinute > 59 || coffeeTime.startMinute < 0) {
            throw new MinutesOutOffRangeError();
        }
        CoffeeTime[] newCoffeeTimes = new CoffeeTime[this.coffeeWeek.get(day).coffeeTimes.length + 1];
        System.arraycopy(this.coffeeWeek.get(day).coffeeTimes, 0, newCoffeeTimes, 0, this.coffeeWeek.get(day).coffeeTimes.length);
        newCoffeeTimes[newCoffeeTimes.length - 1] = coffeeTime;
        this.coffeeWeek.put(day, new CoffeeDay(day, newCoffeeTimes));
        System.out.println("Coffee time added at " + coffeeTime.startHour + ":" + coffeeTime.startMinute + " on " + day);
        this.printCoffeeSchedule();
    }

    @Override
    public void clearDaySchedule(Day day, Current current) {
        this.coffeeWeek.get(day).coffeeTimes = new CoffeeTime[0];
        System.out.println("Schedule on " + day + " cleared.");
        this.printCoffeeSchedule();
    }

    @Override
    public void clearSchedule(Current current) {
        this.coffeeWeek.clear();
        System.out.println("Schedule cleared.");
        this.printCoffeeSchedule();
    }

    private void printCoffeeSchedule() {
        for (Map.Entry<Day, CoffeeDay> entry : this.coffeeWeek.entrySet()) {
            System.out.println("Schedule on " + entry.getValue().dayOfWeek);
            for (CoffeeTime coffeeTime : entry.getValue().coffeeTimes) {
                System.out.println(coffeeTime.startHour + ":" + coffeeTime.startMinute);
            }
        }
    }
}

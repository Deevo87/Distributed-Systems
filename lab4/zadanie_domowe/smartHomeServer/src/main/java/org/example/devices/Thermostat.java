package org.example.devices;

import SmartHome.TemperatureOutOfRangeError;
import com.zeroc.Ice.Current;

public class Thermostat extends Device implements SmartHome.Thermostat {
    private final double MAX_TEMPERATURE = 40.00;
    private final double MIN_TEMPERATURE = 10.00;
    private double temperature;

    public Thermostat() {}

    @Override
    public void setTemperature(double temperature, Current current) throws TemperatureOutOfRangeError {
        if (temperature < MIN_TEMPERATURE || temperature > MAX_TEMPERATURE) {
            throw new TemperatureOutOfRangeError();
        }
        this.temperature = temperature;
        System.out.println("Temperature set to " + temperature);
    }

    @Override
    public double getTemperature(Current current) {
        System.out.println("Current temperature: " + temperature);
        return this.temperature;
    }
}

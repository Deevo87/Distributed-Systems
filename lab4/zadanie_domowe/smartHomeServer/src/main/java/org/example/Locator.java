package org.example;

import com.zeroc.Ice.*;
import com.zeroc.Ice.Object;
import org.example.devices.*;

import java.util.ArrayList;
import java.util.List;

public class Locator implements ServantLocator {
    private final List<String> servants = new ArrayList<>();
    private final String id;

    public Locator(String serverName) {
        this.id = serverName;
    }

    @Override
    public LocateResult locate(Current current) throws UserException {
        String servant = current.id.name;
        ObjectAdapter adapter = current.adapter;
//        if (getServantId(servant).equals(this.id + 10)) {
//            String servantName = servant.substring(0, getFirstServantIdIndex(servant));
        this.servants.add(servant);
        switch (servant) {
            case "CoffeeMaker":
                CoffeeMaker coffeeMaker = new CoffeeMaker();
                adapter.add(coffeeMaker, new Identity(servant, "CoffeeMaker"));
                return new ServantLocator.LocateResult(coffeeMaker, null);
            case "Thermostat":
                Thermostat thermostat = new Thermostat();
                adapter.add(thermostat, new Identity(servant, "Thermostat"));
                return new ServantLocator.LocateResult(thermostat, null);
            case "SmartBlinds":
                SmartBlinds smartBlinds = new SmartBlinds();
                adapter.add(smartBlinds, new Identity(servant, "SmartBlinds"));
                return new ServantLocator.LocateResult(smartBlinds, null);
            case "DailyScheduleBlinds":
                DailyScheduleBlinds dailyScheduleBlinds = new DailyScheduleBlinds();
                adapter.add(dailyScheduleBlinds, new Identity(servant, "DailyScheduleBlinds"));
                return new ServantLocator.LocateResult(dailyScheduleBlinds, null);
            case "AlarmTriggeredBlinds":
                AlarmTriggeredBlinds alarmTriggeredBlinds = new AlarmTriggeredBlinds();
                adapter.add(alarmTriggeredBlinds, new Identity(servant, "AlarmTriggeredBlinds"));
                return new ServantLocator.LocateResult(alarmTriggeredBlinds, null);
            default:
                throw new RuntimeException("Servant with this name does not exists.");

        }
//            System.out.println(servantName);
//        }
    }

    @Override
    public void finished(Current current, Object object, java.lang.Object o) throws UserException {

    }

    @Override
    public void deactivate(String s) {
        System.out.println("Server deactivated: " + s);
    }

    private String getServantId(String servantName) {
        return servantName.substring(getFirstServantIdIndex(servantName));
    }

    private int getFirstServantIdIndex(String servantName) {
        int i = servantName.length();
        while (i > 0 && Character.isDigit(servantName.charAt(i - 1))) {
            i--;
        }
        return i;
    }

    public void servantToString() {
        for (String name : this.servants) {
            System.out.println("Servant name: " + name);
        }
    }


}

package org.example.devices;

import SmartHome.AlarmAlreadyOffError;
import SmartHome.AlarmAlreadyOnError;
import com.zeroc.Ice.Current;

public class AlarmTriggeredBlinds extends SmartBlinds implements SmartHome.AlarmTriggeredBlinds {
    private boolean isAlarmTriggered = false;

    @Override
    public void activateAlarmMode(Current current) throws AlarmAlreadyOnError {
        if (isAlarmTriggered) {
            throw new AlarmAlreadyOnError();
        }
        this.isAlarmTriggered = true;
    }

    @Override
    public void deactivateAlarmMode(Current current) throws AlarmAlreadyOffError {
        if (!isAlarmTriggered) {
            throw new AlarmAlreadyOffError();
        }
        this.isAlarmTriggered = false;
    }
}

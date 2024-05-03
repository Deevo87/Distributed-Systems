package org.example.devices;

import SmartHome.AlreadyWorkingError;
import SmartHome.Info;
import com.zeroc.Ice.Current;

public class Device implements SmartHome.Device {
    private Info info = Info.Working;

    @Override
    public void setInfo(Info info, Current current) throws AlreadyWorkingError {
        if (info == Info.Working) {
            throw new AlreadyWorkingError();
        }
        this.info = info;
    }

    @Override
    public boolean isWorking(Current current) {
        return this.info == Info.Working;
    }

    @Override
    public boolean isFailure(Current current) {
        return this.info == Info.Failure;
    }
}

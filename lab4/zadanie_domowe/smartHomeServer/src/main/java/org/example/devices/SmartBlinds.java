package org.example.devices;

import SmartHome.AlreadyClosedError;
import SmartHome.AlreadyOpenedError;
import SmartHome.AngleOutOfRangeError;
import SmartHome.CoverageOutOfRangeError;
import com.zeroc.Ice.Current;

public class SmartBlinds extends Device implements SmartHome.SmartBlinds {
    private final double SUPPORTED_ANGLE_MAX = 180.00;
    private final double SUPPORTED_ANGLE_MIN = 0.00;
    private final int SUPPORTED_COVERAGE_MAX = 100; //CLOSE
    private final int SUPPORTED_COVERAGE_MIN = 0; //OPEN
    private int coverage = 0;
    private double angle = 90.00;
    private boolean closed = false;

    public SmartBlinds() {}

    @Override
    public void openBlinds(Current current) throws AlreadyOpenedError {
        if (!this.closed) {
            throw new AlreadyOpenedError();
        }
        this.closed = false;
        System.out.println("SmartBlinds opened.");
    }

    @Override
    public void closeBlinds(Current current) throws AlreadyClosedError {
        if (this.closed) {
            throw new AlreadyClosedError();
        }
        this.closed = true;
        System.out.println("SmartBlinds closed.");
    }

    @Override
    public void setAngle(double angle, Current current) throws AngleOutOfRangeError {
        if (angle < SUPPORTED_ANGLE_MIN || angle > SUPPORTED_ANGLE_MAX) {
            throw new AngleOutOfRangeError();
        }
        this.angle = angle;
        System.out.println("SmartBlinds angle changed.");
    }

    @Override
    public void setCustomWindowCoverage(int coverage, Current current) throws CoverageOutOfRangeError {
        if (coverage < SUPPORTED_COVERAGE_MIN || coverage > SUPPORTED_COVERAGE_MAX) {
            throw new CoverageOutOfRangeError();
        }
        this.coverage = coverage;
        System.out.println("SmartBlinds coverage changed.");
    }
}

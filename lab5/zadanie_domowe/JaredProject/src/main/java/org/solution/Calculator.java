package org.solution;

import java.util.List;

public class Calculator {

    public Double add(List<Double> numbers) {
        return numbers.stream().reduce(0.0, Double::sum);
    }

    public double average(List<Double> numbers) {
        return (double) numbers.stream().reduce(0.0, Double::sum) / numbers.size();
    }

    public double power(List<Double> numbers) {
        return Math.pow(numbers.stream().reduce(0.0, Double::sum), 2);
    }
}

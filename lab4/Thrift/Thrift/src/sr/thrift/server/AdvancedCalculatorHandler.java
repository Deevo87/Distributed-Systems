package sr.thrift.server;

import org.apache.thrift.TException;

import sr.rpc.thrift.OperationType;
import sr.rpc.thrift.InvalidArguments;
import sr.rpc.thrift.AdvancedCalculator;

// Generated code

import java.util.List;
import java.util.Set;

public class AdvancedCalculatorHandler implements AdvancedCalculator.Iface {

	int id;

	public AdvancedCalculatorHandler(int id) {
		this.id = id;
	}

	public int add(int n1, int n2) {
		System.out.println("AdvCalcHandler#" + id + " add(" + n1 + "," + n2 + ")");
		//try { Thread.sleep(9000); } catch(java.lang.InterruptedException ex) { }
		System.out.println("DONE");
		return n1 + n2;
	}


	@Override
	public double op(OperationType type, Set<Double> val) throws TException 
	{
		System.out.println("AdvCalcHandler#" + id + " op() with " + val.size() + " arguments");
		
		if(val.size() == 0) {
			throw new InvalidArguments(0, "no data");
		}
		
		double res = 0;
		switch (type) {
		case SUM:
			for (Double d : val) res += d;
			return res;
		case AVG:
			for (Double d : val) res += d;
			return res/val.size();
		case MIN:
			return 0;
		case MAX:
			return 0;
		}
		
		return 0;
	}

	
	@Override
	public int subtract(int num1, int num2) throws TException {
		return num1 - num2;
	}

	@Override
	public List<Integer> sorting(List<Integer> arr) throws TException {
		return arr.stream().sorted().toList();
	}

}


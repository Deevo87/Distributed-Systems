package org.solution;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import io.grpc.stub.StreamObserver;
import org.Executors.ExecutionRequest;
import org.Executors.ExecutionResponse;
import org.Executors.ExecutionServiceGrpc;

import java.lang.reflect.Method;
import java.net.MalformedURLException;
import java.net.URL;
import java.net.URLClassLoader;
import java.nio.file.Path;
import java.nio.file.Paths;

public class ExecutionServiceImpl extends ExecutionServiceGrpc.ExecutionServiceImplBase {
    private String JAR_PATH = "../jars/";

    private final Gson gson = new GsonBuilder()
            .setLenient()
            .create();

    @Override
    public void execute(ExecutionRequest request, StreamObserver<ExecutionResponse> responseObserver) {
        System.out.println("Executing request: " + request);
        ExecutionResponse.Builder responseBuilder = ExecutionResponse.newBuilder();
        Path jarPath = Paths.get(JAR_PATH, request.getJarLocation());
        URL jarUrl;
        try {
            jarUrl = jarPath.toUri().toURL();
            Class c;
            ClassLoader classLoader = Main.class.getClassLoader();
            URLClassLoader urlClassLoader = new URLClassLoader(new URL[]{jarUrl}, classLoader);
            try {
                c = Class.forName(request.getClassName(), true, urlClassLoader);
                System.out.println("Class found: " + c.getName());
            } catch (ClassNotFoundException e) {
                System.out.println("Class not found: " + request.getClassName());
                throw new RuntimeException(e);
            }
            Method method = this.getMethod(c, request.getMethodName(), responseBuilder);
            Object execRes;
            try {
                execRes = this.getExecutionResult(c, method, request.getData(), responseBuilder);
            } catch (Exception e) {
                System.out.println(method);
                System.out.println("Error executing request: " + request);
                throw new RuntimeException(e);
            }
            responseBuilder.setData(gson.toJson(execRes));
        } catch (MalformedURLException e) {
            System.err.println("Malformed URL: " + e.getMessage());
            throw new RuntimeException(e);
        }
        responseObserver.onNext(responseBuilder.build());
        responseObserver.onCompleted();
    }

    private Object getExecutionResult(Class c, Method method, String data, ExecutionResponse.Builder responseBuilder) throws Exception {
        Object result = null;
        Object object = c.getDeclaredConstructor().newInstance();
        Class[] pTypes = method.getParameterTypes();
        if (1 == pTypes.length) {
            result = method.invoke(object, gson.fromJson(data, pTypes[0]));
        } else if (0 == pTypes.length) {
            result = method.invoke(object);
        } else {
            System.out.println("Wrong number of parameters: " + pTypes.length);
        }
        return result;
    }

    private Method getMethod(Class c, String methodName, ExecutionResponse.Builder responseBuilder) {
        Method[] methods = c.getDeclaredMethods();
        System.out.println("Methods found: " + methods.length);
        for (Method method : methods) {
            System.out.println(method.getName());
            if (method.getName().equals(methodName)) {
                Class[] pTypes = method.getParameterTypes();
                System.out.println("Method found: " + methodName);

                try {
                    if (pTypes.length == 1) {
                        Class pType = pTypes[0];
                        return c.getMethod(methodName, new Class[]{pType});
                    } else if (pTypes.length == 0) {
                        return c.getMethod(methodName);
                    }
                } catch (NoSuchMethodException e) {
                    System.out.println("Method not found: " + methodName);
                    break;
                }
            }
        }
        return null;
    }
}

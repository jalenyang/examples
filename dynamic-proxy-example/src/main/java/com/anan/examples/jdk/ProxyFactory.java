package com.anan.examples.jdk;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class ProxyFactory {

    private Object obj;

    public ProxyFactory(Object obj) {
        this.obj = obj;
    }

    public Object getInstance() {
        return Proxy.newProxyInstance(this.obj.getClass().getClassLoader(),
                this.obj.getClass().getInterfaces(), new InvocationHandler() {
                    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                        System.out.println("JDK:调用前日志处理");
                        return method.invoke(obj, args);
                    }
                });
    }
}

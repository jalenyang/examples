package com.anan.examples.cglib.interceptor;

import net.sf.cglib.proxy.MethodInterceptor;
import net.sf.cglib.proxy.MethodProxy;

import java.lang.reflect.Method;


public class LogInterceptor implements MethodInterceptor {
    /**
     * @param o           表示要进行增强的对象
     * @param method      表示拦截的方法
     * @param objects     数组表示参数列表，基本数据类型需要传入其包装类型，如int-->Integer、long-Long、double-->Double
     * @param methodProxy 表示对方法的代理，invokeSuper方法表示对被代理对象方法的调用
     * @return 执行结果
     * @throws Throwable 异常
     */
    public Object intercept(Object o, Method method, Object[] objects, MethodProxy methodProxy) throws Throwable {
        // 注意这里是调用invokeSuper而不是invoke，否则死循环;
        // methodProxy.invokeSuper执行的是原始类的方法;
        // method.invoke执行的是子类的方法;
        before(method.getName());
        Object result = methodProxy.invokeSuper(o, objects);
        after(method.getName());
        return result;
    }

    /**
     * 调用invoke方法之前执行
     */
    private void before(String methodName) {
        System.out.println("调用方法" + methodName + "之【前】的日志处理");
    }


    /**
     * 调用invoke方法之后执行
     */
    private void after(String methodName) {
        System.out.println("调用方法" + methodName + "之【后】的日志处理");
    }
}

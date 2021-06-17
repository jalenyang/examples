package com.anan.examples.jdk;

import com.anan.examples.jdk.dao.IUserDao;
import com.anan.examples.jdk.dao.UserDaoImpl;

public class JdkMain {
    public static void main(String[] args) {
        UserDaoImpl userDao = new UserDaoImpl();
        ProxyFactory proxyFactory = new ProxyFactory(userDao);
        IUserDao userProxy = (IUserDao) proxyFactory.getInstance();
        userProxy.findAllUsers();
    }
}

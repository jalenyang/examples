package com.anan.examples.jdk.dao;

public interface IUserDao {
    void findAllUsers();

    String findUsernameById(int id);
}

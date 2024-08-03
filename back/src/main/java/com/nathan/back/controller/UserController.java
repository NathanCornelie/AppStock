package com.nathan.back.controller;

import com.nathan.back.entity.User;

import com.nathan.back.service.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class UserController {
    @Autowired
    private UsersService service;

    @QueryMapping
    public List<User> users(){
        return service.getUsers();
    }

    @QueryMapping
    public User userById(@Argument Integer id){
        return service.getUserById(id);
    }

    @MutationMapping
    public User createUser(@Argument User user){
        return service.createUser(user);
    }

    @MutationMapping
    public Integer deleteUser(@Argument Integer id){
        return service.deleteUser(id);
    }
    @MutationMapping
    public User updateUser(@Argument User user){
        return service.updateUser(user);
    }

}

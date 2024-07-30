package com.nathan.back.service;


import com.nathan.back.entity.Service;
import com.nathan.back.model.ProductCreate;
import com.nathan.back.repository.ProductRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
@Service
public class UsersService{
    @Autowired
    private UserRepository repository;

    public List<User> getUsers(){
        return repository.findAll();
    }
    public User getUserById(Intger id){
        return repository.findById(id).orElse(null);    
    }

    public User createUser(@Argument User user){
        return repository.save(user);
    }
    public Integer deleteUser(@Argument Integer id){
        repository.deleteById(id);
        return id;
    }
    //TODO : pour la methode update voir quelle synthaxe utiliser:   
}

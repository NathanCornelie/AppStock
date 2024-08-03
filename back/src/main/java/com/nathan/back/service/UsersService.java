package com.nathan.back.service;




import com.nathan.back.entity.User;
import com.nathan.back.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.stereotype.Service;

import java.util.List;
@Service
public class UsersService{
    @Autowired
    private UserRepository repository;

    public List<User> getUsers(){
        return repository.findAll();
    }
    public User getUserById(Integer id){
        return repository.findById(id).orElse(null);    
    }

    public User createUser(@Argument User user){
        return repository.save(user);
    }
    public Integer deleteUser(@Argument Integer id){
        repository.deleteById(id);
        return id;
    }

    public User updateUser(@Argument User user){
        User existingUSer= repository.findById(user.getId()).orElse(null);

        assert  existingUSer != null;
        if(user.getEmail()!=null) existingUSer.setEmail(user.getEmail());
        if(user.getPassword()!=null) existingUSer.setPassword(user.getPassword());
        if(user.getFirstname()!=null) existingUSer.setFirstname(user.getFirstname());
        if(user.getLastname()!=null) existingUSer.setLastname(user.getLastname());

        return repository.save(user);
    }

}

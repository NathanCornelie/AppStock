package com.nathan.back.repository;

import com.nathan.back.entity.User;

import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository  extends JpaRepository<User, Integer> {

}

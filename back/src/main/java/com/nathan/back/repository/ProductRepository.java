package com.nathan.back.repository;

import com.nathan.back.entity.Product;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ProductRepository extends JpaRepository<Product, Integer> {

    public Product findByName(String name);
}

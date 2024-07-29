package com.nathan.back.controller;

import com.nathan.back.entity.Product;
import com.nathan.back.service.ProductService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class ProductController {
    @Autowired
    private ProductService service;

    @QueryMapping(name = "products")
    public List<Product> products(){
        return service.getProducts();
    }

    @QueryMapping
    public Product productById(@Argument Integer id){
        return service.getProductById(id);
    }

    @MutationMapping
    public Product createProduct(@Argument Product product){
        return service.saveProduct(product);
    }

    @MutationMapping
    public Integer deleteProduct(@Argument Integer id){
        return service.deleteProduct(id);
    }
    @MutationMapping
    public Product updateProduct(@Argument Product product){
        return service.updateProduct(product);
    }

}

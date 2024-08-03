package com.nathan.back.service;

import com.nathan.back.entity.Product;

import com.nathan.back.repository.ProductRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProductService {
    @Autowired
    private ProductRepository repository;

    public Product saveProduct(Product product) {
        return repository.save(product);
    };

    public List<Product> saveProducts(List<Product> products) {
        return repository.saveAll(products);
    }
    public List<Product> getProducts(){
        return repository.findAll();
    }
    public Product getProductById(Integer id){
        return repository.findById(id).orElse(null);
    }
    public Product getProductByName(String name){
        return repository.findByName(name);
    }

    public Integer deleteProduct(int id){
        repository.deleteById(id);
        return id;
    }

    public Product updateProduct(Product product){
       Product existingProduct = repository.findById(product.getId()).orElse(null);
        assert existingProduct != null;
        if(product.getName() !=null)existingProduct.setName(product.getName());
        if(product.getCategory() !=null) existingProduct.setCategory(product.getCategory());
        if(product.getDescription()!=null) existingProduct.setDescription(product.getDescription());

        if(product.getPrice()>0)existingProduct.setPrice(product.getPrice());
        existingProduct.setStock(product.getStock());
       return repository.save(existingProduct);

    }

}

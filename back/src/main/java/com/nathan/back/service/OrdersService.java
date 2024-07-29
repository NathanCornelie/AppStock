package com.nathan.back.service;

import com.nathan.back.entity.Order;
import com.nathan.back.entity.Product;
import com.nathan.back.repository.OrderRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.stereotype.Service;

import java.util.List;
@Service
public class OrdersService {
    @Autowired
    private OrderRepository repository;

    public List<Order> getOrders(){
        return repository.findAll();
    }

    public Order getOrderById(Integer id){
        return repository.findById(id).orElse(null);
    }

    public Order createOrder(@Argument Order order){return repository.save(order);}
    public Integer deleteOrder(@Argument Integer id){ repository.deleteById(id);return id;}
    public Order updateOrder(@Argument Order order){
        Order existingOrder = repository.findById(order.getId()).orElse(null);
        assert existingOrder != null;
        existingOrder.setDiscount(order.getDiscount());
        existingOrder.setQuantity(order.getQuantity());
        existingOrder.setDocumentId(order.getDocumentId());
        existingOrder.setItemId(order.getItemId());
        //TODO: a tester !! si on ne specifie pas tous les champs !!
        return repository.save(existingOrder);}
}

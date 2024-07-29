package com.nathan.back.controller;


import com.nathan.back.entity.Order;
import com.nathan.back.service.OrdersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class OrderController {
    @Autowired
    private OrdersService service;

    @QueryMapping
    public List<Order> orders(){
        return service.getOrders();
    }

    @QueryMapping
    public Order ordersById(@Argument Integer id){
        return service.getOrderById(id);
    }

    @MutationMapping
    public Order createOrder(@Argument Order order){
        return service.createOrder(order);
    }

    @MutationMapping
    public Integer deleteOrder(@Argument Integer id){
        return service.deleteOrder(id);
    }

    @MutationMapping
    public Order updateOrder(@Argument Order order){
        return service.updateOrder(order);
    }

}

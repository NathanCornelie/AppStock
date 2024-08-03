package com.nathan.back.controller;

import com.nathan.back.entity.Document;
import com.nathan.back.entity.Order;
import com.nathan.back.service.DocumentsService;
import com.nathan.back.service.OrdersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.graphql.data.method.annotation.SchemaMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class DocumentController {
    @Autowired
    private DocumentsService service;
    private OrdersService ordersService;

    @QueryMapping
    public List<Document> documents(){
        return service.getDocuments();
    }

    @QueryMapping
    public Document documentById(@Argument Integer id){
        return service.getDocumentById(id);
    }

    @MutationMapping
    public Document createDocument(@Argument Document document){
        return service.createDocument(document);
    }

    @MutationMapping
    public Integer deleteDocument(@Argument Integer id){
        return service.deleteDocument(id);
    }
    @MutationMapping
    public Document updateDocument(@Argument Document document){
        return service.updateDocument(document);
    }

    @SchemaMapping
    public List<Order> orders(Document document){
        return ordersService.getOrdersByDocumentId(document.getId());
    }
}

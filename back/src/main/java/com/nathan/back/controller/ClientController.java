package com.nathan.back.controller;

import com.nathan.back.entity.Client;
import com.nathan.back.service.ClientsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.MutationMapping;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class ClientController {
    @Autowired
    private ClientsService service;

    @QueryMapping
    public List<Client> clients(){return service.getClients();}

    @QueryMapping
    public Client clientById(@Argument Integer id){return service.getClientById(id);}

    @MutationMapping
    public Client createClient(@Argument Client client){return service.createClient(client);}

    @MutationMapping
    public Client updateClient(@Argument Client client){return service.updateClient(client);}

    @MutationMapping
    public Integer deleteClient(@Argument Integer id){return service.deleteClient(id);}

}

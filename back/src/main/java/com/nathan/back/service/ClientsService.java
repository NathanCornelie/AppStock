package com.nathan.back.service;

import com.nathan.back.entity.Client;
import com.nathan.back.repository.ClientRepository;
import jakarta.persistence.criteria.CriteriaBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
@Service
public class ClientsService {
    @Autowired
    private ClientRepository repository;
    public List<Client> getClients(){
        return repository.findAll();
    }
    public Client getClientById(int id){
        return repository.findById(id).orElse(null);
    }
    public Client createClient(Client client){
        return repository.save(client);
    }
    public Client updateClient(Client client){
        Client existingClient = repository.findById(client.getId()).orElse(null);

        assert existingClient != null;
        existingClient.setFirstname(client.getFirstname());
        existingClient.setLastname(client.getLastname());
        existingClient.setEmail(client.getEmail());

        return repository.save(existingClient);
    }

    public Integer deleteClient(Integer id){
        repository.deleteById(id);
        return id;
    }
}

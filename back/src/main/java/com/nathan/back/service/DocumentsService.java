package com.nathan.back.service;

import com.nathan.back.entity.Document;
import com.nathan.back.repository.DocumentRepository;
import jdk.dynalink.linker.LinkerServices;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.print.Doc;
import java.util.List;

@Service
public class DocumentsService {
    @Autowired
    private DocumentRepository repository;

    public List<Document> getDocuments(){
        return repository.findAll();
    }
    public Document getDocumentById(Integer id){
        return repository.findById(id).orElse(null);
    }

    public Document createDocument(Document document){
        return repository.save(document);
    }
    public Integer deleteDocument(Integer id){
        repository.deleteById(id);
        return id;}

    public Document updateDocument( Document document){
        Document existingDocument = repository.findById(document.getId()).orElse(null);

        assert existingDocument != null;
        if(document.getType() !=null) existingDocument.setType(document.getType());
        if(document.getDate() != null) existingDocument.setDate(document.getDate());
        if(document.getClientId() >0) existingDocument.setClientId(document.getClientId());
        if(document.getClientId()>0) existingDocument.setClientId(document.getClientId());

        return repository.save(existingDocument);
    }

}

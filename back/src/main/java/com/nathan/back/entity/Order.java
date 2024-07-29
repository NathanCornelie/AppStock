package com.nathan.back.entity;



import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;


@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name= "ORDERS")
public class Order {
    @Id
    @GeneratedValue
    private int id;
    private int documentId;
    private int itemId;
    private int quantity;
    private int discount;
}

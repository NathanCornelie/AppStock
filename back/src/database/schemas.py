from logging.config import listen
from typing import List, Union
from pydantic import BaseModel
import datetime


class ItemBase(BaseModel):
    """Base item model"""
    ref: str
    description: str


class ItemCreate(ItemBase):
    """model used to create an item"""
    pass


class Item(ItemBase):
    prix : float
    quantite: int
    casier: str

    class Config:
        orm_mode = True


class ItemUpdate(BaseModel):
    ref: str | None
    desc: str | None
    quantite : str | None 
    casier : str | None 
    prix: float | None

class ClientBase(BaseModel):
    """Base client model"""
    nom: str
    prenom: str
    adresse: str
    codePostal: int
    remise: int


class ClientCreate(ClientBase):
    """model used to create a client"""
    pass


class Client(ClientBase):
    
    id: int
    actif: bool

    class Config:
        orm_mode = True


class AchatBase(BaseModel):
    ref: str
    quantite: int
    remise: int
    doc_id: str
    id: str


class AchatCreate(AchatBase):
    pass


class Achat(AchatBase):

    class Config:
        orm_mode = True


class DocumentBase(BaseModel):
    doc_id: str
    type: str
    
    client_id: int
    vendeur: str
    date: datetime.date

class DocumentCreate(DocumentBase):
    pass



class Document(DocumentBase):
    client : Client
    achats: List[Achat]
    valide : bool

    class Config:
        orm_mode = True

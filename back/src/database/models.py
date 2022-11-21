
from sqlalchemy import Boolean, Column, Text, ForeignKey, Integer, String, Date
from sqlalchemy.orm import relationship
from .__init__ import Base


class Client(Base):
    __tablename__ = "clients"

    id = Column(Integer, primary_key=True, index=True,nullable=False)
    nom = Column(String, index=True)
    prenom = Column(String, index=True)
    adresse = Column(String, index=True)
    codePostal = Column(Integer, index=True)
    remise = Column(Integer, default=0)
    actif = Column(Boolean,nullable=False)
    docs = relationship("Document",back_populates="client")

class Item(Base):
    __tablename__ = "pieces"

    ref = Column(String, primary_key=True, index=True, nullable=False)
    description = Column(String, index=True)
    quantite = Column(Integer, index=True, default=0)
    casier = Column(String, default="not assigned")
    prix = Column(Integer)
    

class Document(Base):
    __tablename__ = "documents"

    doc_id = Column(Text, primary_key=True, index=True, nullable=False)
    type = Column(Text, nullable=False)
    
    client_id = Column(Integer, ForeignKey("clients.id"), nullable=False)
    client = relationship("Client",back_populates="docs")
    vendeur = Column(Text, nullable=False)
    achats = relationship("Achats",back_populates="doc")
    date = Column(Date)
    valide = Column(Boolean,nullable=False)

class Achats(Base):
    __tablename__ = "achats"

    id = Column(Text,primary_key=True,index=True, nullable=False)
    doc_id = Column(Text,ForeignKey(Document.doc_id ),nullable=False)
    ref = Column(ForeignKey(Item.ref),nullable = False)
    quantite = Column(Integer)
    remise = Column(Integer,ForeignKey(Client.remise))
    doc = relationship("Document",back_populates="achats")
    
from operator import getitem
from pydoc import doc
from sqlalchemy.orm import Session

from . import models, schemas


def get_client(db: Session, client_id: int):
    return db.query(models.Client).filter(models.Client.id == client_id).first()


def get_client_by_first_name(db: Session, client_fname: str):
    return db.query(models.Client).filter(models.Client.Prénom == client_fname).first()


def get_client_by_name(db: Session, client_name: str):
    return db.query(models.Client).filter(models.Client.Nom == client_name).first()


def get_clients(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Client).offset(skip).limit(limit).all()


def create_client(db: Session, client: schemas.ClientCreate) -> models.Client:
    db_client = models.Client(
        actif=True,
        **client.dict(),


    )
    db.add(db_client)
    db.commit()
    db.refresh(db_client)
    return db_client


def get_item(db: Session, item_id: str) -> models.Item | None:
    return db.query(models.Item).filter(models.Item.ref == item_id).first()


def get_items(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Item).offset(skip).limit(limit).all()


def create_item(db: Session, item: schemas.ItemCreate) -> models.Item:
    db_item = models.Item(
        ref=item.ref,
        description=item.description,
        prix =0
    )
    db.add(db_item)
    db.commit()
    db.refresh(db_item)
    return db_item


def update_item(db: Session, item: schemas.ItemUpdate, ref: str) -> models.Item | None:
    db_item = get_item(db, ref)
    if db_item is None:
        return None
    if item.ref is not None:
        db_item.ref = item.ref
    if item.casier is not None: 
        db_item.casier = item.casier
    if item.desc is not None :
        db_item.description = item.desc
    if item.quantite is not None: 
        db_item.quantite = item.quantite
    if item.prix is not None:
        db_item.prix = item.prix
    db.commit()
    return db_item

def achat_item(db:Session,ref:str , qte_achat : int ) -> models.Item:
    db_item = get_item(db,ref)
    if db_item is None: 
        return None 
    if db_item.quantite is not None : 
        db_item.quantite = db_item.quantite - qte_achat
    db.commit()
    return db_item

def get_achat(db: Session, achat_id: str):
    return db.query(models.Achats).filter(models.Achats.id == achat_id).first()


def get_achat_by_client_id(db: Session, client_id: int):
    liste = []
    documents = db.query(models.Document).filter(
        models.Document.client.id == client_id).all()

    for doc in documents:
        for piece in doc.pieces:
            liste.append(piece)
    return liste


def get_achat_by_doc_id(db: Session, doc_id: str) -> list[models.Achats]:
    return db.query(models.Achats).filter(models.Achats.doc_id == doc_id).all()


def get_achat_by_ref(db: Session, ref: str) -> list[models.Achats]:
    return db.query(models.Achats).filter(models.Achats.ref == ref).all()


def create_achat(db: Session, achat: schemas.AchatCreate) -> models.Achats:
    db_achat = models.Achats(
        **achat.dict(),


    )
    db.add(db_achat)
    db.commit()
    db.refresh(db_achat)
    return db_achat


def delete_achat(db: Session, achat_id: str) -> models.Achats:
    db_achat = get_achat(db, achat_id=achat_id)
    if db_achat is None:
        return None

    db.delete(db_achat)
    db.commit()
    return db_achat


def get_doc(db: Session, doc_id: str)->schemas.Document:
    return db.query(models.Document).filter(models.Document.doc_id == doc_id).first()


def get_docs(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.Document).offset(skip).limit(limit).all()


def create_doc(db: Session, doc=schemas.DocumentCreate):
    db_doc = models.Document(
        **doc.dict(),
        valide=False
    )
    db.add(db_doc)
    db.commit()
    db.refresh(db_doc)
    return db_doc

def ajoutAchats(db:Session, doc_id:str , achats:list[schemas.Achat]):
    db_doc = get_doc(db=db , doc_id=doc_id)
    if len(achats):
        db_doc.achats = achats
    db.commit()
    return db_doc

def deleteAchats(db:Session,doc_id:str):
    db_doc = get_doc(db=db,doc_id=doc_id)
    if db_doc is None : 
        return None 
    for achat in db_doc.achats :
         db.delete(achat)
    db.commit()
    return db_doc


def updateDoc(db:Session, achats: list[schemas.Achat],doc_id:str,valide:bool)->schemas.Document:
    doc_db = get_doc(db=db, doc_id=doc_id)
    if doc_db is None : 
        return None 
    
    doc_db.achats = achats
    doc_db.valide = valide
    db.commit()
    return doc_db
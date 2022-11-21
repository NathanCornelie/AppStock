from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from database.models import Achats
from database import schemas, crud, get_db

router = APIRouter(
    prefix="/documents",
    tags=["documents"]
)

@router.post("/",response_model=schemas.Document)
def create_doc(doc:schemas.DocumentCreate, db:Session=Depends(get_db)):
    db_doc = crud.get_doc(db,doc_id=doc.doc_id)
    if db_doc:
        raise HTTPException(status_code=400, detail="Document already registered")
    return crud.create_doc(db=db,doc=doc)

@router.get("/",response_model=list[schemas.Document])
def read_docs(skip:int=0, limit:int= 100, db:Session=Depends(get_db)):
    docs = crud.get_docs(db,skip=skip,limit=limit)
    return docs

@router.get("/{doc_id}",response_model=schemas.Document)
def read_doc(doc_id:str,db:Session=Depends(get_db)):
    db_doc = crud.get_doc(db=db,doc_id=doc_id)
    if db_doc is None :
        raise HTTPException(status_code=404,detail="Document not found")
    return db_doc

@router.post("/achat",response_model=schemas.Achat)
def create_achat(achat:schemas.AchatCreate,db:Session=Depends(get_db)):
    return crud.create_achat(db=db,achat=achat)


@router.get("/{client_id}",response_model=list[schemas.Achat])
def get_docs(client_id:int,db:Session=Depends(get_db)):
    return crud.get_achat_by_client_id(db,client_id)

@router.post("/update/{doc_id}",response_model=schemas.Document)
def addAchats(doc_id:str,achats:list[schemas.Achat],db:Session=Depends(get_db)):
    return crud.ajoutAchats(db=db,doc_id=doc_id,achats=achats)

@router.get("/delete/{doc_id}",response_model=schemas.Document)
def deleteAchats(doc_id:str,db:Session=Depends(get_db)):
    return crud.deleteAchats(db=db,doc_id=doc_id)

@router.get("/update/{doc_id}",response_model=schemas.Document)
def updateDoc(doc_id:str,achats:list[schemas.Achat],db:Session=Depends(get_db)):
    return crud.updateDoc(db=db,doc_id=doc_id,achats=achats) 
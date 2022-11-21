from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from database import schemas, crud, get_db


router = APIRouter(
    prefix="/clients",
    tags=["clients"]
    
)


@router.post("/", response_model = schemas.Client)
def create_client(client:schemas.ClientCreate, db: Session=Depends(get_db)):
    return crud.create_client(db=db, client=client)

@router.get("/",response_model=list[schemas.Client])
def read_clients(skip:int=0,limit:int=100,db:Session = Depends(get_db)):
    clients = crud.get_clients(db,skip=skip,limit=limit)
    return clients

@router.get("/{client_id}",response_model=schemas.Client)
def read_client(client_id:int, db:Session=Depends(get_db)):
    db_client = crud.get_client(db,client_id=client_id)
    if db_client is None :
        raise HTTPException(status_code=404, detail='Client not Found')
    return db_client

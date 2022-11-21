from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from database import schemas, crud, get_db

router = APIRouter(
    prefix="/items",
    tags=["items"]

)


@router.post("/", response_model=schemas.Item)
def create_item(item: schemas.ItemCreate, db: Session = Depends(get_db)):
    db_item = crud.get_item(db, item_id=item.ref)
    if db_item:
        raise HTTPException(status_code=400, detail="Item already registered")
    return crud.create_item(db=db,item=item)


@router.get("/", response_model=list[schemas.Item])
def read_items(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)):
    items = crud.get_items(db, skip=skip, limit=limit)
    return items


@router.get("/{item_id}", response_model=schemas.Item)
def read_item(item_id: str, db:Session = Depends(get_db)):
    db_item = crud.get_item(db,item_id=item_id)
    if db_item is None :
        raise HTTPException(status_code=404, detail='Item not Found')
    return db_item

@router.post("/update/{item_id}",response_model=schemas.Item)
def update_item(item_id:str,item_update: schemas.ItemUpdate, db:Session = Depends(get_db))->schemas.Item:
    return crud.update_item(db=db,item=item_update,ref=item_id)

@router.post("achat/{item_id}",response_model=schemas.Item)
def achat_item(item_id:str , qte:int, db:Session =Depends(get_db))->schemas.Item:
    return crud.achat_item(db=db,ref=item_id,qte_achat=qte)


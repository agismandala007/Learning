from sqlalchemy.orm import Session
from models.shop import Shop

def get_shops(db: Session):
    return db.query(Shop).all()

def get_shop(db: Session, shop_id: int):
    return db.query(Shop).filter(Shop.id == shop_id).first()

def create_shop(db: Session, name: str, description: str, price: int, inventory: int, images: str):
    shop = Shop(name=name, description=description, price=price, inventory=inventory, images=images)
    db.add(shop)
    db.commit()
    db.refresh(shop)
    return shop

def update_shop(db: Session, shop_id: int, name: str = None, description: str = None, price: int = None, inventory: int = None, images: str = None):
    shop = db.query(Shop).filter(Shop.id == shop_id).first()
    if shop:
        if name:
            shop.name = name
        if description:
            shop.description = description
        if price:
            shop.price = price
        if inventory:
            shop.inventory = inventory
        if images:
            shop.images = images
        db.commit()
        db.refresh(shop)
    return shop

def delete_shop(db: Session, shop_id: int):
    shop = db.query(Shop).filter(Shop.id == shop_id).first()
    if shop:
        db.delete(shop)
        db.commit()
        return True
    return False
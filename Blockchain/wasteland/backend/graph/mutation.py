import strawberry
from strawberry.file_uploads import Upload
from schemas.shop import ShopType
from crud.shop import create_shop, update_shop, delete_shop
from db.database import get_db

@strawberry.type
class Mutation:
    @strawberry.mutation
    def create_shop(
        self,
        name: str,
        description: str,
        price: int,
        inventory: int,
        images: Upload
    ) -> ShopType:
        db = next(get_db())
        
        file_location = f"../db/img/{images.filename}"
        with open(file_location, "wb") as buffer:
            buffer.write(images.read())
        
        shop = create_shop(db, name, description, price, inventory, file_location)
        return ShopType(
            id=shop.id,
            name=shop.name,
            description=shop.description,
            price=shop.price,
            inventory=shop.inventory,
            images=shop.images
        )

    @strawberry.mutation
    def update_shop(
        self,
        shop_id: int,
        name: str = None,
        description: str = None,
        price: int = None,
        inventory: int = None,
        images: Upload = None
    ) -> ShopType:
        db = next(get_db())
        
        file_location = None
        if images:
            file_location = f"../db/img/{images.filename}"
            with open(file_location, "wb") as buffer:
                buffer.write(images.read())
                
        shop = update_shop(db, shop_id, name, description, price, inventory, file_location)
        return ShopType(
            id=shop.id,
            name=shop.name,
            description=shop.description,
            price=shop.price,
            inventory=shop.inventory,
            images=shop.images
        )

    @strawberry.mutation
    def delete_shop(self, shop_id: int) -> bool:
        db = next(get_db())
        return delete_shop(db, shop_id)
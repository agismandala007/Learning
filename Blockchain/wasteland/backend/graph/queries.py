import strawberry
from schemas.shop import ShopType
from crud.shop import get_shops, get_shop
from db.database import get_db

@strawberry.type
class Query:
    @strawberry.field
    def shops(self) -> list[ShopType]:
        db = next(get_db())
        shops = get_shops(db)
        return [ShopType(
            id=shop.id,
            name=shop.name,
            description=shop.description,
            price=shop.price,
            inventory=shop.inventory,
            images=shop.images
        ) for shop in shops]

    @strawberry.field
    def shop(self, shop_id: int) -> ShopType:
        db = next(get_db())
        shop = get_shop(db, shop_id)
        if shop:
            return ShopType(
                id=shop.id,
                name=shop.name,
                description=shop.description,
                price=shop.price,
                inventory=shop.inventory,
                images=shop.images
            )
        return None
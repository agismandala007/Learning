import strawberry

@strawberry.type
class ShopType:
    id: str
    name: str
    description: str
    price: int
    inventory: int
    images: str
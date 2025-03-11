from sqlalchemy import Column, Integer, String, Text
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()

class Shop(Base):
    __tablename__ = "shop"
    id = Column(String(255), primary_key=True, index=True)
    name = Column(String(255), index=True)
    description = Column(Text, index=True)
    price = Column(Integer, index=True)
    inventory = Column(Integer, index=True)
    images = Column(String(255), index=True)
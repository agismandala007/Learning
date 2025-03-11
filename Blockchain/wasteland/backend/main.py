from fastapi import FastAPI
import strawberry
from strawberry.fastapi import GraphQLRouter
from graph.queries import Query
from graph.mutation import Mutation
from fastapi.middleware.cors import CORSMiddleware


schema = strawberry.Schema(query=Query, mutation=Mutation)

graphql_app = GraphQLRouter(schema)

app = FastAPI()

origins = [
    "http://localhost",
    "http://localhost:3000",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
app.include_router(graphql_app, prefix="/graphql")
import yaml
from pydantic import BaseModel
from typing import List, Optional
import constants as cte


class Item(BaseModel):
    name: str
    cmd: str
    args: Optional[List[str]] = []


class Category(BaseModel):
    name: str
    items: Optional[List[Item]] = []
    cats: Optional[List['Category']] = []  # Use forward reference


Category.model_rebuild()  # Resolve forward references


class AppList(BaseModel):
    cats: List[Category]
    items: Optional[List[Item]] = []


def load_app_list(app_list_file):
    with open(app_list_file, 'r') as file:
        content = file.read()
    return AppList(**yaml.safe_load(content)).model_dump()


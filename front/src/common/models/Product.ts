export class CreateProduct  {
    name: string=""
    category:string =""
    price:number=0
    stock:number=0

    constructor(name: string="", category:string = "", price:number=0, stock:number=0) {
        this.name = name;
        this.category = category;
        this.price = price;
        this.stock = stock;
    }

}

export class Product extends CreateProduct {
    id : string =""

}
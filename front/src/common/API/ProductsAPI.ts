import {Product} from "@models/Product.ts";

export default abstract class {
    static async GetProducts(): Promise<Product[]> {
        const res = await fetch(`${import.meta.env.VITE_API_URL}/products`, {
            method: "GET"
        })

        const json = await res.json()
        if (json.error) {
            throw new Error(json.error.message)
        } else {
            return json ? json as Product[] : []
        }
    }

}
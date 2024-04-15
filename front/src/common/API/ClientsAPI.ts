import {Client,CreateClient} from "@models/Client.ts";


export default abstract class{
    static async GetClients() : Promise<Client[]>{
        const res = await fetch(`${import.meta.env.VITE_API_URL}/clients` , {
            method : "GET"
        })

        const json = await res.json()
        if(json.error){
            throw new Error(json.error.message)
        }else{
            return json ? json as Client[] : []
        }
    }
    static async CreateClient(client : CreateClient) : Promise<Client|null>{
        const res = await fetch(`${import.meta.env.VITE_API_URL}/clients`, {
            method : "POST",
            body : JSON.stringify(client)
        })
        const json = await res.json()
        if(json.error){
            throw new Error(json.error.message)

        }else {
            return json ? json as Client : null
        }

    }
}
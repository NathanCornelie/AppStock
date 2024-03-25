import Document from "@models/Document.ts";


export default abstract class{
    static async GetDocumentsByUserId(userId : string) : Promise<Document[]>{
        const res = await fetch(`${import.meta.env.VITE_API_URL}/documents/users/${userId}` , {
            method : "GET"
        })
        const json = await res.json()
        if(json.error){
            throw new Error(json.error.message)
        }else{
            return json ? json as Document[] : []
        }
    }
}
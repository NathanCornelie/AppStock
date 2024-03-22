import Document from "@models/Document.ts";


export default abstract class{
    async GetDocumentByUserId(userId : string) : Promise<Document[]>{
        const res = await fetch(`http://localhost:8080/documents/user/${userId}` , {
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
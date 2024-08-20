import Document from "@models/Document.ts";


export default abstract class{
    static async GetDocumentsByUserId(userId : string) : Promise<Document[]>{
        const res = await fetch(`${import.meta.env.VITE_API_URL}/documents/users/${userId}` , {
            method : "GET"
        })

        const json = await res.json()
        console.log(json)
        if(json.error){
            throw new Error(json.error.message)
        }else{
            return json ? json as Document[] : []
        }
    }

    // * JAVA API
    static async GetDocumentsByUserId2(userId : string) : Promise<Document[]>{
        const res = await fetch(`${import.meta.env.VITE_API_URL}/api` , {
            method : "POST",
            headers: {
                "Content-Type": "application/json"
            },

            body: JSON.stringify({
                query: `{
                      documents {
                        id
                        type
                        date
                        clientId
                        userId
                        
                      }
                    }`
            })
        })
        console.log(userId)
        const json = await res.json()
        console.log(json)
        if(json.error){
            throw new Error(json.error.message)
        }else{
            return json ? json.data.documents as Document[] : []
        }
    }
}
class Document  {
    id: string = ""
    date : string = (new Date()).toISOString()
    clientId : string=""
    userId : string=""
    type : string=""
    commands : Command[]=[]

     public getNumberOfProducts() {
        return this.commands.map(c=>c.quantity).reduce((a,b)=>a+b)
    }
}
export default Document
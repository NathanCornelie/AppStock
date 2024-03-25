type Document  ={
    id: string;
    date : string;
    clientId : string;
    userId : string
    type : string
    commands : Command[]
}
export default Document
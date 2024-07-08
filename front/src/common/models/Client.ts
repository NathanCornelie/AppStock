
export class CreateClient {
    name:string=""
    lastname:string=""
    email:string=""
    phoneNumber:string=""
    constructor(name:string, lastname:string, email:string, phoneNumber:string) {
        this.name = name;
        this.lastname = lastname;
        this.email = email;
        this.phoneNumber = phoneNumber;
    }
}
export class Client extends CreateClient{
    id: string = ""

}

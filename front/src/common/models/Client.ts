
export class CreateClient {
    firstname:string=""
    lastname:string=""
    email:string=""
    phoneNumber:string=""
    constructor(firstname:string, lastname:string, email:string, phoneNumber:string) {
        this.firstname = firstname;
        this.lastname = lastname;
        this.email = email;
        this.phoneNumber = phoneNumber;
    }
}
export class Client extends CreateClient{
    id: string = ""

}

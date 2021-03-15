import { GenericClass } from './genericClass';

export class Customer extends GenericClass {
    public CustomerID;
    public customer;
    constructor() {
        super('Customer');
        this.init();
    }

    public init() {}
}
export class ApiResponses {
    public CustomerFetch = {
        CustNodeID: {
            mandatory: true,
            name: 'Customer ID',
            type: 'string'
        },
        TaxationInfo: {
            mandatory: true,
            name: 'Taxation info',
            type: 'string'
        },
    };
}
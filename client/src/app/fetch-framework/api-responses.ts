export class ApiResponses {
    // Objects returned by ClassFetchDetails AND FactFetchDetails API's
    public caseClassFactFetch = {
        CaseID : {
            readonly: true,
            mandatory: true,
            width: 100,
            type: 'string'
        },
        StartDate : {
            readonly: true,
            mandatory: true,
            width: 100,
            type: 'date'
        }
    };
    public custAcctClassFactFetch = {
        CustNodeID: {
            readonly: true,
            mandatory: true,
            name: 'Customer Node ID',
            type: 'string'
        },
        TaxationInfo: {
            readonly: true,
            mandatory: true,
            name: 'Taxation info',
            type: 'string'
        },
    };
}
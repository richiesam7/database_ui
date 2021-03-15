export class Fact {

    public Name: string;
    public DisplayName: string;
    public FactKey: string;
    public Type: string;
    public ObjectType: string;
    public Description: string;
    public SecurityGroup: string;
    public DerivationType: number;
    public ExpiryDuration: number;
    public EvalTiming: number;
    public TableName: string;
    public ColumnName: string;
    public RowKey: string;
    public InitialState: number;
    public Usage: number;
    public AttributeFlags: number;
    public MaxLength: number;
    public Ceiling: string;
    public RefClass: any;
    public DisplayFormat: string;
    public EditMask: string;
    public Icon: string;
    public Local_Fact: boolean;
    public Foreign_Fact: boolean;
    public AssocXref_Fact: any;
    public Statuses: number;

    // Fact State
    public Optional: boolean;
    public OptionalReadOnly: boolean;
    public NotAllowed: boolean;
    public NotAllowedReadOnly: boolean;
    public Mandatory: boolean;
    public MandatoryReadOnly: boolean;
    public Null: boolean;
    public Encrypted: boolean;
    public Masked: boolean;
    
    // To handle validations
    public rules: any[];
}
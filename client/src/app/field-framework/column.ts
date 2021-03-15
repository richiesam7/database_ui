export class Column {
    public Name: string;
    public Type: string;
    public Description: string;
    public Mandatory: boolean;
    // To handle validations
    public rules: any[];
}
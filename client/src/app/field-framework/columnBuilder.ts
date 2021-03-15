import { Column } from "./column";

export class ColumnBuilder {
    public validations = [];
    public column: Column = new Column();

    constructor() {}

    public name(name: string) {
        this.column.Name = name;
        return this;
    }

    public type(type: string) {
        this.column.Type = type;
        return this;
    }

    public mandatory() {
        this.column.Mandatory = true;
        this.validations.push('mandatory');
        return this;
    }

    public description(description: string) {
        this.column.Description = description;
        return this;
    }
}

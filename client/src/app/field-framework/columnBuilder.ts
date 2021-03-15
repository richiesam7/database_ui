import { Fact } from "./column";

export class ColumnBuilder {
    public validations = [];
    public fact: Fact = new Fact();

    constructor() {
    }

    public name(name: string) {
        this.fact.Name = name;
        return this;
    }

    public type(type: string) {
        this.fact.Type = type;
        return this;
    }

    public mandatory() {
        this.fact.Mandatory = true;
        this.validations.push('mandatory');
        return this;
    }

    public key(key: string) {
        this.fact.FactKey = key;
        return this;
    }

    public displayName(displayName: string) {
        this.fact.DisplayName = displayName;
        return this;
    }

    public description(description: string) {
        this.fact.Description = description;
        return this;
    }

    public securityGroup(securityGroup: any) {
        this.fact.SecurityGroup = securityGroup;
        return this;
    }

    public readonly() {
        this.fact.MandatoryReadOnly = true;
        return this;
    }

    /* 
    Unfinished business
    public label(label: string) {
        this.fact.label = label;
        return this;
    }
    public version(version: string) {
        this.fact.version = version;
        return this;
    }
    public visible() {
        this.fact. = true;
        return this;
    }
    public unique() {
        this.fact.unique = true;
        this.validations.push('unique');
        return this;
    }
    public width(displayWidth: number) {
        this.fact. = displayWidth;
        return this;
    }

    public readOnly() {
        this.fact.MandatoryReadOnly = true;
        return this;
    }
    */
}

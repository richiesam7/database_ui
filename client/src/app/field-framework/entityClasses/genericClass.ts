import { ApiResponses } from 'src/app/fetch-framework/api-responses';
import { ColumnBuilder } from '../columnBuilder';

// All entity types inherit from here
export class GenericClass {
    public entity;
    public associatedColumns;
    public columnBuilder: ColumnBuilder;
    private apiResponses;

    constructor(className: string) {
        this.entity = [];
        this.setAssociatedColumns();        
        this.columnBuilder = new ColumnBuilder();
        this.setClassEntity();
    }

    public setAssociatedColumns() {
        this.apiResponses = new ApiResponses();
        this.associatedColumns = this.apiResponses.CustomerFetch; // parameterize
    }

    public setClassEntity() {
        for (let columnName in this.associatedColumns) {
            this.entity[columnName] = [];
            for (let attributeName in this.associatedColumns[columnName]) {
                this.entity[columnName][attributeName] = this.columnBuilder[attributeName](this.associatedColumns[columnName][attributeName]);
            };
            this.entity[columnName]['value'] = null;
        };
    }

    public getField(columnName: string) {
        return this.entity[columnName];
    }

    public getEntity() {
        return this.entity;
    }
}
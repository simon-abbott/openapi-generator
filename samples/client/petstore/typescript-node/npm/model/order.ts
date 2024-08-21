/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

/**
* An order for a pets from the pet store
*/
export class Order {
    'id'?: number;
    'petId'?: number;
    'quantity'?: number;
    'shipDate'?: Date;
    /**
    * Order Status
    */
    'status'?: Order.StatusEnum;
    'complete'?: boolean = false;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "number"
        },
        {
            "name": "petId",
            "baseName": "petId",
            "type": "number"
        },
        {
            "name": "quantity",
            "baseName": "quantity",
            "type": "number"
        },
        {
            "name": "shipDate",
            "baseName": "shipDate",
            "type": "Date"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "Order.StatusEnum"
        },
        {
            "name": "complete",
            "baseName": "complete",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return Order.attributeTypeMap;
    }
}

export namespace Order {
    export type StatusEnum = 'placed' | 'approved' | 'delivered';

    export const StatusEnum = {
        Placed: 'placed' as StatusEnum,
        Approved: 'approved' as StatusEnum,
        Delivered: 'delivered' as StatusEnum
    };
}

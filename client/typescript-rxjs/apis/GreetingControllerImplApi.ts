// tslint:disable
/**
 * OpenAPI definition
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: v0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { Observable } from 'rxjs';
import { BaseAPI, throwIfNullOrUndefined, encodeURI, OperationOpts, RawAjaxResponse } from '../runtime';
import {
    ErrorResponse,
} from '../models';

export interface GreetHumanRequest {
    name: string;
}

/**
 * no description
 */
export class GreetingControllerImplApi extends BaseAPI {

    /**
     */
    greetHuman({ name }: GreetHumanRequest): Observable<string>
    greetHuman({ name }: GreetHumanRequest, opts?: OperationOpts): Observable<RawAjaxResponse<string>>
    greetHuman({ name }: GreetHumanRequest, opts?: OperationOpts): Observable<string | RawAjaxResponse<string>> {
        throwIfNullOrUndefined(name, 'name', 'greetHuman');

        return this.request<string>({
            url: '/greet/{name}'.replace('{name}', encodeURI(name)),
            method: 'GET',
        }, opts?.responseOpts);
    };

    /**
     */
    greetWorld(): Observable<string>
    greetWorld(opts?: OperationOpts): Observable<RawAjaxResponse<string>>
    greetWorld(opts?: OperationOpts): Observable<string | RawAjaxResponse<string>> {
        return this.request<string>({
            url: '/greet',
            method: 'GET',
        }, opts?.responseOpts);
    };

}

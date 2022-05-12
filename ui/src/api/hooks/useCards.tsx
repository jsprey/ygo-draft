import {useQuery, UseQueryResult} from "react-query";
import {Card, Deck} from "../CardModel";
import {PUBLIC_URL} from "../../index";
import {QueryKey} from "react-query/types/core/types";

export function useCard(id: number, queryOptions: any = {}): UseQueryResult<Card> {
    return useMagicMethod<Card>(["card", id], `cards/${id}`, queryOptions)
}

export function useRandomCards(clientID: string, size: number, queryOptions: any = {}): UseQueryResult<Deck> {
    return useMagicMethod<Deck>(["random", clientID], `randomCards?size=${size}`, queryOptions)
}

function useMagicMethod<GenericJsonType>(queryKey: QueryKey, apiV1Path: string, queryOptions: any = {}): UseQueryResult<GenericJsonType> {
    return useQuery(queryKey, () => {
        return new Promise<GenericJsonType>((resolve, reject) => {
            return fetch(`${PUBLIC_URL}/api/v1/${apiV1Path}`).then(function (response) {
                if (!response.ok) {
                    return response.json().then(res => {
                            reject(res)
                        }
                    )
                }
                return response.json().then(function (jsonContent) {
                    let content: GenericJsonType = {} as GenericJsonType
                    if (!jsonContent) {
                        return content
                    }

                    let cardJson: GenericJsonType = jsonContent
                    return resolve(cardJson)
                })
            }).catch((e) => {
                reject(e);
            })
        })
    }, queryOptions);
}

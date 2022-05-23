import {useQuery, UseQueryResult} from "react-query";
import {Card, Deck} from "../CardModel";
import {PUBLIC_URL} from "../../index";
import {QueryKey} from "react-query/types/core/types";
import {CardFilter, FilterToQuery} from "../CardFilter";

export function useCard(id: number, queryOptions: any = {}): UseQueryResult<Card> {
    return useMagicMethod<Card>(["card", id], `cards/${id}`, new Map<string, string>(), queryOptions)
}

export function MapToQuery(queryMap: Map<string, string>): string {
    let result: string = "?"
    queryMap.forEach((value, key) => {
        if (result !== "?") {
            result += "&"
        }
        result += key + "=" + value
    })

    return result
}

export function useRandomCards(clientID: string, size: number, cardFilter: CardFilter, queryOptions: any = {}): UseQueryResult<Deck> {
    const queryParams = FilterToQuery(cardFilter)
    queryParams.set("size", String(size))

    return useMagicMethod<Deck>(["random", clientID], `randomCards`, queryParams, queryOptions)
}

function useMagicMethod<GenericJsonType>(queryKey: QueryKey, apiV1Path: string, queryMap: Map<string, string>, queryOptions: any = {}): UseQueryResult<GenericJsonType> {
    const queryParams = MapToQuery(queryMap)

    return useQuery(queryKey, () => {
        return new Promise<GenericJsonType>((resolve, reject) => {
            return fetch(`${PUBLIC_URL}/api/v1/${apiV1Path}${queryParams}`).then(function (response) {
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

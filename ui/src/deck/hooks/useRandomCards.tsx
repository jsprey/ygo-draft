import {useQuery, UseQueryResult} from "react-query";

const contextPath = process.env.PUBLIC_URL || "";

export type Deck = {
    cards: number[]
}

export function useRandomCards(name: string, size: number): UseQueryResult<Deck> {
    return useQuery(["cards", "decks", name], () => {
        return new Promise<Deck>((resolve, reject) => {
            return fetch(`http://localhost:8080/api/v1/randomCards?size=${size}`).then(function (response) {
                if (!response.ok) {
                    return response.json().then(res => {
                            reject(res)
                        }
                    )
                }

                return response.json().then(function (jsonContent) {
                    console.log(jsonContent)
                    let content: Deck = {} as Deck
                    if (!jsonContent) {
                        return content
                    }

                    let cardJson: Deck = jsonContent
                    return resolve(cardJson)
                })
            }).catch((e) => {
                reject(e);
            })
        })
    });
}
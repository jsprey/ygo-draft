import {useQuery, UseQueryResult} from "react-query";

const contextPath = process.env.PUBLIC_URL || "";

export type CardSet = {
    set_name: string;
    set_code: string;
    set_rarity: string;
    set_rarity_code: string;
    set_price: string;
}

export type CardImage = {
    id: number;
    image_url: string;
    image_url_small: string;
}

export type CardPrice = {
    cardmarket_price: string;
    tcgplayer_price: string;
    ebay_price: string;
    amazon_price: string;
    coolstuffinc_price: string;
}

export type Card = {
    id: number;
    name: string;
    type: string;
    desc: string;
    atk: number;
    def: number;
    level: number;
    race: string;
    attribute: string;
    card_sets: CardSet[];
    card_images: CardImage[];
    card_prices: CardPrice[];
}

export function useRandomCard(identifier: string): UseQueryResult<Card> {
    return useQuery(["api", identifier], () => {
        return new Promise<Card>((resolve, reject) => {
            return fetch(`http://localhost:8080/api/v1/randomCard`).then(function (response) {
                if (!response.ok) {
                    return response.json().then(res => {
                            reject(res)
                        }
                    )
                }
                return response.json().then(function (jsonContent) {
                    let content: Card = {} as Card
                    if (!jsonContent) {
                        return content
                    }

                    let cardJson: Card = jsonContent
                    return resolve(cardJson)
                })
            }).catch((e) => {
                reject(e);
            })
        })
    });
}


export function useCard(id: string): UseQueryResult<Card> {
    return useQuery(["api", id], () => {
        return new Promise<Card>((resolve, reject) => {
            return fetch(`http://localhost:8080/api/v1/cards/${id}`).then(function (response) {
                if (!response.ok) {
                    return response.json().then(res => {
                            reject(res)
                        }
                    )
                }
                return response.json().then(function (jsonContent) {
                    let content: Card = {} as Card
                    if (!jsonContent) {
                        return content
                    }

                    let cardJson: Card = jsonContent
                    return resolve(cardJson)
                })
            }).catch((e) => {
                reject(e);
            })
        })
    });
}

import {Deck} from "./CardModel";
import {PUBLIC_URL} from "../index";

export function getRandomCards(size: number): Promise<Deck> {
    return new Promise<Deck>((resolve, reject) => {
        return fetch(`${PUBLIC_URL}/api/v1/cards/random?size=${size}`).then(function (response) {
            if (!response.ok) {
                return response.json().then(res => {
                        reject(res)
                    }
                )
            }

            return response.json().then(function (jsonContent) {
                let content: Deck = {cards:[]} as Deck
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
}
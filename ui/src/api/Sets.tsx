import {Card} from "./CardModel";

export type CardSet = {
    set_name: string;
    set_code: string;
    set_rarity: string;
    set_rarity_code: string;
}

export type SetList = {
    sets: CardSet[];
}

export type SetWithCards = {
    set: CardSet;
    cards: Card[]
}

export function sortSets(list: CardSet[]): CardSet[] {
    list = list.sort((a, b) => a.set_name.localeCompare(b.set_name))
    return list
}
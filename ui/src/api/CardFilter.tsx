import {CardType, GetCardTypeString} from "./CardType";
import {CardSet} from "./Sets";

export type CardFilter = {
    types: CardType[]
    sets: CardSet[]
}

export const MainDeckFilter = {
    types: getMainCardTypes(),
    sets: [] as CardSet[]
}

export const ExtraDeckFilter = {
    types: getExtraCardTypes(),
    sets: [] as CardSet[]
}

function getMainCardTypes(): CardType[] {
    return Object.values(CardType).filter((type) => {
        return type < 100
    }) as CardType[]
}

function getExtraCardTypes(): CardType[] {
    return Object.values(CardType).filter((type) => {
        return type >= 100
    }) as CardType[]
}

export function FilterToQuery(filter: CardFilter): Map<string, string> {
    const myMap = new Map<string, string>()
    if (filter?.types?.length > 0) {
        myMap.set("types", getTypesAsQueryParam(filter.types))
    }
    if (filter?.sets?.length > 0) {
        myMap.set("sets", getSetsAsQueryParam(filter.sets))
    }
    return myMap
}

function getTypesAsQueryParam(cardTypes: CardType[]): string {
    if (cardTypes === undefined) return ""

    let result = ""
    cardTypes.forEach((value, index) => {
        result += encodeURIComponent(GetCardTypeString(value))
        if (index < cardTypes.length - 1) {
            result += ","
        }
    })
    return result
}

function getSetsAsQueryParam(cardSets: CardSet[]): string {
    if (cardSets === undefined) return ""

    let result = ""
    cardSets.forEach((value, index) => {
        result += encodeURIComponent(value.set_name)
        if (index < cardSets.length - 1) {
            result += ","
        }
    })
    
    return result
}
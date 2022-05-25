import {CardType, GetCardTypeString} from "./CardType";

export type CardFilter = {
    types: CardType[]
}

export const MainDeckFilter = {
    types: getMainCardTypes()
}

export const ExtraDeckFilter = {
    types: getExtraCardTypes()
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
    return myMap
}

function getTypesAsQueryParam(cardTypes: CardType[]): string {
    if (cardTypes === undefined) return ""

    let result = ""
    cardTypes.forEach((value, index) => {
        result += GetCardTypeString(value)
        if (index < cardTypes.length - 1) {
            result += ","
        }
    })
    return result
}
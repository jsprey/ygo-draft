import {CardType, GetCardType} from "./CardType";
import {CardSet} from "./Sets";

export type Card = {
    id: number;
    name: string;
    type: string;
    desc: string;
    atk: number;
    def: number;
    level: number;
    race: string;
    sets: string;
    attribute: string;
    card_sets: CardSet[];
}

export type Deck = {
    cards: Card[]
}

export function SortDeck(deck: Deck): Deck {
    if (deck === undefined) return deck

    let newDeck: Deck = {cards: [] as Card[]} as Deck
    newDeck.cards = SortCards(deck.cards)
    return newDeck
}

export function SortCards(cards: Card[]): Card[] {
    let sortedCards: Card[] = [] as Card[]

    const values2 = Object.values(CardType).filter((v) => !isNaN(Number(v)));
    values2.forEach((value) => {
        let filteredCards = FilterByType(cards, [value as CardType])

        filteredCards = filteredCards.sort(SortAscending)
        filteredCards.forEach(value1 => sortedCards.push(value1))
    });

    return sortedCards
}

export function SortAscending(a: Card, b: Card) {
    return a.id - b.id
}

export function FilterByType(cards: Card[], types: CardType[]): Card[] {
    if (cards === undefined) return [] as Card[]

    return cards.filter(card => {
        for (const typesKey in types) {
            if (GetCardType(card) === types[typesKey]) {
                return true
            }
        }

        return false
    })
}

export function FilterByMainCards(cards: Card[]): Card[] {
    if (cards === undefined) return [] as Card[]

    return cards.filter(card => {
        return GetCardType(card) < 100
    })
}

export function FilterByExtraCards(cards: Card[]): Card[] {
    if (cards === undefined) return [] as Card[]

    return cards.filter(card => {
        return GetCardType(card) >= 100
    })
}

export function ToYdkFileString(deck: Deck): string {
    let deckYdkFileString = "#main\n"

    let mainDeckCards = FilterByMainCards(deck.cards)
    for (const cardKey in mainDeckCards) {
        deckYdkFileString += mainDeckCards[cardKey].id + "\n"
    }

    deckYdkFileString += "\n#extra\n"
    let extraDeckCards = FilterByExtraCards(deck.cards)
    for (const cardKey in extraDeckCards) {
        deckYdkFileString += extraDeckCards[cardKey].id + "\n"
    }

    return deckYdkFileString
}
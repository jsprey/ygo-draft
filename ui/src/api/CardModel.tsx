
export type CardSet = {
    set_name: string;
    set_code: string;
    set_rarity: string;
    set_rarity_code: string;
    set_price: string;
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
}

export type Deck = {
    cards: Card[]
}

export enum CardType {
    Unknown,
    NormalMonster,
    NormalTunerMonster,
    EffectMonster,
    FlipEffectMonster,
    FlipTunerEffectMonster,
    GeminiMonster,
    PendulumEffectMonster,
    PendulumFlipEffectMonster,
    PendulumNormalMonster,
    PendulumTunerEffectMonster,
    RitualEffectMonster,
    RitualMonster,
    ToonMonster,
    SpiritMonster,
    TunerMonster,
    UnionEffectMonster,
    SkillCard,
    SpellCard,
    TrapCard,
    FusionMonster = 100,
    LinkMonster,
    PendulumEffectFusionMonster,
    SynchroMonster,
    SynchroPendulumEffectMonster,
    SynchroTunerMonster,
    XYZMonster,
    XYZPendulumEffectMonster,
}

function GetCardType(card: Card): CardType {
    switch (card.type) {
        case "Normal Monster":
            return CardType.NormalMonster
        case "Effect Monster":
            return CardType.EffectMonster
        case "Flip Effect Monster":
            return CardType.FlipEffectMonster
        case "Flip Tuner Effect Monster":
            return CardType.FlipTunerEffectMonster
        case "Gemini Monster":
            return CardType.GeminiMonster
        case "Normal Tuner Monster":
            return CardType.NormalTunerMonster
        case "Pendulum Effect Monster":
            return CardType.PendulumEffectMonster
        case "Pendulum Flip Effect Monster":
            return CardType.PendulumFlipEffectMonster
        case "Pendulum Normal Monster":
            return CardType.PendulumNormalMonster
        case "Pendulum Tuner Effect Monster":
            return CardType.PendulumTunerEffectMonster
        case "Ritual Effect Monster":
            return CardType.RitualEffectMonster
        case "Ritual Monster":
            return CardType.RitualMonster
        case "Skill Card":
            return CardType.SkillCard
        case "Spell Card":
            return CardType.SpellCard
        case "Spirit Monster":
            return CardType.SpiritMonster
        case "Trap Card":
            return CardType.TrapCard
        case "Toon Monster":
            return CardType.ToonMonster
        case "Tuner Monster":
            return CardType.TunerMonster
        case "Union Effect Monster":
            return CardType.UnionEffectMonster
        case "Fusion Monster":
            return CardType.FusionMonster
        case "Link Monster":
            return CardType.LinkMonster
        case "Pendulum Effect Fusion Monster":
            return CardType.PendulumEffectFusionMonster
        case "Synchro Monster":
            return CardType.SynchroMonster
        case "Synchro Pendulum Effect Monster":
            return CardType.SynchroPendulumEffectMonster
        case "Synchro Tuner Monster":
            return CardType.SynchroTunerMonster
        case "XYZ Monster":
            return CardType.XYZMonster
        case "XYZ Pendulum Effect Monster":
            return CardType.XYZPendulumEffectMonster
        default:
            return CardType.Unknown
    }
}

export function SortDeck(deck: Deck): Deck {
    let newDeck:Deck = {cards: [] as Card[]} as Deck
    newDeck.cards = SortCards(deck.cards)
    return newDeck
}

export function SortCards(cards: Card[]): Card[] {
    let sortedCards:Card[] = [] as Card[]

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
    return cards.filter(card => {
        return GetCardType(card) < 100
    })
}

export function FilterByExtraCards(cards: Card[]):Card[] {
    return cards.filter(card => {
        return GetCardType(card) >= 100
    })
}

export function ToYdkFileString(deck: Deck): string {
    let deckYdkFileString= "#main\n"

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
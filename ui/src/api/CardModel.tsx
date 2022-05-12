
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

export function SortDeck(deck: Deck,  compareFn?: ((a: Card, b: Card) => number) | undefined): Deck {
    deck.cards = deck.cards.sort(compareFn)
    return deck
}

export function SortByType(a: Card, b: Card) {
    return GetCardType(a) - GetCardType(b)
}

export function FilterByMainCards(deck: Deck): Deck {
    let newDeck:Deck = {} as Deck
    newDeck.cards = deck.cards.filter(card => {
        return GetCardType(card) < 100
    })
    return newDeck
}

export function FilterByExtraCards(deck: Deck): Deck {
    let newDeck:Deck = {} as Deck
    newDeck.cards = deck.cards.filter(card => {
        return GetCardType(card) >= 100
    })
    return newDeck
}

export function ToYdkFileString(deck: Deck): string {
    let deckYdkFileString= "#main\n"

    let mainDeck = FilterByMainCards(deck)
    for (let i = 0; i <mainDeck.cards.length; i++) {
        deckYdkFileString += mainDeck.cards[i].id + "\n"
    }

    deckYdkFileString += "\n#extra\n"
    let extraDeck = FilterByExtraCards(deck)
    for (let i = 0; i <extraDeck.cards.length; i++) {
        deckYdkFileString += extraDeck.cards[i].id + "\n"
    }

    return deckYdkFileString
}
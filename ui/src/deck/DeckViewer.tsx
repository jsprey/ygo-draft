import CardViewer from "./CardViewer";
import {
    Card, CardType,
    Deck,
    FilterByExtraCards,
    FilterByMainCards,
    FilterByType,
    SortDeck
} from "../api/CardModel";

export type DeckViewerProps = {
    deck: Deck
}

function DeckViewer(props: DeckViewerProps) {
    let myInt = 50000

    let deck = SortDeck(props.deck)

    let mainDeckCards = FilterByMainCards(deck.cards)
    const mainDeckTrapCardsCount = FilterByType(mainDeckCards, [CardType.TrapCard]).length
    const mainDeckSpellCardsCount = FilterByType(mainDeckCards, [CardType.SpellCard]).length
    const mainDeckMonsterCardsCount = mainDeckCards.length - mainDeckTrapCardsCount - mainDeckSpellCardsCount
    let mainDeckBody = mainDeckCards.map((card: Card) =>
        <span key={myInt++}><CardViewer card={card}/></span>
    );

    let extraDeckCards = FilterByExtraCards(deck.cards)
    let extraDeckTrapCardsCount = FilterByType(extraDeckCards, [CardType.TrapCard]).length
    let extraDeckSpellCardsCount = FilterByType(extraDeckCards, [CardType.SpellCard]).length
    let extraDeckMonsterCardsCount = extraDeckCards.length - extraDeckSpellCardsCount - extraDeckTrapCardsCount
    let extraDeckBody = extraDeckCards.map((card: Card) =>
        <span key={myInt++}><CardViewer card={card}/></span>
    );

    return <>
        <span className={"fw-bold font-monospace text-xl"}>Main Deck</span>
        <div>
            <span className={"mr-2 font-monospace fw-light"}>{mainDeckMonsterCardsCount} Monster Cards |</span>
            <span className={"mr-2 font-monospace fw-light"}>{mainDeckSpellCardsCount} Spell Cards |</span>
            <span className={"mr-2 font-monospace fw-light"}>{mainDeckTrapCardsCount} Trap Cards</span>
        </div>
        <div className={"p-2 grid grid-cols-10 gap-2 bg-dark mt-2 mb-4"}>{mainDeckBody}</div>

        <span className={"fw-bold font-monospace text-xl"}>Extra Deck</span>
        <div>
            <span className={"mr-2 fw-light"}>{extraDeckMonsterCardsCount} Monster Cards |</span>
            <span className={"mr-2 fw-light"}>{extraDeckSpellCardsCount} Spell Cards |</span>
            <span className={"mr-2 fw-light"}>{extraDeckTrapCardsCount} Trap Cards</span>
        </div>
        <div className={"p-2 grid grid-cols-10 gap-2 bg-dark mt-2 mb-2"}>{extraDeckBody}</div>
    </>
}

export default DeckViewer
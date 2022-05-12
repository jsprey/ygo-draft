import CardViewer from "./CardViewer";
import {Card, Deck, FilterByExtraCards, FilterByMainCards, SortByType, SortDeck} from "../api/CardModel";

export type DeckViewerProps = {
    deck: Deck
}

function DeckViewer(props: DeckViewerProps) {
    let myInt = 50000

    let deck = SortDeck(props.deck, SortByType)

    let mainDeck = FilterByMainCards(deck)
    let mainDeckBody = mainDeck.cards.map((card: Card) =>
        <span key={myInt++}><CardViewer card={card}/></span>
    );

    let extraDeck = FilterByExtraCards(deck)
    let extraDeckBody = extraDeck.cards.map((card: Card) =>
        <span key={myInt++}><CardViewer card={card}/></span>
    );

    return <>
        <h3>Main Deck</h3>
        <div className={"p-2 grid grid-cols-10 gap-2 bg-dark mb-4"}>{mainDeckBody}</div>
        <h3>Extra Deck</h3>
        <div className={"p-2 grid grid-cols-10 gap-2 bg-dark mb-2"}>{extraDeckBody}</div>
    </>
}

export default DeckViewer
import CardViewer from "./CardViewer";
import {Card, Deck, FilterByExtraCards, FilterByMainCards, SortAscending, SortDeck} from "../api/CardModel";

export type DeckViewerProps = {
    deck: Deck
}

function DeckViewer(props: DeckViewerProps) {
    let myInt = 50000

    let deck = SortDeck(props.deck)

    let mainDeckCards = FilterByMainCards(deck.cards)
    let mainDeckBody = mainDeckCards.map((card: Card) =>
        <span key={myInt++}><CardViewer card={card}/></span>
    );

    let extraDeckCards = FilterByExtraCards(deck.cards)
    let extraDeckBody = extraDeckCards.map((card: Card) =>
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
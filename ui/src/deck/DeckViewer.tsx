import {
    Deck,
    FilterByExtraCards,
    FilterByMainCards,
    SortDeck
} from "../api/CardModel";
import MultiCardViewer from "./MultiCardViewer";

export type DeckViewerProps = {
    deck: Deck
}

function DeckViewer(props: DeckViewerProps) {
    let deck = SortDeck(props.deck)
    let mainDeckCards = FilterByMainCards(deck.cards)
    let extraDeckCards = FilterByExtraCards(deck.cards)

    return <>
        <MultiCardViewer name={"Main Deck"} showDetails={true} cards={mainDeckCards}/>
        <MultiCardViewer name={"Extra Deck"} showDetails={true} cards={extraDeckCards}/>
    </>
}

export default DeckViewer
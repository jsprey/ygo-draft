import {
    Deck,
    FilterByExtraCards,
    FilterByMainCards,
    SortDeck
} from "../../../api/CardModel";
import MultiCardViewer from "./MultiCardViewer";
import classNames from "classnames";

export type DeckViewerProps = {
    deck: Deck
    className?: string
}

function DeckViewer(props: DeckViewerProps) {
    let deck = SortDeck(props.deck)
    let mainDeckCards = FilterByMainCards(deck.cards)
    let extraDeckCards = FilterByExtraCards(deck.cards)

    const rootCN = classNames(props.className ? props.className : "")
    return <div className={rootCN}>
        <MultiCardViewer name={"Main Deck"} showDetails={true} cards={mainDeckCards}/>
        <div style={{height: "1rem"}}/>
        <MultiCardViewer name={"Extra Deck"} showDetails={true} cards={extraDeckCards}/>
    </div>
}

export default DeckViewer
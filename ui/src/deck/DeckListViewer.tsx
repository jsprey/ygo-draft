import {
    FilterByExtraCards,
    FilterByMainCards,
    SortDeck
} from "../api/CardModel";
import MultiCardViewer from "./MultiCardViewer";
import {useCardsBulk} from "../api/hooks/cards/useCardsBulk";
import React from "react";
import Spinner from "../core/Spinner";
import Alert from "../core/Alert";
import classNames from "classnames";

export type DeckListViewerProps = {
    deckList: string[]
    className?: string
}

function DeckListViewer(props: DeckListViewerProps) {
    const {isLoading, data, error} = useCardsBulk(props.deckList)

    if (isLoading) {
        return <Spinner/>
    }

    if (error || !data) {
        return <Alert variant={'danger'}>Failed to load deck!</Alert>
    }

    let deck = SortDeck(data)
    let mainDeckCards = FilterByMainCards(deck.cards)
    let extraDeckCards = FilterByExtraCards(deck.cards)

    const rootCN = classNames(props.className ? props.className : "")
    return <div className={rootCN}>
        <MultiCardViewer name={"Main Deck"} showDetails={true} cards={mainDeckCards}/>
        <MultiCardViewer name={"Extra Deck"} showDetails={true} cards={extraDeckCards}/>
    </div>
}

export default DeckListViewer
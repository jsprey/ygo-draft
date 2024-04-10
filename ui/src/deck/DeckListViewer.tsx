import {
    Deck,
    FilterByExtraCards,
    FilterByMainCards,
    SortDeck
} from "../api/CardModel";
import MultiCardViewer from "./MultiCardViewer";
import {useCardsBulk} from "../api/hooks/cards/useCardsBulk";
import {Alert, Spinner} from "react-bootstrap";
import React from "react";

export type DeckListViewerProps = {
    deckList: string[]
}

function DeckListViewer(props: DeckListViewerProps) {
    const {isLoading, data, error} = useCardsBulk(props.deckList)

    if (isLoading) {
        return <Spinner animation={"border"}></Spinner>
    }

    if (error || !data) {
        return <Alert variant={'danger'}>Failed to load deck!</Alert>
    }

    let deck = SortDeck(data)
    let mainDeckCards = FilterByMainCards(deck.cards)
    let extraDeckCards = FilterByExtraCards(deck.cards)

    return <>
        <MultiCardViewer name={"Main Deck"} showDetails={true} cards={mainDeckCards}/>
        <MultiCardViewer name={"Extra Deck"} showDetails={true} cards={extraDeckCards}/>
    </>
}

export default DeckListViewer
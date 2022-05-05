import {Card, useRandomCard} from "./hooks/useRandomCard";
import CardViewer from "./CardViewer";
import {Deck, useRandomCards} from "./hooks/useRandomCards";
import DeckViewer from "./DeckViewer";
import React from "react";

function RandomDeckGenerator() {
    const {isLoading, error, data} = useRandomCards("mydeck", 40)

    let body;
    if (isLoading) {
        body = <p>Loading...</p>
    } else if (!data || error) {
        body = <p>Fehler beim Laden des Decks!</p>
    } else if (data) {
        body = <DeckViewer deck={data}/>
    } else {
        body = <p>no data</p>
    }

    return <>
        {body}
    </>
}

export default RandomDeckGenerator

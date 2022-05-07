import {Deck, useRandomCards} from "./hooks/useRandomCards";
import DeckViewer from "./DeckViewer";
import React, {useState} from "react";

const emptyDeck: Deck = {cards: []}

function DeckGeneratorPage() {
    const [myDeck, setDeck] = useState(emptyDeck)
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
        <h1 className={"mt-3"}> Deck Generation </h1>
        {/*{body}*/}
    </>
}

export default DeckGeneratorPage

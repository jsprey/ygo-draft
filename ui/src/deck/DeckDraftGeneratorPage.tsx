import React, {Dispatch, SetStateAction, useState} from "react";
import {Deck, ToYdkFileString} from "../api/CardModel";
import {useRandomCards} from "../api/hooks/useCards";
import DeckViewer from "./DeckViewer";
import {Alert, Button, Spinner} from "react-bootstrap";
import {YgoQueryClient} from "../index";
import MultiCardViewer from "./MultiCardViewer";
import {ExportDeck} from "./DeckRandomGeneratorPage";

const emptyDeck: Deck = {cards: []}

function DeckDraftGeneratorPage() {
    const [myDeck, setDeck] = useState(emptyDeck)
    const [draftDeck, setDraftDeck] = useState(emptyDeck)
    const [draftSize, setDraftSize] = useState(5)

    const {data, isLoading, error} = useRandomCards("draft_generator", draftSize, {enabled: true, staleTime: Infinity})

    let body;
    if (draftDeck == emptyDeck) {
        if (isLoading) {
            body = <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading Deck...</span>
            </Spinner>
        } else if (error) {
            body = <Alert variant={"danger"}>
                Could not load deck!
            </Alert>
        } else if (data) {
            setDraftDeck(data)
        }
    }

    return <>
        <h1 className={"mt-3 mb-3"}>
            Deck Generation
            <Button className={"ml-4 object-center"}
                    variant="primary"
                    disabled={isLoading}
                    onClick={() => !isLoading ? ExportDeck(myDeck) : null}>
                Export
            </Button>
        </h1>
        {body}
        <MultiCardViewer name={"Draft Area"} showDetails={false} cards={draftDeck.cards} />
        <p className={"mt-5 text-3xl"}>Generated Deck</p>
        <DeckViewer deck={myDeck}/>
    </>
}

export default DeckDraftGeneratorPage

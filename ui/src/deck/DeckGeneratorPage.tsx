import React, {Dispatch, SetStateAction, useState} from "react";
import {Deck, ToYdkFileString} from "../api/CardModel";
import {useRandomCards} from "../api/hooks/useCards";
import DeckViewer from "./DeckViewer";
import {Alert, Button, Spinner} from "react-bootstrap";
import {YgoQueryClient} from "../index";

const emptyDeck: Deck = {cards: []}

function DeckGeneratorPage() {
    const {data, isLoading, error} = useRandomCards("deck_generator", 40, {enabled: true, staleTime: Infinity})
    const [myDeck, setDeck] = useState(emptyDeck)

    let body;
    if (myDeck == emptyDeck) {
        if (isLoading) {
            body = <Spinner animation="border" role="status">
                <span className="visually-hidden">Loading Deck...</span>
            </Spinner>
        } else if (error) {
            body = <Alert variant={"danger"}>
                Could not load deck!
            </Alert>
        } else if (data) {
            setDeck(data)
        }
    } else {
        body = <>
            <DeckViewer deck={myDeck}/>
        </>
    }

    return <>
        <h1 className={"mt-3 mb-3"}>
            Deck Generation
            <Button className={"ml-4 object-center"}
                    variant="primary"
                    disabled={isLoading}
                    onClick={() => !isLoading ? resetDeck(setDeck) : null}>
                Recreate
            </Button>
            <Button className={"ml-4 object-center"}
                    variant="primary"
                    disabled={isLoading}
                    onClick={() => !isLoading ? exportDeck(myDeck) : null}>
                Export
            </Button>
        </h1>
        {body}
    </>
}

function resetDeck(setDeck: Dispatch<SetStateAction<Deck>>) {
    setDeck(emptyDeck)

    YgoQueryClient.removeQueries(["random", "deck_generator"])
}

function exportDeck(myDeck: Deck) {
    download("mydeck.ydk", ToYdkFileString(myDeck))
}

function download(filename:string, text:string) {
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

export default DeckGeneratorPage

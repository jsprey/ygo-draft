import React, {Dispatch, SetStateAction, useState} from "react";
import {Deck, ToYdkFileString} from "../../../api/CardModel";
import {useRandomCards} from "../../../api/hooks/cards/useCards";
import DeckViewer from "./DeckViewer";
import {YgoQueryClient} from "../../../index";
import {CardFilter} from "../../../api/CardFilter";
import Spinner from "../../../core/Spinner";
import Alert from "../../../core/Alert";
import Button from "../../../core/Button";
import {ShowSuccessfulSnack} from "../../../core/Snacks";

const emptyDeck: Deck = {cards: []}

function DeckRandomGeneratorPage() {
    const {data, isLoading, error} = useRandomCards("deck_generator", 40, {} as CardFilter, {
        enabled: true,
        staleTime: Infinity
    })
    const [myDeck, setDeck] = useState(emptyDeck)

    let body;
    if (myDeck === emptyDeck) {
        if (isLoading) {
            body = <Spinner/>
        } else if (error) {
            body = <Alert variant={'danger'}>
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

    return <div>
        {body}
        <h1 className={"dark:text-light mt-2 flex justify-end"}>
            <Button className={"object-center"}
                    variant={"primary"}
                    disabled={isLoading}
                    onClick={() => !isLoading ? resetDeck(setDeck) : null}>
                Recreate
            </Button>
            <Button className={"ml-2 object-center"}
                    variant={"primary"}
                    disabled={isLoading}
                    onClick={() => !isLoading ? ExportDeck(myDeck) : null}>
                Export
            </Button>
        </h1>
    </div>
}

function resetDeck(setDeck: Dispatch<SetStateAction<Deck>>) {
    setDeck(emptyDeck)

    YgoQueryClient.removeQueries(["random", "deck_generator"])
}

export function ExportDeck(myDeck: Deck) {
    downloadDeck("mydeck.ydk", ToYdkFileString(myDeck))
    ShowSuccessfulSnack("Deck is begin downloaded. Check the downloads of your browser.")
}

function downloadDeck(filename: string, text: string) {
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

export default DeckRandomGeneratorPage

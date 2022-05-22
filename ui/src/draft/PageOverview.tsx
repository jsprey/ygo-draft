import React from "react";
import {Deck} from "../api/CardModel";
import DeckViewer from "../deck/DeckViewer";
import {ExportDeck} from "../deck/DeckRandomGeneratorPage";
import {Button} from "react-bootstrap";

export type PageOverviewProps = {
    deck: Deck
}

function PageOverview(props: PageOverviewProps) {
    return <>
        <DeckViewer deck={props.deck}/>
        <div className={"flex place-content-end"}>
            <Button className={"ml-4 object-center"}
                    variant="primary"
                    onClick={() => ExportDeck(props.deck)}>
                Export
            </Button>
        </div>
    </>
}


export default PageOverview

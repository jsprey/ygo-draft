import React from "react";
import {Deck} from "../../api/CardModel";
import DeckViewer from "../shared/deck/DeckViewer";
import Button from "../../core/Button";
import {ExportDeck} from "../shared/deck/DeckRandomGeneratorPage";

export type LocaleDraftWizardResultPageProps = {
    deck: Deck
    onAbort: () => void
}

function LocaleDraftWizardResultPage(props: LocaleDraftWizardResultPageProps) {
    return <>
        <DeckViewer className={"mt-2"} deck={props.deck}/>
        <div className={"flex place-content-end"}>
            <Button variant={"danger"}
                    className={"mt-2 mr-2"}
                    onClick={() => props.onAbort()}>
                Reset
            </Button>
            <Button variant={"primary"}
                    className={"mt-2"}
                    onClick={() => ExportDeck(props.deck)}>
                Export
            </Button>
        </div>
    </>
}

export default LocaleDraftWizardResultPage

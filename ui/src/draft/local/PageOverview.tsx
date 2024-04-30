import React from "react";
import {Deck} from "../../api/CardModel";
import DeckViewer from "../../deck/DeckViewer";
import Button from "../../core/Button";

export type PageOverviewProps = {
    deck: Deck
    onSubmit: (deck:Deck) => void
    submitName: string
}

function PageOverview(props: PageOverviewProps) {
    return <>
        <DeckViewer className={"mt-2"} deck={props.deck}/>
        <div className={"flex place-content-end"}>
            <Button variant={"primary"}
                    className={"mt-2"}
                    onClick={() => props.onSubmit(props.deck)}>
                {props.submitName}
            </Button>
        </div>
    </>
}


export default PageOverview

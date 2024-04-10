import React from "react";
import {Deck} from "../../api/CardModel";
import DeckViewer from "../../deck/DeckViewer";
import {Button} from "react-bootstrap";

export type PageOverviewProps = {
    deck: Deck
    onSubmit: (deck:Deck) => void
    submitName: string
}

function PageOverview(props: PageOverviewProps) {
    return <>
        <DeckViewer deck={props.deck}/>
        <div className={"flex place-content-end"}>
            <Button className={"ml-4 object-center"}
                    variant="primary"
                    onClick={() => props.onSubmit(props.deck)}>
                {props.submitName}
            </Button>
        </div>
    </>
}


export default PageOverview

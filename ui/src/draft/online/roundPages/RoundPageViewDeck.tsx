import React from "react";
import DeckListViewer from "../../shared/deck/DeckListViewer";
import Button from "../../../core/Button";
import {useNavigate} from "react-router";
import {Draft} from "../../../api/Draft";
import {DraftPagePath} from "../../../routes/AppRouter";

type RoundPageViewDeckProps = {
    deckList: string[]
    draft: Draft
};

function RoundPageViewDeck(props: RoundPageViewDeckProps) {
    const navigate = useNavigate()

    return <div className={"flex flex-col"}>
        <DeckListViewer className={"flex flex-col gap-2"} deckList={props.deckList}/>
        <Button variant={"primary"} className={"mt-2 self-end"}
                onClick={() => navigate(DraftPagePath(props.draft.id))}>Back</Button>
    </div>
}

export default RoundPageViewDeck

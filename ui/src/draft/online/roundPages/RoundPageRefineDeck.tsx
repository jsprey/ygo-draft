import React from "react";
import Button from "../../../core/Button";
import {useNavigate} from "react-router";
import {Draft} from "../../../api/Draft";
import {DraftPagePath} from "../../../routes/AppRouter";

type RoundPageRefineDeckProps = {
    draft: Draft
};

function RoundPageRefineDeck(props: RoundPageRefineDeckProps) {
    const navigate = useNavigate()

    return <div className={"flex flex-col"}>
        <span>Lets refine!</span>
        <Button variant={"primary"} className={"mt-2 self-end"} onClick={() => navigate(DraftPagePath(props.draft.id))}>Back</Button>
    </div>
}

export default RoundPageRefineDeck

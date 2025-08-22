import React from "react";
import {Draft} from "../../../api/Draft";
import {useNavigate} from "react-router";
import {DraftPagePath} from "../../../routes/AppRouter";
import OnlineDeckDraftWizard from "../OnlineDeckDraftWizard";

type RoundPageDraftInitialDeckProps = {
    roundID: string
    draft: Draft
};

function RoundPageDraftInitialDeck(props: RoundPageDraftInitialDeckProps) {
    const navigate = useNavigate()

    function onAbortDraft() {
        navigate(DraftPagePath(props.draft.id))
    }

    return <>
        <OnlineDeckDraftWizard onAbort={onAbortDraft}
                               draftID={""+props.draft.id}
                               roundID={props.roundID}
                               settings={props.draft.settings}/>
    </>
}

export default RoundPageDraftInitialDeck

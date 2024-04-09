import React from "react";
import {useDraftRoundDecks} from "../../api/hooks/drafts/useDraftRoundDecks";
import {useParams} from "react-router";
import {Alert, Spinner} from "react-bootstrap";
import {useDraft} from "../../api/hooks/drafts/useDraft";
import {usePrompt} from "../../api/hooks/usePromptBlocker";

type DraftMyDeckPageParams = {
    id: string
    roundID: string
};

function DraftMyDeckPage() {
    usePrompt("Your unfinished deck is going to be deleted when leaving the page. Are you sure you want to leave?", true);

    let params = useParams<DraftMyDeckPageParams>() as DraftMyDeckPageParams;
    const draftQuery = useDraft(params.id)
    const roundDecksQuery = useDraftRoundDecks(params.roundID)

    if (roundDecksQuery.isLoading || draftQuery.isLoading) {
        return <div className={"p-2 flex justify-content-center align-items-center"}>
            <Spinner animation={"border"} />
        </div>
    }

    if (roundDecksQuery.error) {
        return <div className={"pt-2 flex justify-content-center align-items-center"}>
            <Alert variant={"danger"}>Failed to load the draft round deck!</Alert>
        </div>
    }

    if (draftQuery.error) {
        return <div className={"pt-2 flex justify-content-center align-items-center"}>
            <Alert variant={"danger"}>Failed to load the draft settings!</Alert>
        </div>
    }

    return <div>
        {JSON.stringify(draftQuery.data?.settings)}
    </div>
}

export default DraftMyDeckPage

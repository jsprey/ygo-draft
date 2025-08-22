import React from "react";
import DeckViewer from "../shared/deck/DeckViewer";
import Button from "../../core/Button";
import {Deck, ToStringList} from "../../api/CardModel";
import {SubmitDeckRequest, useSubmitDeck} from "../../api/hooks/drafts/useSubmitDeck";
import {ShowErrorSnack, ShowSuccessfulSnack} from "../../core/Snacks";
import {DraftPagePath} from "../../routes/AppRouter";
import {useNavigate} from "react-router";
import Spinner from "../../core/Spinner";
import {removeStorageKeys} from "../../api/hooks/storage/useLocalState";
import {OnlineDeckDraftWizardStoragePrefix} from "./OnlineDeckDraftWizard";

type OnlineDeckDraftWizardResultPageProps = {
    deck: Deck
    draftID: string
    roundID: string
    onAbort: () => void
};

function OnlineDeckDraftWizardResultPage(props: OnlineDeckDraftWizardResultPageProps) {
    var navigate = useNavigate();

    const submitDeckMutation = useSubmitDeck({
        onSuccess: () => {
            ShowSuccessfulSnack("Deck submitted")
            navigate(DraftPagePath(props.draftID))
            removeStorageKeys(OnlineDeckDraftWizardStoragePrefix)
        },
        onError: () => ShowErrorSnack("Failed to submit deck.")
    })

    function onSubmitDeck(deck: Deck) {
        const request: SubmitDeckRequest = {
            roundID: parseInt(props.roundID),
            deck: ToStringList(deck)
        }
        submitDeckMutation.mutate(request)
    }

    return <div>
        <DeckViewer className={"mt-2"} deck={props.deck}/>
        <div className={"flex place-content-end"}>
            <Button variant={"primary"}
                    className={"mt-2"}
                    onClick={() => onSubmitDeck(props.deck)}
                    disabled={submitDeckMutation.isLoading}>
                {submitDeckMutation.isLoading ? <Spinner/> : null} Submit
            </Button>
        </div>
    </div>
}

export default OnlineDeckDraftWizardResultPage

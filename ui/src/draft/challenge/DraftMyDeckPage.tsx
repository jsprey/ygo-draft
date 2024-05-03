import React, {useState} from "react";
import {useDraftRoundDecks, useDraftRoundDecksPayload} from "../../api/hooks/drafts/useDraftRoundDecks";
import {useNavigate, useParams} from "react-router";
import {useDraft} from "../../api/hooks/drafts/useDraft";
import {usePrompt} from "../../api/hooks/usePromptBlocker";
import OnlineDeckDraftWizard from "./OnlineDeckDraftWizard";
import {Deck, ToStringList} from "../../api/CardModel";
import {SubmitDeckRequest, useSubmitDeck} from "../../api/hooks/drafts/useSubmitDeck";
import {enqueueSnackbar} from "notistack";
import DeckListViewer from "../../deck/DeckListViewer";
import Spinner from "../../core/Spinner";
import Alert from "../../core/Alert";
import Button from "../../core/Button";

type DraftMyDeckPageParams = {
    id: string
    roundID: string
};

function DraftMyDeckPage() {
    const [prompt, setPrompt] = useState<boolean>(true)
    usePrompt("Your unfinished deck is going to be deleted when leaving the page. Are you sure you want to leave?", prompt);

    const params = useParams<DraftMyDeckPageParams>() as DraftMyDeckPageParams;
    const navigate = useNavigate()
    const draftQuery = useDraft(params.id)
    const roundDecksQuery = useDraftRoundDecks(params.roundID, {
        onSuccess: (data: useDraftRoundDecksPayload) => {
            if (data) {
                setPrompt(false)
            }
        },
    })

    const submitDeckMutation = useSubmitDeck({
        onSuccess: () => {
            ShowSuccessfulSnack("Deck submitted")
            setPrompt(false)
            navigate(`/draft/${draftQuery.data?.id}`)
        },
        onError: () => ShowErrorSnack("Failed to submit deck.")
    })

    if (roundDecksQuery.isLoading || draftQuery.isLoading) {
        return <div className={"p-2 flex justify-center items-center"}>
            <Spinner/>
        </div>
    }

    if (roundDecksQuery.error || !roundDecksQuery.data) {
        return <div className={"pt-2 flex justify-center items-center"}>
            <Alert variant={'danger'}>Failed to load the draft round deck!</Alert>
        </div>
    }

    if (draftQuery.error || !draftQuery.data) {
        return <div className={"pt-2 flex justify-center items-center"}>
            <Alert variant={'danger'}>Failed to load the draft settings!</Alert>
        </div>
    }

    function onSubmitDeck(deck: Deck) {
        const request: SubmitDeckRequest = {
            roundID: parseInt(params.roundID),
            deck: ToStringList(deck)
        }
        submitDeckMutation.mutate(request)
    }

    function onAbortDraft() {
        if (draftQuery.data) {
            setPrompt(false)
            navigate(`/draft/${draftQuery.data.id}`)
        }
    }

    if (draftQuery.data && roundDecksQuery.data && roundDecksQuery.data.user_deck.length > 0) {
        // inspect deck show deck
        return <div className={"flex flex-col"}>
            <DeckListViewer className={"flex flex-col gap-2"} deckList={roundDecksQuery.data.user_deck}/>
            <Button variant={"primary"} className={"mt-2 self-end"} onClick={() => navigate(`/draft/${params.id}`)}>Back</Button>
        </div>
    } else if (roundDecksQuery.data && draftQuery.data.current_round_number === 1) {
        // create initial deck
        return <div>
            <OnlineDeckDraftWizard submitName={"Finish"}
                                   onSubmit={onSubmitDeck}
                                   onAbort={onAbortDraft}
                                   settings={draftQuery.data.settings}/>
        </div>
    } else if (roundDecksQuery.data && draftQuery.data.current_round_number >= 1) {
        // refine existing deck
        return <div className={"flex flex-col"}>
            <span>Lets refine!</span>
            <Button variant={"primary"} className={"mt-2 self-end"} onClick={() => navigate(`/draft/${params.id}`)}>Back</Button>
        </div>
    }

    return <span>This should not happen</span>
}

export function ShowSuccessfulSnack(message: string) {
    enqueueSnackbar(message, {
        autoHideDuration: 6000,
        variant: "success"
    })
}

export function ShowErrorSnack(message: string) {
    enqueueSnackbar(message, {
        autoHideDuration: 6000,
        variant: "error"
    })
}

export default DraftMyDeckPage

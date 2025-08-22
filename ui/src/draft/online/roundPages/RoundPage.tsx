import React from "react";
import {useDraftRoundDecks} from "../../../api/hooks/drafts/useDraftRoundDecks";
import {useParams} from "react-router";
import {useDraft} from "../../../api/hooks/drafts/useDraft";
import Spinner from "../../../core/Spinner";
import Alert from "../../../core/Alert";
import RoundPageRefineDeck from "./RoundPageRefineDeck";
import RoundPageDraftInitialDeck from "./RoundPageDraftInitialDeck";
import RoundPageViewDeck from "./RoundPageViewDeck";

type RoundPageParams = {
    id: string
    roundID: string
};

function RoundPage() {
    const params = useParams<RoundPageParams>() as RoundPageParams;
    const draftQuery = useDraft(params.id)
    const roundDecksQuery = useDraftRoundDecks(params.roundID)

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

    if (draftQuery.data && roundDecksQuery.data && roundDecksQuery.data.user_deck.length > 0) {
        return <RoundPageViewDeck deckList={roundDecksQuery.data.user_deck}
                                  draft={draftQuery.data}/>
    } else if (roundDecksQuery.data && draftQuery.data.current_round_number === 1) {
        return <RoundPageDraftInitialDeck draft={draftQuery.data}
                                          roundID={params.roundID}/>
    } else if (roundDecksQuery.data && draftQuery.data.current_round_number >= 1) {
        return <RoundPageRefineDeck draft={draftQuery.data}/>
    }

    return <span>This should not happen</span>
}

export default RoundPage

import React from "react";
import {Draft} from "../../api/Draft";
import {useDraftRounds} from "../../api/hooks/drafts/useDraftRounds";
import {Alert} from "react-bootstrap";
import DraftRoundListEntry from "./DraftRoundListEntry";

type DraftRoundListProps = {
    draft: Draft
};

function DraftRoundList(props: DraftRoundListProps) {
    const draftRoundsRequest = useDraftRounds(props.draft.id)

    let content = <></>
    if (draftRoundsRequest.isLoading) {
        content = <p className={"placeholder placeholder-glow vw-100 vh-100"}></p>
    } else if (draftRoundsRequest.isError) {
        content = <Alert variant={"danger"}>Failed to load draft rounds!</Alert>
    } else if (draftRoundsRequest.data) {
        const rounds = draftRoundsRequest.data.rounds

        const roundsElements = rounds.map(round => {
            return <div key={`draft_overview_round_${round.id}`}>
                <DraftRoundListEntry draft={props.draft} round={round}/>
            </div>
        })

        content = <div className={"dark:text-white mt-2"}>
            {roundsElements}
        </div>
    }

    return content
}

export default DraftRoundList

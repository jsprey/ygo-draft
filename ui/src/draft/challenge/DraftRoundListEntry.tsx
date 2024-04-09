import React from "react";
import {DraftRound} from "../../api/DraftRound";
import classNames from "classnames";
import {useDraftRoundDecks} from "../../api/hooks/drafts/useDraftRoundDecks";
import {Alert, Spinner} from "react-bootstrap";
import {useNavigate} from "react-router";

type DraftRoundListEntryProps = {
    round: DraftRound
};

function DraftRoundListEntry(props: DraftRoundListEntryProps) {
    const round = props.round
    const navigate = useNavigate();
    const roundDecksQuery = useDraftRoundDecks(""+round.id)

    function createPlayerElement() {
        const playerCN = classNames("m-2")

        if (roundDecksQuery.isLoading) {
            return <Spinner animation={"border"}/>
        }

        if (roundDecksQuery.error || !roundDecksQuery.data) {
            return <Alert variant={"danger"}>Failed to load deck!</Alert>
        }

        let isDeckEmpty = roundDecksQuery.data.user_deck.length === 0
        return <div className={playerCN}>
            <span className={"btn btn-primary"} onClick={() => navigate(`/draft/${round.draft_id}/${round.id}`)}>
                {isDeckEmpty ? "Draft Deck" : "Inspect Deck"}
            </span>
        </div>
    }

    function createEnemyElement() {
        const cnList = classNames("m-2")

        if (roundDecksQuery.isLoading) {
            return <Spinner animation={"border"}/>
        }

        if (roundDecksQuery.error || !roundDecksQuery.data) {
            return <Alert variant={"danger"}>Failed to load deck!</Alert>
        }

        const isDeckEmpty = roundDecksQuery.data.enemy_deck.length === 0
        return <div className={cnList}>
            <span>
                {isDeckEmpty ? "Waiting for Draft" : "Ready"}
            </span>
        </div>
    }

    const containerCN = classNames("flex justify-content-between align-items-center", "w-100 p-0", "rounded", "bg-gray-200 dark:bg-gray-600", "border")
    return <div className={containerCN} key={`draft_overview_round_${round.id}`}>
        {createPlayerElement()}
        <div
            className={classNames("flex flex-col align-items-center", "p-2", "bg-green-200 dark:bg-green-950", "border-start border-end")}>
            <span className={"fw-bold text-xl"}>Round {round.round_number}</span>
            <span>{getStatusDisplayMessage(round)}</span>
        </div>
        {createEnemyElement()}
    </div>
}

function getStatusDisplayMessage(round: DraftRound) {
    if (round.status === "preparation") {
        return round.round_number === 1 ? "Deck Draft Phase" : "Deck Refinement Phase"
    }

    if (round.status === "fighting") {
        return "Waiting for Results"
    }

    if (round.status === "finished") {
        return "Finished"
    }
}

export default DraftRoundListEntry

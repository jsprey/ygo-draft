import React from "react";
import {DraftRound} from "../../api/DraftRound";
import classNames from "classnames";
import {useDraftRoundDecks} from "../../api/hooks/drafts/useDraftRoundDecks";
import {Alert, Spinner} from "react-bootstrap";
import {useNavigate} from "react-router";
import SvgIconButton from "../../core/SvgIconButton";
import {SubmitWinnerRequest, useSubmitWinner} from "../../api/hooks/drafts/useSubmitWinner";
import {ShowSuccessfulSnack} from "./DraftMyDeckPage";
import {Draft} from "../../api/Draft";
import {useCurrentUser} from "../../api/hooks/users/useUser";
import {useQueryClient} from "react-query";

const IconEye = <SvgIconButton size={25} classNames={"fill-white mr-1"}>
    <path d="M10.5 8a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z"/>
    <path
        d="M0 8s3-5.5 8-5.5S16 8 16 8s-3 5.5-8 5.5S0 8 0 8zm8 3.5a3.5 3.5 0 1 0 0-7 3.5 3.5 0 0 0 0 7z"/>
</SvgIconButton>

type DraftRoundListEntryProps = {
    round: DraftRound
    draft: Draft
};

function DraftRoundListEntry(props: DraftRoundListEntryProps) {
    const round = props.round
    const navigate = useNavigate();
    const currentUser = useCurrentUser()
    const queryClient = useQueryClient()
    const winnerMutation = useSubmitWinner({
        onSuccess: () => {
            ShowSuccessfulSnack("You won, congratulation!")
            queryClient.refetchQueries({queryKey: ["draft", round.draft_id, "rounds"]})
            queryClient.refetchQueries({queryKey: ["draft", round.draft_id]})
        },
        onError: () => ShowSuccessfulSnack("Keep up, you will get your revenge!")
    })
    const roundDecksQuery = useDraftRoundDecks("" + round.id)

    function createPlayerElement() {
        if (roundDecksQuery.isLoading || currentUser.isLoading) {
            return <Spinner animation={"border"}/>
        }

        if (roundDecksQuery.error || !roundDecksQuery.data || currentUser.error || !currentUser.data) {
            return <Alert variant={"danger"}>Failed to load data!</Alert>
        }

        const playerCN = classNames("m-2")
        let isDeckEmpty = roundDecksQuery.data.user_deck.length === 0

        let playerActions: JSX.Element
        if (isDeckEmpty && props.draft.status === "surrender") {
            playerActions = <></>
        } else if (isDeckEmpty) {
            playerActions =
                <span className={"btn btn-primary"} onClick={() => navigate(`/draft/${round.draft_id}/${round.id}`)}>
                Draft Deck
            </span>
        } else {
            playerActions =
                <span className={"btn btn-primary"} onClick={() => navigate(`/draft/${round.draft_id}/${round.id}`)}>
                <div className={"flex justiy-content-between"}>
                {IconEye}Inspect Deck
            </div>
            </span>
        }

        return <div className={playerCN}>
            {playerActions}
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

        let content = <></>
        if (round.status == "preparation" && !isDeckEmpty) {
            content = <div className={"text-green-700 dark:text-green-300 text-uppercase text-2xl"}>Ready</div>
        } else if (round.status == "preparation" && isDeckEmpty) {
            content = <div className={"text-yellow-700 dark:text-yellow-300 text-uppercase text-2xl"}>Drafting</div>
        } else if (round.status == "fighting") {
            content = <>In Duel</>
        } else if (round.status == "finished") {
            content = <></>
        }

        return <div className={cnList}>
            <span>
                {content}
            </span>
        </div>
    }

    function getLeftPlayerStatus() {
        if (round.status === "fighting") {
            return <div className={"mr-2 flex align-items-center justify-content-end"}>
                <button className={"btn btn-success"}
                        onClick={onUserWinButton}>
                    I won
                </button>
            </div>
        } else if (round.status === "finished") {
            if (round.winner_user_id == currentUser.data?.id) {
                return <div
                    className={"mr-2 flex text-uppercase text-xl fw-bold text-green-800 dark:text-green-400 align-items-center justify-content-end"}>
                    Win
                </div>
            } else {
                return <div
                    className={"mr-2 flex text-uppercase text-xl fw-bold text-red-800 dark:text-red-400 align-items-center justify-content-end"}>
                    Lose
                </div>
            }
        } else {
            return <div></div>
        }
    }

    function getRightPlayerStatus() {
        if (round.status === "fighting") {
            return <div className={"ml-2 flex align-items-center justify-content-start"}>
                <button className={"btn btn-danger"}
                        onClick={onEnemyWinButton}>
                    Enemy won
                </button>
            </div>
        } else if (round.status === "finished") {
            if (round.winner_user_id != currentUser.data?.id) {
                return <div
                    className={"ml-2 flex text-uppercase text-xl fw-bold text-green-800 dark:text-green-400 align-items-center justify-content-start"}>
                    Win
                </div>
            } else {
                return <div
                    className={"ml-2 flex text-uppercase text-xl fw-bold text-red-800 dark:text-red-400 align-items-center justify-content-start"}>
                    Lose
                </div>
            }
        } else {
            return <div></div>
        }
    }

    function onUserWinButton() {
        if (!currentUser.data) {
            return
        }

        const request: SubmitWinnerRequest = {
            roundID: round.id,
            winner: currentUser.data.id
        }

        winnerMutation.mutate(request)
    }

    function onEnemyWinButton() {
        const request: SubmitWinnerRequest = {
            roundID: round.id,
            winner: props.draft.challenger_id == currentUser.data?.id ? props.draft.receiver_id : props.draft.challenger_id
        }

        winnerMutation.mutate(request)
    }

    const containerCN = classNames("grid-cols-5 grid", "w-100 p-0", "rounded", "bg-gray-200 dark:bg-gray-600", "border")
    return <div className={containerCN}>
        <div className={"flex align-items-center"}>{createPlayerElement()}</div>
        {getLeftPlayerStatus()}
        <div
            className={classNames("flex flex-col align-items-center", "p-2", "bg-green-200 dark:bg-green-950", "border-start border-end")}>
            <span className={"fw-bold text-xl"}>Round {round.round_number}</span>
            <span>{getStatusDisplayMessage(round)}</span>
        </div>
        {getRightPlayerStatus()}
        <div className={"flex align-items-center justify-content-end"}>{createEnemyElement()}</div>
    </div>
}

function getStatusDisplayMessage(round: DraftRound) {
    if (round.status === "preparation") {
        return round.round_number === 1 ? "Deck Draft Phase" : "Deck Refinement Phase"
    }

    if (round.status === "fighting") {
        return "Dueling"
    }

    if (round.status === "finished") {
        return "Finished"
    }
}

export default DraftRoundListEntry

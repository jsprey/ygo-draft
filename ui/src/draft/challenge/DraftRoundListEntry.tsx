import React from "react";
import {DraftRound} from "../../api/DraftRound";
import classNames from "classnames";
import {useDraftRoundDecks} from "../../api/hooks/drafts/useDraftRoundDecks";
import {useNavigate} from "react-router";
import SvgIconButton from "../../core/SvgIconButton";
import {SubmitWinnerRequest, useSubmitWinner} from "../../api/hooks/drafts/useSubmitWinner";
import {ShowSuccessfulSnack} from "./DraftMyDeckPage";
import {Draft} from "../../api/Draft";
import {useCurrentUser} from "../../api/hooks/users/useUser";
import {useQueryClient} from "react-query";
import Spinner from "../../core/Spinner";
import Alert from "../../core/Alert";
import Button from "../../core/Button";

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
            queryClient.refetchQueries({queryKey: ["draft", round.draft_id, "rounds"]})
            queryClient.refetchQueries({queryKey: ["draft", round.draft_id]})
        },
        onError: () => ShowSuccessfulSnack("Something wrong :o")
    })
    const roundDecksQuery = useDraftRoundDecks("" + round.id)

    function createPlayerElement() {
        if (roundDecksQuery.isLoading || currentUser.isLoading) {
            return <Spinner/>
        }

        if (roundDecksQuery.error || !roundDecksQuery.data || currentUser.error || !currentUser.data) {
            return <Alert variant={'danger'}>Failed to load data!</Alert>
        }

        const playerCN = classNames("m-2")
        let isDeckEmpty = roundDecksQuery.data.user_deck.length === 0

        let playerActions: JSX.Element
        if (isDeckEmpty && props.draft.status === "surrender") {
            playerActions = <></>
        } else if (isDeckEmpty) {
            playerActions =
                <Button variant={"primary"} className={""} onClick={() => navigate(`/draft/${round.draft_id}/${round.id}`)}>
                Draft Deck
            </Button>
        } else {
            playerActions =
                <Button variant={"primary"} className={""} onClick={() => navigate(`/draft/${round.draft_id}/${round.id}`)}>
                <div className={"flex justify-between"}>
                {IconEye}Inspect Deck
            </div>
            </Button>
        }

        return <div className={playerCN}>
            {playerActions}
        </div>
    }

    function createEnemyElement() {
        const cnList = classNames("m-2")

        if (roundDecksQuery.isLoading) {
            return <Spinner/>
        }

        if (roundDecksQuery.error || !roundDecksQuery.data) {
            return <Alert variant={'danger'}>Failed to load deck!</Alert>
        }

        const isDeckEmpty = roundDecksQuery.data.enemy_deck.length === 0

        let content = <></>
        if (round.status === "preparation" && !isDeckEmpty) {
            content = <div className={"text-secondary uppercase text-2xl"}>Ready</div>
        } else if (round.status === "preparation" && isDeckEmpty) {
            content = <div className={"text-secondary uppercase text-2xl"}>Drafting</div>
        } else if (round.status === "fighting") {
            content = <>In Duel</>
        } else if (round.status === "finished") {
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
            return <div className={"mr-2 flex items-center justify-end"}>
                <Button variant={"success"}
                        onClick={onUserWinButton}>
                    I won
                </Button>
            </div>
        } else if (round.status === "finished") {
            if (round.winner_user_id === currentUser.data?.id) {
                return <div
                    className={"mr-2 flex uppercase text-xl font-bold text-success-dark items-center justify-end"}>
                    Win
                </div>
            } else {
                return <div
                    className={"mr-2 flex uppercase text-xl font-bold text-danger-dark items-center justify-end"}>
                    Lose
                </div>
            }
        } else {
            return <div></div>
        }
    }

    function getRightPlayerStatus() {
        if (round.status === "fighting") {
            return <div className={"ml-2 flex items-center justify-start"}>
                <Button variant={"danger"}
                        onClick={onEnemyWinButton}>
                    Enemy won
                </Button>
            </div>
        } else if (round.status === "finished") {
            if (round.winner_user_id !== currentUser.data?.id) {
                return <div
                    className={"ml-2 flex uppercase text-xl fw-bold text-success-dark items-center justify-start"}>
                    Win
                </div>
            } else {
                return <div
                    className={"ml-2 flex uppercase text-xl fw-bold text-danger-dark items-center justify-start"}>
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
            winner: props.draft.challenger_id === currentUser.data?.id ? props.draft.receiver_id : props.draft.challenger_id
        }

        winnerMutation.mutate(request)
    }

    const containerCN = classNames("grid-cols-5 grid", "w-100 p-0", "rounded", "bg-light-1 dark:bg-dark-1", "border border-light-3 dark:border-dark-3")
    return <div className={containerCN}>
        <div className={"flex items-center"}>{createPlayerElement()}</div>
        {getLeftPlayerStatus()}
        <div
            className={classNames("flex flex-col items-center", "p-2", "text-dark", "bg-secondary-light", "border-l border-r border-light-3 dark:border-dark-3")}>
            <span className={"fw-bold text-xl"}>Round {round.round_number}</span>
            <span>{getStatusDisplayMessage(round)}</span>
        </div>
        {getRightPlayerStatus()}
        <div className={"flex items-center justify-end"}>{createEnemyElement()}</div>
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

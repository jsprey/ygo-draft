import React, {useState} from "react";
import classNames from "classnames";
import {Link} from "react-router-dom";
import {ChallengeDraftState} from "../draft/challenge/ChallengeDraftPage";
import {useDraftChallenges} from "../api/hooks/drafts/useDraftChallenges";
import {Friend} from "../api/UserModel";
import {Draft} from "../api/Draft";
import DraftChallengeDetailModal from "../draft/challenge/DraftChallengeDetailModal";
import {useDrafts} from "../api/hooks/drafts/useDrafts";
import {useNavigate} from "react-router";
import Spinner from "../core/Spinner";
import Alert from "../core/Alert";
import Button from "../core/Button";
import {ChallengeUserPath} from "../routes/AppRouter";

export interface FriendListEntryProps {
    friend: Friend
    highlightBackground: boolean
    borderBottom: boolean
}

function FriendListEntry(props: FriendListEntryProps) {
    const draftChallenges = useDraftChallenges()
    const runningDrafts = useDrafts()
    const navigate = useNavigate()
    const [showChallengeModal, setShowChallengeModal] = useState<boolean>(false)
    const [inspectChallenge, setInspectChallenge] = useState<Draft>({} as Draft)

    const friendChallengeState: ChallengeDraftState = {
        friendID: props.friend.id,
        friendName: props.friend.name
    }

    let actions = <></>
    if (draftChallenges.isLoading || runningDrafts.isLoading) {
        actions = <div className={"flex content-center"}>
            <Spinner/>
        </div>
    } else if (draftChallenges.error) {
        actions = <Alert variant={'danger'} className={"mb-0"}>Failed to load challenges!</Alert>
    } else if (runningDrafts.error) {
        actions = <Alert variant={'danger'} className={"mb-0"}>Failed to load drafts!</Alert>
    } else if (draftChallenges.data && runningDrafts.data) {
        const receivedChallenges: Draft[] = draftChallenges.data.drafts.filter(value => value.challenger_id === props.friend.id)
        const sendChallenges: Draft[] = draftChallenges.data.drafts.filter(value => value.receiver_id === props.friend.id)
        const currentlyRunningDrafts: Draft[] = runningDrafts.data.drafts.filter(value => value.receiver_id === props.friend.id || value.challenger_id === props.friend.id)

        if (currentlyRunningDrafts.length === 1) {
            // there is a running draft
            actions = <div className={"flex items-center"}>
                <span className={"mr-2"}>Running Draft: </span>
                <Button className={"!p-1"}
                        variant={"primary"}
                        onClick={() => {
                            navigate(`/draft/${currentlyRunningDrafts[0].id}`)
                        }
                        }>View</Button>
            </div>
        } else if (receivedChallenges.length === 1) {
            // there is a challenge
            actions = <div className={"flex items-center"}>
                <span className={"mr-2"}>You received a challenge: </span>
                <Button variant={"primary"}
                    className={"!p-1"} onClick={() => {
                    setInspectChallenge(receivedChallenges[0]);
                    setShowChallengeModal(true)
                }
                }>View</Button>
            </div>
        } else if (sendChallenges.length === 1) {
            // there is an outgoing challenge
            actions = <div className={"flex items-center"}>
                <span className={"mr-2"}>Challenge send. Waiting for response.</span>
            </div>
        } else {
            // no challenge
            actions = <Link state={friendChallengeState} to={ChallengeUserPath}>
                <Button variant={"primary"}
                        onClick={() => {
                            navigate(ChallengeUserPath, {state:friendChallengeState})
                        }}
                        className={"!p-1"}>
                    Challenge
                </Button>
            </Link>
        }
    }

    const cNames = classNames("flex justify-between p-2 border-l border-r border-border", props.highlightBackground ? "bg-light-1 dark:bg-dark-1" : "bg-light-2 dark:bg-dark-2", props.borderBottom ? "border-b" : "")
    return <div className={cNames}>
        <div className={"self-center dark:text-white"}>
            <b>
                {props.friend.name}
            </b>
        </div>
        <div>
            {actions}
        </div>
        <DraftChallengeDetailModal challenge={inspectChallenge} isShowing={showChallengeModal}
                                   setShow={setShowChallengeModal}/>
    </div>
}

export default FriendListEntry

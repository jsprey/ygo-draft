import React, {useState} from "react";
import {Alert, Nav, Spinner} from "react-bootstrap";
import classNames from "classnames";
import {Link} from "react-router-dom";
import {ChallengeDraftState} from "../draft/challenge/ChallengeDraftPage";
import {usePendingChallenges} from "../api/hooks/challenges/usePendingChallenges";
import {Friend} from "../api/UserModel";
import {DraftChallenge} from "../api/Draft";
import DraftChallengeDetailModal from "../draft/challenge/DraftChallengeDetailModal";

export interface FriendListEntryProps {
    friend: Friend
    highlightBackground: boolean
    borderBottom: boolean
}

function FriendListEntry(props: FriendListEntryProps) {
    const pendingChallenges = usePendingChallenges()
    const [showChallengeModal, setShowChallengeModal] = useState<boolean>(false)
    const [inspectChallenge, setInspectChallenge] = useState<DraftChallenge>({} as DraftChallenge)

    const friendChallengeState: ChallengeDraftState = {
        friendID: props.friend.id,
        friendName: props.friend.name
    }

    let actions = <></>
    if (pendingChallenges.isLoading) {
        actions = <div className={"flex align-content-center"}>
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (pendingChallenges.error) {
        actions = <Alert variant={"danger"} className={"mb-0"}>Failed to load challenges!</Alert>
    } else if (pendingChallenges.data) {
        const receivedChallenges: DraftChallenge[] = pendingChallenges.data.challenges.filter(value => value.challenger_id == props.friend.id)

        if (receivedChallenges.length == 1) {
            // there is a challenge
            actions = <div className={"flex align-items-center"}>
                <span className={"mr-2"}>You received a challenge: </span>
                <span className={"btn btn-primary"} onClick={() => {
                setInspectChallenge(receivedChallenges[0]);
                setShowChallengeModal(true)
                }
                }>View</span>
            </div>
        } else {
            // no challenge
            actions = <Nav.Link as={Link} state={friendChallengeState} to={"/challenge"}>
                <span className={"btn btn-primary"}>Challenge</span>
            </Nav.Link>
        }

    }

    const cNames = classNames("flex justify-content-between p-2 border-start border-end", props.highlightBackground ? "bg-blue-100 dark:bg-gray-700" : "bg-blue-50 dark:bg-gray-600", props.borderBottom ? "border-bottom" : "")
    return <div key={props.friend.id}
                className={cNames}>
        <div className={"align-self-center dark:text-white"}>
            <b>
                {props.friend.name}
            </b>
        </div>
        <div>
            {actions}
        </div>
        <DraftChallengeDetailModal challenge={inspectChallenge} isShowing={showChallengeModal} setShow={setShowChallengeModal}/>
    </div>
}

export default FriendListEntry

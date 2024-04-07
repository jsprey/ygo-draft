import React, {useState} from "react";
import {Alert, Nav, Spinner} from "react-bootstrap";
import classNames from "classnames";
import {Link} from "react-router-dom";
import {ChallengeDraftState} from "../draft/challenge/ChallengeDraftPage";
import {useDraftChallenges} from "../api/hooks/drafts/useDraftChallenges";
import {Friend} from "../api/UserModel";
import {Draft} from "../api/Draft";
import DraftChallengeDetailModal from "../draft/challenge/DraftChallengeDetailModal";
import {useDrafts} from "../api/hooks/drafts/useDrafts";
import {useNavigate} from "react-router";

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
        actions = <div className={"flex align-content-center"}>
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (draftChallenges.error) {
        actions = <Alert variant={"danger"} className={"mb-0"}>Failed to load challenges!</Alert>
    } else if (runningDrafts.error) {
        actions = <Alert variant={"danger"} className={"mb-0"}>Failed to load drafts!</Alert>
    } else if (draftChallenges.data && runningDrafts.data) {
        const receivedChallenges: Draft[] = draftChallenges.data.drafts.filter(value => value.challenger_id === props.friend.id)
        const sendChallenges: Draft[] = draftChallenges.data.drafts.filter(value => value.receiver_id === props.friend.id)
        const currentlyRunningDrafts: Draft[] = runningDrafts.data.drafts.filter(value => value.receiver_id === props.friend.id || value.challenger_id === props.friend.id)

        if (currentlyRunningDrafts.length === 1) {
            // there is a running draft
            actions = <div className={"flex align-items-center"}>
                <span className={"mr-2"}>Running Draft: </span>
                <span className={"btn btn-primary"} onClick={() => {
                    navigate(`/draft/${currentlyRunningDrafts[0].id}`)
                }
                }>View</span>
            </div>
        } else if (receivedChallenges.length === 1) {
            // there is a challenge
            actions = <div className={"flex align-items-center"}>
                <span className={"mr-2"}>You received a challenge: </span>
                <span className={"btn btn-primary"} onClick={() => {
                    setInspectChallenge(receivedChallenges[0]);
                    setShowChallengeModal(true)
                }
                }>View</span>
            </div>
        } else if (sendChallenges.length === 1) {
            // there is an outgoing challenge
            actions = <div className={"flex align-items-center"}>
                <span className={"mr-2"}>Challenge send. Waiting for response.</span>
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
        <DraftChallengeDetailModal challenge={inspectChallenge} isShowing={showChallengeModal}
                                   setShow={setShowChallengeModal}/>
    </div>
}

export default FriendListEntry

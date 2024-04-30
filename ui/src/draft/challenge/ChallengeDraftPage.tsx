import React, {useState} from "react";
import {useLocation} from "react-router-dom";
import PageSettings from "../local/PageSettings";
import {usePrompt} from "../../api/hooks/usePromptBlocker";
import {PostDraftChallengeRequest, useChallengeUser} from "../../api/hooks/drafts/useChallengeUser";
import {enqueueSnackbar} from "notistack";
import {DraftSettings} from "../../api/Draft";
import {useNavigate} from "react-router";
import Spinner from "../../core/Spinner";
import Alert from "../../core/Alert";
import {UserPath} from "../../routes/AppRouter";

export type ChallengeDraftState = {
    friendID: number,
    friendName: string
}

function ChallengeDraftPage() {
    const navigate = useNavigate()
    const [challengedSend, setChallengedSend] = useState<boolean>(false)
    const location = useLocation();

    usePrompt("You did not challenge your friend. Are you sure you want to leave?", !challengedSend);

    const onMutationError = () => {
        enqueueSnackbar('Failed to send challenge. Try again an/or contact the support.', {
            autoHideDuration: 6000,
            variant: 'error'
        })
    }
    const onMutationSuccess = () => {
        enqueueSnackbar('Challenge send', {
            autoHideDuration: 6000,
            variant: 'success'
        })
        navigate(UserPath)
    }
    const challengeFriendMutation = useChallengeUser({onSuccess: onMutationSuccess, onError: onMutationError});

    function sendChallenge(settings: DraftSettings, state: ChallengeDraftState) {
        const challengeFriendRequest: PostDraftChallengeRequest = {
            friend_id: state.friendID,
            settings: settings
        }

        setChallengedSend(true)
        challengeFriendMutation.mutate(challengeFriendRequest)
    }

    if (!location.state) {
        return <div className={"pt-2 pb-1"}>
            <Alert variant={'danger'}>Missing friend to invite. Try again.</Alert>
        </div>
    } else {
        const state = location.state as ChallengeDraftState;
        return <div>
            <div className={"flex justify-content-center pt-4 pb-3"}>
                <p className={"text-5xl align-text-center uppercase dark:text-neutral-50"}>Challenge: {state.friendName}</p>
            </div>

            {challengeFriendMutation.isLoading ?
                <div className={"dark:text-neutral-50"}><Spinner/> Sending Challenge
                </div> :
                <PageSettings local={false} submitButtonName={"Challenge"}
                              onSettingsSubmit={(settings: DraftSettings) => sendChallenge(settings, state)}/>}
        </div>
    }

}


export default ChallengeDraftPage

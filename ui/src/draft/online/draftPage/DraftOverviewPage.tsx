import React, {useState} from "react";
import {useParams} from "react-router";
import {useDraft} from "../../../api/hooks/drafts/useDraft";
import DraftHeader from "./DraftHeader";
import DraftRoundList from "./DraftRoundList";
import {useCurrentUser} from "../../../api/hooks/users/useUser";
import {useFriend} from "../../../api/hooks/friends/useFriend";
import {useSurrenderDraft} from "../../../api/hooks/drafts/useSurrenderDraft";
import {useQueryClient} from "react-query";
import Spinner from "../../../core/Spinner";
import Alert from "../../../core/Alert";
import ConfirmModal from "../../../core/ConfirmModal";
import Button from "../../../core/Button";
import {ShowErrorSnack, ShowSuccessfulSnack} from "../../../core/Snacks";

type DraftOverviewPageParams = {
    id: string;
};

function DraftOverviewPage() {
    const [showConfirm, setShowConfirm] = useState<boolean>(false)
    let params = useParams<DraftOverviewPageParams>() as DraftOverviewPageParams;
    const draftRequest = useDraft(params.id)
    const queryClient = useQueryClient()
    const surrenderDraft = useSurrenderDraft({
        onSuccess: () => {
            ShowSuccessfulSnack("Draft surrendered!")
            if (draftRequest.data) {
                queryClient.refetchQueries({queryKey: ["draft", "" + draftRequest.data.id]})
            }
        }, onError: () => ShowErrorSnack("Failed to surrender draft!")
    })
    const userRequest = useCurrentUser()

    function getFriendID(): number {
        if (!draftRequest.isSuccess || !userRequest.isSuccess) {
            return -1;
        }

        return draftRequest.data.challenger_id === userRequest.data.id ? draftRequest.data.receiver_id : draftRequest.data.challenger_id;
    }

    const friendRequest = useFriend(getFriendID(), {enabled: draftRequest.isSuccess && userRequest.isSuccess})

    function createConfirmModal(): JSX.Element {
        return <ConfirmModal show={showConfirm}
                             setShow={setShowConfirm}
                             title={"Surrender Draft"}
                             confirmName={"Surrender"}
                             description={"You are about to surrender this draft. Are you sure? This cannot be reversed."}
                             onConfirm={() => {
                                 if (draftRequest.data) {
                                     surrenderDraft.mutate(draftRequest.data.id)
                                 }
                             }}/>
    }

    let content = <></>
    if (draftRequest.isLoading || userRequest.isLoading || friendRequest.isLoading
    ) {
        content = <Spinner/>
    } else if (draftRequest.isError) {
        content = <Alert className={"mb-0"} variant={'danger'}>Failed to load current draft!</Alert>
    } else if (userRequest.isError) {
        content = <Alert className={"mb-0"} variant={'danger'}>Failed to load current user!</Alert>
    } else if (friendRequest.isError) {
        content = <Alert className={"mb-0"} variant={'danger'}>Failed to load enemy user!</Alert>
    } else if (draftRequest.data && userRequest.data && friendRequest.data) {
        const currentDraft = draftRequest.data
        const isDraw = currentDraft.winner_user_id === -100
        const winnerName = userRequest.data.id === currentDraft.winner_user_id && !isDraw ? userRequest.data.display_name : friendRequest.data.name
        content = <div className={"w-full"}>
            <DraftHeader draft={currentDraft} player={userRequest.data} enemy={friendRequest.data}/>
            <DraftRoundList draft={currentDraft}/>
            {(currentDraft.status === "finished" || currentDraft.status === "surrender") ?
                <div
                    className={"mt-2 p-2 bg-gray-200 dark:bg-gray-600 text-6xl flex justify-content-center text-uppercase"}>
                    {isDraw ? "Draw" : `Winner: ${winnerName}`}
                </div> : <></>}
            {currentDraft.status === "running" ? <div className={"mt-2 flex justify-end"}>
                <Button variant={"danger"} onClick={() => setShowConfirm(true)}>
                    {surrenderDraft.isLoading ? <Spinner/> : "Surrender"}
                </Button>
            </div> : <></>}
            {createConfirmModal()}
        </div>
    }

    return <div className={"flex text-dark dark:text-light"}>{content}</div>
}


export default DraftOverviewPage

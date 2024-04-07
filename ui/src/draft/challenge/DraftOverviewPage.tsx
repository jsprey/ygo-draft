import React from "react";
import {Alert} from "react-bootstrap";
import {useParams} from "react-router";
import {useDraft} from "../../api/hooks/drafts/useDraft";
import DraftHeader from "./DraftHeader";
import DraftRoundList from "./DraftRoundList";
import {useCurrentUser} from "../../api/hooks/users/useUser";
import {useFriend} from "../../api/hooks/friends/useFriend";

type DraftOverviewPageParams = {
    id: string;
};

function DraftOverviewPage() {
    let params = useParams<DraftOverviewPageParams>() as DraftOverviewPageParams;
    const draftRequest = useDraft(params.id)
    const userRequest = useCurrentUser()

    function getFriendID(): number {
        if (!draftRequest.isSuccess || !userRequest.isSuccess) {
            return -1;
        }

        return draftRequest.data.challenger_id === userRequest.data.id ? draftRequest.data.receiver_id : draftRequest.data.challenger_id;
    }

    const friendRequest = useFriend(getFriendID(), {enabled: draftRequest.isSuccess && userRequest.isSuccess})

    let content = <></>
    if (draftRequest.isLoading && userRequest.isLoading && friendRequest.isLoading) {
        content = <p className={"placeholder placeholder-glow vw-100 vh-100"}></p>
    } else if (draftRequest.isError) {
        content = <Alert variant={"danger"}>Failed to load current draft!</Alert>
    } else if (userRequest.isError) {
        content = <Alert variant={"danger"}>Failed to load current user!</Alert>
    } else if (friendRequest.isError) {
        content = <Alert variant={"danger"}>Failed to load enemy user!</Alert>
    } else if (draftRequest.data && userRequest.data && friendRequest.data) {

        content = <div className={"flex dark:text-white pb-2 pt-2"}>
            <div className={"w-100"}>
                <DraftHeader draft={draftRequest.data} player={userRequest.data} enemy={friendRequest.data}/>
                <DraftRoundList draft={draftRequest.data}/>
            </div>
        </div>
    }

    return content
}


export default DraftOverviewPage

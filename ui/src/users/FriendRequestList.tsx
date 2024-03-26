import React from "react";
import {Spinner} from "react-bootstrap";
import {useFriendRequests} from "../api/hooks/friends/useFriendRequests";
import {getTimeDifferenceString} from "../api/UserModel";
import {useFriendsAcceptRequest} from "../api/hooks/friends/useFriendsAcceptRequest";
import {useQueryClient} from "react-query";
import {enqueueSnackbar} from "notistack";
import classNames from "classnames";

function FriendRequestList() {
    const {data, isLoading, error} = useFriendRequests()
    const queryClient = useQueryClient()

    const onMutationError = () => {
        enqueueSnackbar('Failed to accept friend request. Try again an/or contact the support.', {
            autoHideDuration: 6000,
            variant: 'error'
        })
    }
    const onMutationSuccess = () => {
        enqueueSnackbar('Friend added.', {
            autoHideDuration: 6000,
            variant: 'success'
        })
        queryClient.invalidateQueries({queryKey: "friends"})
        queryClient.invalidateQueries({queryKey: "friendRequests"})
    }
    const sendFriendRequestMutation = useFriendsAcceptRequest({
        onSuccess: onMutationSuccess,
        onError: onMutationError
    })

    let friendRequestsContainer = <></>
    if (isLoading) {
        friendRequestsContainer = <div className={"flex align-content-center"}>
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (error) {
        friendRequestsContainer = <div className={"flex align-content-center"}>
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get friends!</div>
        </div>
    } else if (data) {
        let requestEntries: JSX.Element[] = [];
        let isHighlightedBackground = true
        data.forEach((request, index) => {
            let cNames = classNames("flex p-2 border-start border-end", isHighlightedBackground ? "bg-blue-100 dark:bg-gray-700" : "bg-blue-50 dark:bg-gray-600", index === data.length-1 ? "border-bottom" : "")
            let entry = <div key={request.id}
                             className={cNames}>
                <span
                    className={"rounded bg-cyan-700 text-white p-1 mr-2"}>{getTimeDifferenceString(request.invitation_date)}</span>
                <div className={"align-self-center dark:text-white"}>You got a friend request from <b>{request.name}</b>
                </div>
                <button className={"btn btn-success p-1 ml-auto"} onClick={() => {
                    sendFriendRequestMutation.mutate(request.id)
                }}>
                    {sendFriendRequestMutation.isLoading ? <Spinner size={"sm"} animation={"border"}/> : "Accept"}
                </button>
                <button className={"btn btn-danger p-1 ml-1"}>Decline</button>
            </div>
            isHighlightedBackground = !isHighlightedBackground
            requestEntries.push(entry)
        })

        if (requestEntries.length === 0) {
            friendRequestsContainer = <></>
        } else {
            friendRequestsContainer = <div>
                {requestEntries}
            </div>
        }
    }

    return friendRequestsContainer
}

export default FriendRequestList

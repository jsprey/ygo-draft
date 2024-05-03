import React from "react";
import {useFriendRequests} from "../api/hooks/friends/useFriendRequests";
import {getTimeDifferenceString} from "../api/UserModel";
import {useFriendsAcceptRequest} from "../api/hooks/friends/useFriendsAcceptRequest";
import {useQueryClient} from "react-query";
import {enqueueSnackbar} from "notistack";
import classNames from "classnames";
import Spinner from "../core/Spinner";
import Button from "../core/Button";

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
        queryClient.refetchQueries({queryKey: "friends"})
        queryClient.refetchQueries({queryKey: "friendRequests"})
    }
    const sendFriendRequestMutation = useFriendsAcceptRequest({
        onSuccess: onMutationSuccess,
        onError: onMutationError
    })

    let friendRequestsContainer = <></>
    if (isLoading) {
        friendRequestsContainer = <div className={"flex align-center"}>
            <Spinner/>
        </div>
    } else if (error) {
        friendRequestsContainer = <div className={"flex align-center"}>
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get friends!</div>
        </div>
    } else if (data) {
        let requestEntries: JSX.Element[] = [];
        let isHighlightedBackground = true
        data.forEach((request, index) => {
            let cNames = classNames("flex p-2 border-l border-r border-border", isHighlightedBackground ? "bg-light-1 dark:bg-dark-1" : "bg-light-2 dark:bg-dark-2", index === data.length-1 ? "border-b" : "")
            let entry = <div key={request.id}
                             className={cNames}>
                <span
                    className={"rounded bg-cyan-700 text-white p-1 mr-2"}>{getTimeDifferenceString(request.invitation_date)}</span>
                <div className={"self-center dark:text-white"}>You got a friend request from <b>{request.name}</b>
                </div>
                <Button className={"btn btn-success !p-1 mr-1 ml-auto"}
                        variant={"success"}
                        onClick={() => {
                    sendFriendRequestMutation.mutate(request.id)
                }}>
                    {sendFriendRequestMutation.isLoading ? <Spinner/> : "Accept"}
                </Button>
                <Button variant={"danger"} className={"!p-1"}>Decline</Button>
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

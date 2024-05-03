import React from "react";
import {useFriends} from "../api/hooks/friends/useFriends";
import FriendListEntry from "./FriendListEntry";
import Spinner from "../core/Spinner";

function FriendList() {
    const {data, isLoading, error} = useFriends()

    let content = <></>
    if (isLoading) {
        content = <div className={"flex content-center"}>
            <Spinner/>
        </div>
    } else if (error) {
        content = <div className={"flex content-center"}>
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get friends!</div>
        </div>
    } else if (data && data.length > 0) {
        let friendsEntries: JSX.Element[] = [];
        let isHighlightedBackground = true
        data.forEach((friend, index) => {
            let entry = <div key={`friend-list-entry-${friend.id}`}>
                <FriendListEntry friend={friend} highlightBackground={isHighlightedBackground}
                                 borderBottom={index === data.length - 1}/>
            </div>
            isHighlightedBackground = !isHighlightedBackground
            friendsEntries.push(entry)

            content = <div>{friendsEntries}</div>
        })
    } else if (data && data.length === 0) {
        content = <div className={"flex bg-blue-100 dark:bg-gray-700 p-2 border-l border-r border-b border-border"}>
            <div className={"self-center dark:text-white"}>
                <b>You currently have no friends.</b>
            </div>
        </div>
    }

    return content
}

export default FriendList

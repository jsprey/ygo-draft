import React from "react";
import {Spinner} from "react-bootstrap";
import {useFriends} from "../api/hooks/friends/useFriends";
import FriendListEntry from "./FriendListEntry";

function FriendList() {
    const {data, isLoading, error} = useFriends()

    let content = <></>
    if (isLoading) {
        content = <div className={"flex align-content-center"}>
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (error) {
        content = <div className={"flex align-content-center"}>
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
        content = <div className={"flex bg-blue-100 dark:bg-gray-700 p-2 border-start border-end border-bottom"}>
            <div className={"align-self-center dark:text-white"}><b>You currently have no friends.</b>
            </div>
        </div>
    }

    return content
}

export default FriendList

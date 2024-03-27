import React from "react";
import {useCurrentUser} from "../api/hooks/users/useUser";
import {Spinner} from "react-bootstrap";
import FriendList from "./FriendList";
import FriendRequestList from "./FriendRequestList";
import AddNewFriendWidget from "./AddNewFriendWidget";

function UserPage() {
    const {data, isLoading, error} = useCurrentUser()

    let contentUser = <></>
    if (isLoading) {
        contentUser = <div className={"flex align-content-center"}>
            <Spinner animation={"grow"} size={"sm"}/>
        </div>
    } else if (error) {
        contentUser = <div className={"flex align-content-center"}>
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get user!</div>
        </div>
    } else if (data) {
        contentUser = <>
            <div className={"flex"}>
                <span className={"mr-3"}>Username: </span>
                <span>{data.display_name}</span>
            </div>
            <div className={"flex"}>
                <span className={"mr-3"}>E-Mail: </span>
                <span>{data.email}</span>
            </div>
            <div className={"flex"}>
                <span className={"mr-3"}>Display-Name: </span>
                <span>{data.display_name}</span>
            </div>
        </>
    }


    return <div className={"p-2 dark:text-white"}>
        <h1>User Profile</h1>
        {contentUser}
        <h2 className={"mt-3"}>Friends</h2>
        <AddNewFriendWidget/>
        <FriendRequestList/>
        <FriendList/>
        <div className={"mb-2"}></div>
    </div>
}

export default UserPage

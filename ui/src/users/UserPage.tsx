import React from "react";
import {useCurrentUser} from "../api/hooks/users/useUser";
import FriendList from "./FriendList";
import FriendRequestList from "./FriendRequestList";
import AddNewFriendWidget from "./AddNewFriendWidget";
import Spinner from "../core/Spinner";
import {GetUserReponse} from "../api/UserModel";

function UserPage() {
    const {data, isLoading, error} = useCurrentUser()

    function createUserInformationTable(data: GetUserReponse) {
        return <table>
            <tbody>
            <tr>
                <td>Display Name:</td>
                <td>{data.display_name}</td>
            </tr>
            <tr>
                <td>Email:</td>
                <td>{data.email}</td>
            </tr>
            </tbody>
        </table>
    }

    let contentUser = <></>
    if (isLoading) {
        contentUser = <div className={"flex align-content-center"}>
            <Spinner/>
        </div>
    } else if (error) {
        contentUser = <div className={"flex align-content-center"}>
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get user!</div>
        </div>
    } else if (data) {
        contentUser = createUserInformationTable(data)
    }


    return <div className={"dark:text-white"}>
        <div className={"mb-1 text-xl font-bold "}>
            User Profile
        </div>
        {contentUser}
        <div className={"mt-3 mb-1 text-xl font-bold "}>
            Friends
        </div>
        <AddNewFriendWidget/>
        <FriendRequestList/>
        <FriendList/>
        <div className={"mb-2"}></div>
    </div>
}

export default UserPage

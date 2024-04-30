import React from "react";
import {useCurrentUser} from "../api/hooks/users/useUser";
import UserManagement from "./UserManagement";
import AddNewUserWidget from "./AddNewUserWidget";
import Spinner from "../core/Spinner";
import Alert from "../core/Alert";

function AdminPage() {
    const {data, isLoading, error} = useCurrentUser()

    let content = <></>

    if (isLoading) {
        content = <Spinner/>
    } else if (error) {
        content = <Alert variant={'danger'}>There seems to be an issue!</Alert>
    } else if (data && !data.is_admin) {
        content = <Alert variant={'danger'}>No access for you!</Alert>
    } else if (data && data.is_admin) {
        content = <div className={"dark:text-white"}>
            <AddNewUserWidget/>
            <div className={"mb-3 mt-3 ml-1 text-xl uppercase text-dark dark:text-light fw-bold"}>Users</div>
            <UserManagement/>
        </div>
    }

    return content
}

export default AdminPage

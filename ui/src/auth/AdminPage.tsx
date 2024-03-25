import React from "react";
import {useCurrentUser} from "../api/hooks/useUser";
import {Alert, Spinner} from "react-bootstrap";
import UserManagement from "./UserManagement";

function AdminPage() {
    const {data, isLoading, error} = useCurrentUser()

    let content = <></>

    if (isLoading) {
        content = <Spinner animation={"border"}/>
    } else if (error) {
        content = <Alert variant={"danger"}>There seems to be an issue!</Alert>
    } else if (data && !data.is_admin) {
        content = <Alert variant={"danger"}>No access for you!</Alert>
    } else if (data && data.is_admin) {
        content = <div className={"p-2 dark:text-white"}>
            <h1>Administration</h1>
            <h2 className={"mt-3"}>Users</h2>
            <UserManagement/>
            <div className={"mb-2"}></div>
        </div>
    }

    return content
}

export default AdminPage

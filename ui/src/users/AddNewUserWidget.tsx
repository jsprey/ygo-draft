import React, {useState} from "react";
import {AddUserPayload, useAddUser} from "../api/hooks/users/useAddUser";
import {enqueueSnackbar} from "notistack";
import {useQueryClient} from "react-query";
import classNames from "classnames";
import Spinner from "../core/Spinner";
import YgoIcon from "../core/YgoIcon";

function AddNewUserWidget() {
    const [collapsed, setCollapsed] = useState<boolean>(true)
    const [email, setEmail] = useState<string>("")
    const [displayName, setDisplayName] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const [isAdministrator, setIsAdministrator] = useState<boolean>(false)
    const queryClient = useQueryClient()

    const onMutationError = () => {
        enqueueSnackbar(`Failed to create user.`, {
            autoHideDuration: 6000,
            variant: 'error'
        })
    }
    const onMutationSuccess = () => {
        enqueueSnackbar('User created', {
            autoHideDuration: 6000,
            variant: 'success'
        })
        queryClient.refetchQueries({queryKey: "users"})
    }
    const addUserMutation = useAddUser({onSuccess: onMutationSuccess, onError: onMutationError});

    function getWidgetBody() {
        return <>
            <div className="mb-2 flex flex-col">
                <label className="font-bold" htmlFor={"newUser_email"}>Email address</label>
                <input value={email} type="email" className="" id={"newUser_email"}
                       onChange={event => setEmail(event.target.value)}/>
            </div>
            <div className="mb-2 flex flex-col">
                <label className="font-bold" htmlFor={"newUser_name"}>Display Name</label>
                <input value={displayName} type="email" className="" id={"newUser_name"}
                       onChange={event => setDisplayName(event.target.value)}/>
            </div>
            <div className="mb-2 flex flex-col">
                <label className="font-bold" htmlFor={"newUser_password"}>Password</label>
                <input value={password} type="password" className="" id={"newUser_password"}
                       onChange={event => setPassword(event.target.value)}/>
            </div>
            <div className="mb-2 form-check">
                <input checked={isAdministrator} type="checkbox" className="form-check-input" id={"newUser_isAdmin"}
                       onChange={event => setIsAdministrator(event.target.checked)}/>
                <label className="ml-2 form-check-label" htmlFor={"newUser_isAdmin"}>Administrator</label>
            </div>
            <button disabled={addUserMutation.isLoading} className="btn btn-success" onClick={() => {
                const payload: AddUserPayload = {
                    email: email,
                    password: password,
                    display_name: displayName,
                    is_admin: isAdministrator
                }
                addUserMutation.mutate(payload)
            }}>
                {addUserMutation.isLoading ? <Spinner/> : <>Create</>}
            </button>
        </>;
    }

    return <div>
        <div className={classNames("flex", collapsed ? "" : "mb-2")}>
            <div className={classNames("ml-1 mr-2 text-xl uppercase text-dark dark:text-light fw-bold")}>Add New
                Users
            </div>
            <YgoIcon icon={collapsed ? "collapse" : "expand"}
                     onClick={() => setCollapsed(!collapsed)}
                     size={18}
                     classNames={"text-dark hover:text-dark-1 active:text-dark-2 dark:fill-light hover:dark:text-light-1 active:dark:text-light-2"}/>
        </div>
        {collapsed ? <></> : getWidgetBody()}
    </div>
}

export default AddNewUserWidget

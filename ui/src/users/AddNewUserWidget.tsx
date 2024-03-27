import React, {useState} from "react";
import SvgIconButton from "../core/SvgIconButton";
import {AddUserPayload, useAddUser} from "../api/hooks/users/useAddUser";
import {enqueueSnackbar} from "notistack";
import {useQueryClient} from "react-query";
import {Spinner} from "react-bootstrap";
import classNames from "classnames";

const CollapsedIcon = <SvgIconButton size={18} classNames={"dark:fill-gray-400 dark:fill-neutral-50"}>
    <path
        d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z"/>
</SvgIconButton>
const ExpandIcon = <SvgIconButton size={18} classNames={"fill-neutral-900 dark:fill-neutral-50"}>
        <path
            d="M7.247 11.14 2.451 5.658C1.885 5.013 2.345 4 3.204 4h9.592a1 1 0 0 1 .753 1.659l-4.796 5.48a1 1 0 0 1-1.506 0z"/>
</SvgIconButton>

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
        queryClient.invalidateQueries({queryKey: "users"})
    }
    const addUserMutation = useAddUser({onSuccess: onMutationSuccess, onError: onMutationError});

    function getWidgetBody() {
        return <>
            <div className="mb-2">
                <label className="form-label" htmlFor={"newUser_email"}>Email address</label>
                <input value={email} type="email" className="form-control" id={"newUser_email"}
                       onChange={event => setEmail(event.target.value)}/>
            </div>
            <div className="mb-2">
                <label className="form-label" htmlFor={"newUser_name"}>Display Name</label>
                <input value={displayName} type="email" className="form-control" id={"newUser_name"}
                       onChange={event => setDisplayName(event.target.value)}/>
            </div>
            <div className="mb-2">
                <label className="form-label" htmlFor={"newUser_password"}>Password</label>
                <input value={password} type="password" className="form-control" id={"newUser_password"}
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
                {addUserMutation.isLoading ? <Spinner animation={"border"} size={"sm"}/> : <>Create</>}
            </button>
        </>;
    }

    return <div>
        <div className={classNames("flex", collapsed ? "" : "mb-3")} onClick={() => setCollapsed(!collapsed)}>
            <div className={classNames("mt-3 ml-1 mr-2 text-xl uppercase text-gray-700 dark:text-gray-400 fw-bold")}>Add New Users</div>
            <div className={"mt-3 align-self-center"}>
                {collapsed ? CollapsedIcon : ExpandIcon}
            </div>
        </div>
        {collapsed ? <></> : getWidgetBody()}
    </div>
}

export default AddNewUserWidget

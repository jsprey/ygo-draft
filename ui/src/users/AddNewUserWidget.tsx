import React, {useState} from "react";
import {AddUserPayload, useAddUser} from "../api/hooks/users/useAddUser";
import {enqueueSnackbar} from "notistack";
import {useQueryClient} from "react-query";
import classNames from "classnames";
import Spinner from "../core/Spinner";
import YgoIcon from "../core/YgoIcon";
import Input from "../core/Input";
import Button from "../core/Button";

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
            <div className="mb-3 flex flex-col">
                <Input value={email}
                       type="email"
                       rootClassNames={"w-full"}
                       inputClassNames={"w-full"}
                       label={"E-Mail"}
                       tooltip={"The email that the user uses to log into the app."}
                       onChange={event => setEmail(event.target.value)}/>
            </div>
            <div className="mb-3 flex flex-col">
                <Input value={displayName}
                       type="email"
                       rootClassNames={"w-full"}
                       inputClassNames={"w-full"}
                       label={"Display Name"}
                       tooltip={"The name that is displayed in the application for the user."}
                       onChange={event => setDisplayName(event.target.value)}/>
            </div>
            <div className="mb-3 flex flex-col">
                <Input value={password}
                       type="password"
                       rootClassNames={"w-full"}
                       inputClassNames={"w-full"}
                       label={"Password"}
                       tooltip={"The password that the user uses to log into the app."}
                       onChange={event => setPassword(event.target.value)}/>
            </div>
            <div className="mb-3 form-check">
                <Input value={isAdministrator ? "true" : "false"}
                       type="checkbox"
                       label={"Administrator"}
                       rootClassNames={"w-full"}
                       tooltip={"Decides whether the user is an administrator or not."}
                       onChange={event => setIsAdministrator(event.target.checked)}/>
            </div>
            <Button variant={"primary"} disabled={addUserMutation.isLoading} className="btn btn-success" onClick={() => {
                const payload: AddUserPayload = {
                    email: email,
                    password: password,
                    display_name: displayName,
                    is_admin: isAdministrator
                }
                addUserMutation.mutate(payload)
            }}>
                {addUserMutation.isLoading ? <Spinner/> : <>Create</>}
            </Button>
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

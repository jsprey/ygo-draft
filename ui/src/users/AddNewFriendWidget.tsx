import React, {useState} from "react";
import {useSendFriendRequest} from "../api/hooks/friends/useSendFriendRequest";
import {enqueueSnackbar} from "notistack";
import {useCurrentUser} from "../api/hooks/users/useUser";
import Spinner from "../core/Spinner";
import Button from "../core/Button";
import classNames from "classnames";

function AddNewFriendWidget() {
    var user = useCurrentUser();
    const [newFriendName, setNewFriendName] = useState<string>("")
    const [invalidInput, setInvalidInput] = useState<string>("Value cannot be empty.")
    const [showError, setShowError] = useState<boolean>(false)

    const onMutationError = () => {
        enqueueSnackbar('Failed to send friend request. Try again an/or contact the support.', {
            autoHideDuration: 6000,
            variant: 'error'
        })
    }
    const onMutationSuccess = () => {
        enqueueSnackbar('Friend request send', {
            autoHideDuration: 6000,
            variant: 'success'
        })
    }
    let sendFriendRequest = useSendFriendRequest({onSuccess: onMutationSuccess, onError: onMutationError});

    function validateInput(value: string) {
        if (value === "") {
            setInvalidInput("Value cannot be empty.")
            return
        }

        if (value === user.data?.email) {
            setInvalidInput("You cannot send yourself a friend request.")
            return
        }

        if (!value.match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        )) {
            setInvalidInput("Value needs to be a valid email.")
            return
        }

        setInvalidInput("")
    }

    const errorBlock = <div className={"flex bg-rose-100 dark:bg-rose-700 p-2 border-l border-r border-b border-border"}>
        <div className={"self-center dark:text-white"}><b>{invalidInput}</b>
        </div>
    </div>

    return <div>
        <div className={"flex"}>
            <input
                className={classNames("flex-grow pl-2", "outline-none", "text-dark dark:text-light", "placeholder-dark-3 dark:placeholder-light-3", "bg-light-3 dark:bg-dark-3", "rounded-tl-lg", "border-b border-l border-t border-border")}
                placeholder={"add a new friend"}
                value={newFriendName}
                onBlur={() => setShowError(false)}
                onFocus={() => setShowError(true)}
                onChange={event => {
                    setNewFriendName(event.target.value)
                    validateInput(event.target.value)
                }
                }>
            </input>
            <Button disabled={invalidInput !== ""}
                    className={"!rounded-l-none !rounded-br-none !border-norder"}
                    variant={"success"}
                    onClick={() => {
                        sendFriendRequest.mutate(newFriendName)
                        setNewFriendName("")
                    }
                    }>
                {sendFriendRequest.isLoading ? <Spinner/> : <span>Add Friend</span>}
            </Button>
        </div>
        {(invalidInput !== "" && showError) ? errorBlock : <></>}
    </div>
}

export default AddNewFriendWidget

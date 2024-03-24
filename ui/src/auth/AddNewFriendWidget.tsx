import React, {useState} from "react";
import {useSendFriendRequest} from "../api/hooks/friends/useSendFriendRequest";
import {Spinner} from "react-bootstrap";
import {enqueueSnackbar} from "notistack";
import {useCurrentUser} from "../api/hooks/useUser";

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
        console.log(`Validate the value '${value}'`)

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

    const errorBlock = <div className={"flex bg-rose-100 dark:bg-rose-700 p-2 border-start border-end border-bottom"}>
        <div className={"align-self-center dark:text-white"}><b>{invalidInput}</b>
        </div>
    </div>

    return <div>
        <div className={"flex"}>
            <input
                className={"rounded-tl-lg flex-grow-1 border-bottom border-start border-top focus:no-border pl-2 dark:text-white bg-gray-200 dark:bg-gray-600 is-invalid"}
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
            <button disabled={invalidInput !== ""}
                    className={"bg-ygo-success hover:bg-ygo-success-hover active:bg-ygo-success-active disabled:bg-ygo-success-disabled rounded-tr-lg p-2 border-bottom border-end border-top text-neutral-50 disabled:text-gray-500"}
                    onClick={() => {
                        sendFriendRequest.mutate(newFriendName)
                        setNewFriendName("")
                    }
                    }>
                {sendFriendRequest.isLoading ? <Spinner size={"sm"} animation={"border"}/> : <span>Add Friend</span>}
            </button>
        </div>
        {(invalidInput !== "" && showError) ? errorBlock : <></>}
    </div>
}

export default AddNewFriendWidget

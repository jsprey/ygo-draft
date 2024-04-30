import React, {useState} from "react";
import {User, useUsers} from "../api/hooks/users/useUsers";
import classNames from "classnames";
import {useQueryClient} from "react-query";
import {enqueueSnackbar} from "notistack";
import {useDeleteUser} from "../api/hooks/users/useDeleteUser";
import Spinner from "../core/Spinner";
import ConfirmModal from "../core/ConfirmModal";
import Alert from "../core/Alert";
import YgoIcon from "../core/YgoIcon";
import Button from "../core/Button";

function UserManagement() {
    const [currentPage, setCurrentPage] = useState<number>(0);
    const [currentPageSize] = useState<number>(15);
    const {data, isLoading, error} = useUsers(currentPage, currentPageSize);
    const queryClient = useQueryClient()

    const onMutationError = () => {
        enqueueSnackbar('Failed to delete user. Try again an/or contact the support.', {
            autoHideDuration: 6000,
            variant: 'error'
        })
        setToBeDeletedUserEmail("")
    }
    const onMutationSuccess = () => {
        enqueueSnackbar('User deleted.', {
            autoHideDuration: 6000,
            variant: 'success'
        })
        setToBeDeletedUserEmail("")
        queryClient.refetchQueries({queryKey: "users"})
    }
    const deleteUserMutation = useDeleteUser({
        onSuccess: onMutationSuccess,
        onError: onMutationError
    })
    const [showAbortDialog, setShowAbortDialog] = useState<boolean>(false)
    const [toBeDeletedUserEmail, setToBeDeletedUserEmail] = useState<string>("")

    let tableEntries: JSX.Element[] = []
    if (error) {
        tableEntries.push(<tr>
            <td colSpan={5}><Alert variant={'danger'}>Failed to load users!</Alert></td>
        </tr>)
    } else if (isLoading) {
        tableEntries = createPlaceholderEntries()
    } else if (data) {
        tableEntries = createTableWithActualData(data.users)
    }

    function deleteSelectedUser() {
        if (toBeDeletedUserEmail !== "") {
            deleteUserMutation.mutate(toBeDeletedUserEmail)
        }
    }

    function createTableWithActualData(users: User[]): JSX.Element[] {
        const tableEntries: JSX.Element[] = []

        users.forEach(user => {
            const entry = <tr key={`row-${user.id}`}
                              className={"odd:bg-light-1 odd:dark:bg-dark-1 even:bg-light-2 even:dark:bg-dark-2 border-l border-r border-dark dark:border-light"}>
                <td className="px-6 py-2">{user.id}</td>
                <td className="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">{user.email}</td>
                <td className="px-6 py-2">{user.display_name}</td>
                <td className="px-6 py-2">{user.is_admin ? "Admin" : "User"}</td>
                <td className="px-6 py-2">
                    <Button variant={"danger"} disabled={deleteUserMutation.isLoading} className={"!p-1"}
                            onClick={() => {
                                setToBeDeletedUserEmail(user.email)
                                setShowAbortDialog(true)
                            }}>
                        {toBeDeletedUserEmail === user.email && deleteUserMutation.isLoading ?
                            <Spinner/> : <YgoIcon icon={"trash"} size={18} classNames={"fill-white"}/>}
                    </Button>
                </td>
            </tr>
            tableEntries.push(entry)
        })

        return tableEntries
    }

    function createModal(): JSX.Element {
        return <ConfirmModal show={showAbortDialog}
                             setShow={setShowAbortDialog}
                             confirmName={"Delete"}
                             onConfirm={() => deleteSelectedUser()}
                             title={"Delete User"}
                             description={`The user ${toBeDeletedUserEmail} is going to be deleted. This cannot be reversed!`}/>
    }

    function createPlaceholderEntries(): JSX.Element[] {
        const tableEntries: JSX.Element[] = []

        const placeholderElement = <td>
            <p className={"placeholder-glow"}>
                <span className={"placeholder w-100 px-6 py-3"}>
                    1
                </span>
            </p>
        </td>
        for (let i = 0; i < 15; i++) {
            const entry = <tr key={`row-placeholder-${i}`}
                              className={"odd:bg-light-1 odd:dark:bg-dark-1 even:bg-light-2 even:dark:bg-dark-2 border-l border-r border-dark dark:border-light"}>
                {placeholderElement}
                {placeholderElement}
                {placeholderElement}
                {placeholderElement}
                {placeholderElement}
            </tr>
            tableEntries.push(entry)
        }

        return tableEntries
    }

    function getTableHeader() {
        const tableCN = classNames("text-xs text-dark dark:text-light uppercase", "bg-light-3 dark:bg-dark-3", "border border-dark dark:border-light")
        return <thead className={tableCN}>
        <tr>
            <th className="px-6 py-3 rounded-tl text-sm">ID</th>
            <th className="px-6 py-3 text-sm">Email</th>
            <th className="px-6 py-3 text-sm">Display Name</th>
            <th className="px-6 py-3 text-sm">Administrator</th>
            <th className="px-6 py-3 text-sm rounded-tr">Actions</th>
        </tr>
        </thead>;
    }

    function getTableControls() {
        if (isLoading || error || !data) {
            return <></>
        }

        let minUser = (currentPage * currentPageSize)
        let maxUser = ((currentPage + 1) * currentPageSize)
        if (data.numberOfUsers < maxUser) {
            maxUser = data.numberOfUsers
        }

        const allButtonCN = classNames("py-2 px-4", "text-sm font-semibold", "rounded-tr-none rounded-tl-none", "border !border-dark !dark:border-light")
        return <div
            className="flex justify-content-between items-center justify-between">
            <Button
                variant={"primary"}
                disabled={currentPage === 0}
                className={classNames(allButtonCN, "rounded-br-none")}
                onClick={() => {
                    let newPage = currentPage - 1;
                    setCurrentPage(newPage < 0 ? 0 : newPage)
                }
                }>
                Prev
            </Button>
            <span
                className="text-xs xs:text-sm text-dark dark:text-light bg-light-3 dark:bg-dark-3 flex-grow py-2 text-center place-self-stretch border-t border-b border-dark dark:border-light">
                                Showing {minUser} to {maxUser} of {data.numberOfUsers} Users
                        </span>
            <Button
                variant={"primary"}
                disabled={data.numberOfUsers === maxUser}
                className={classNames(allButtonCN, "rounded-bl-none")}
                onClick={() => {
                    let newPage = currentPage + 1;
                    setCurrentPage(newPage)
                }
                }>
                Next
            </Button>
        </div>
    }

    return <div>
        {createModal()}
        <table className={"w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"}>
            {getTableHeader()}
            <tbody>
            {tableEntries}
            </tbody>
        </table>
        {getTableControls()}
    </div>
}

export default UserManagement

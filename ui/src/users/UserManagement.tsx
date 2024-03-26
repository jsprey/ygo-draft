import React, {useState} from "react";
import {User, useUsers} from "../api/hooks/useUsers";
import {Alert, Button, Modal, Spinner} from "react-bootstrap";
import classNames from "classnames";
import {useQueryClient} from "react-query";
import {enqueueSnackbar} from "notistack";
import {useDeleteUser} from "../api/hooks/useDeleteUser";
import SvgIconButton from "../core/SvgIconButton";

const TrashIcon = <SvgIconButton size={18} classNames={"fill-white"}>
    <path
        d="M2.5 1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1H3v9a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V4h.5a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H10a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1zm3 4a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 .5-.5M8 5a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-1 0v-7A.5.5 0 0 1 8 5m3 .5v7a.5.5 0 0 1-1 0v-7a.5.5 0 0 1 1 0"/>
</SvgIconButton>

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
        queryClient.invalidateQueries({queryKey: "users"})
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
            <td colSpan={5}><Alert variant={"danger"}>Failed to load users!</Alert></td>
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
                              className={"odd:bg-gray-50 odd:dark:bg-gray-900 even:bg-gray-100 even:dark:bg-gray-800 border-b dark:border-gray-700"}>
                <td className="px-6 py-2">{user.id}</td>
                <td className="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">{user.email}</td>
                <td className="px-6 py-2">{user.display_name}</td>
                <td className="px-6 py-2">{user.is_admin ? "Admin" : "User"}</td>
                <td className="px-6 py-2">
                    <button disabled={deleteUserMutation.isLoading} className={"btn btn-danger"} onClick={() => {
                        setToBeDeletedUserEmail(user.email)
                        setShowAbortDialog(true)
                    }}>
                        {toBeDeletedUserEmail === user.email && deleteUserMutation.isLoading ?
                            <Spinner size={"sm"} animation={"border"}/> : TrashIcon}
                    </button>
                </td>
            </tr>
            tableEntries.push(entry)
        })

        return tableEntries
    }

    function createModal(): JSX.Element {
        return <Modal show={showAbortDialog}>
            <Modal.Header closeButton className={"bg-ygo-light dark:bg-ygo-dark dark:text-white"}>
                <Modal.Title>Delete User?</Modal.Title>
            </Modal.Header>
            <Modal.Body className={"bg-ygo-light dark:bg-ygo-dark dark:text-white"}>The user {toBeDeletedUserEmail} is
                going to be
                deleted. This cannot be reversed.</Modal.Body>
            <Modal.Footer className={"bg-ygo-light dark:bg-ygo-dark dark:text-white"}>
                <Button variant="secondary" onClick={() => setShowAbortDialog(false)}>
                    No
                </Button>
                <Button variant="danger" onClick={() => {
                    deleteSelectedUser()
                    setShowAbortDialog(false)
                }}>
                    Yes
                </Button>
            </Modal.Footer>
        </Modal>
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
                              className={"odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700"}>
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
        return <thead className={"text-xs text-gray-700 uppercase bg-gray-200 dark:bg-gray-700 dark:text-gray-400"}>
        <tr>
            <th className="px-6 py-3 rounded-tl">ID</th>
            <th className="px-6 py-3">Email</th>
            <th className="px-6 py-3">Display Name</th>
            <th className="px-6 py-3">Administrator</th>
            <th className="px-6 py-3 rounded-tr">Actions</th>
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

        const allButtonCN = classNames("py-2 px-4", "text-sm font-semibold", "bg-gray-300 hover:bg-gray-400 active:bg-gray-500 disabled:bg-gray-200", "text-gray-800 disabled:text-gray-300")
        return <div
            className="flex justify-content-between xs:flex-row items-center xs:justify-between">
            <button
                disabled={currentPage === 0}
                className={classNames(allButtonCN, "rounded-bl")}
                onClick={() => {
                    let newPage = currentPage - 1;
                    setCurrentPage(newPage < 0 ? 0 : newPage)
                }
                }>
                Prev
            </button>
            <span
                className="text-xs xs:text-sm text-gray-700 bg-gray-200 dark:bg-gray-700 dark:text-gray-400 flex-grow-1 py-2 text-center place-self-stretch">
                            Showing {minUser} to {maxUser} of {data.numberOfUsers} Users
                        </span>
            <button
                disabled={data.numberOfUsers === maxUser}
                className={classNames(allButtonCN, "rounded-br")}
                onClick={() => {
                    let newPage = currentPage + 1;
                    setCurrentPage(newPage)
                }
                }>
                Next
            </button>
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

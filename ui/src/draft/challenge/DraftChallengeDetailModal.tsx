import {Modal, Spinner} from "react-bootstrap";
import {Draft} from "../../api/Draft";
import {useAcceptDraftChallenge} from "../../api/hooks/drafts/useAcceptDraftChallenge";
import {useDeclineDraftChallenge} from "../../api/hooks/drafts/useDeclineDraftChallenge";
import {enqueueSnackbar} from "notistack";
import DraftSettingsDetails from "./DraftSettingsDetails";
import React from "react";
import {useQueryClient} from "react-query";

export type DraftChallengeDetailModalProps = {
    challenge: Draft
    isShowing: boolean
    setShow: React.Dispatch<React.SetStateAction<boolean>>
}

function DraftChallengeDetailModal(props: DraftChallengeDetailModalProps) {
    const handleClose = () => props.setShow(false);
    const queryClient = useQueryClient();

    function showSuccess(message: string) {
        enqueueSnackbar(message, {
            autoHideDuration: 6000,
            variant: "success"
        })

        props.setShow(false)
    }
    function showError(message: string) {
        enqueueSnackbar(message, {
            autoHideDuration: 6000,
            variant: "error"
        })
    }

    const acceptChallengeMutation = useAcceptDraftChallenge({
        onSuccess: () => {
            showSuccess("Challenge accepted.")
            queryClient.refetchQueries({queryKey: ["drafts", "challenges"]})
            queryClient.refetchQueries({queryKey: ["drafts", "running"]})
        },
        onError: () => showError("Failed to accept challenge. Try again and/or contact the support.")
    });
    const declineChallengeMutation = useDeclineDraftChallenge({
        onSuccess: () => {
            showSuccess("Challenge declined.")
            queryClient.refetchQueries({queryKey: ["drafts", "challenges"]})
            queryClient.refetchQueries({queryKey: ["drafts", "running"]})
        },
        onError: () => showError("Failed to decline challenge. Try again and/or contact the support.")
    });

    if (!props.isShowing) {
        return <></>
    }

    const settings = props.challenge.settings
    return <Modal show={props.isShowing}
                  onHide={handleClose}
                  size={"xl"}
                  contentClassName={""}>
        <Modal.Header className={"bg-ygo-light dark:bg-ygo-dark border dark:text-white"}>
            <div className={"text-xl fw-bold"}>
                You got a challenge!
            </div>
        </Modal.Header>
        <Modal.Body className={"bg-ygo-light dark:bg-ygo-dark border dark:text-white"}>
            <div className={"fw-bold"}>Settings</div>
            <DraftSettingsDetails settings={settings}/>
        </Modal.Body>
        <Modal.Footer className={"bg-ygo-light dark:bg-ygo-dark border dark:text-white"}>
            <button disabled={declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading} className={"btn btn-success"} onClick={() => {
                acceptChallengeMutation.mutate(props.challenge.id)
            }
            }>
                {declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading ? <Spinner animation={"border"} size={"sm"}></Spinner> : "Accept"}
            </button>
            <button disabled={declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading} className={"btn btn-danger"} onClick={() => {
                declineChallengeMutation.mutate(props.challenge.id)
            }
            }>
                {declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading ? <Spinner animation={"border"} size={"sm"}></Spinner> : "Decline"}
            </button>
        </Modal.Footer>
    </Modal>
}

export default DraftChallengeDetailModal
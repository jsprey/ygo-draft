import {Draft} from "../../../api/Draft";
import {useAcceptDraftChallenge} from "../../../api/hooks/drafts/useAcceptDraftChallenge";
import {useDeclineDraftChallenge} from "../../../api/hooks/drafts/useDeclineDraftChallenge";
import {enqueueSnackbar} from "notistack";
import DraftSettingsDetails from "../../shared/settingsPage/DraftSettingsDetails";
import React from "react";
import {useQueryClient} from "react-query";
import Spinner from "../../../core/Spinner";
import Modal from "../../../core/Modal";
import Button from "../../../core/Button";

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
                  setShow={props.setShow}
                  onHide={handleClose}>
        <div className={"text-2xl font-bold"}>
            You got a challenge!
        </div>
        <div className={"mt-2 font-semibold"}>Settings</div>
        <DraftSettingsDetails settings={settings}/>
        <div className={"mt-2 flex justify-end"}>
            <Button variant={"primary"} disabled={declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading}
                    className={"mr-2"} onClick={() => {
                acceptChallengeMutation.mutate(props.challenge.id)
            }
            }>
                {declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading ? <Spinner/> : "Accept"}
            </Button>
            <Button variant={"danger"} disabled={declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading}
                    className={""} onClick={() => {
                declineChallengeMutation.mutate(props.challenge.id)
            }
            }>
                {declineChallengeMutation.isLoading || acceptChallengeMutation.isLoading ? <Spinner/> : "Decline"}
            </Button>
        </div>
    </Modal>
}

export default DraftChallengeDetailModal